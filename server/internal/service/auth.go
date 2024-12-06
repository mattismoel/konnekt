package service

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/mattismoel/konnekt/internal/password"
	"github.com/mattismoel/konnekt/internal/storage"
	"golang.org/x/crypto/bcrypt"
)

const SESSION_LIFETIME = 24 * 24 * time.Hour
const SESSION_EXTEND_LIFETIME_BUFFER = 12 * 24 * time.Hour
const SESSION_COOKIE_NAME = "session"

const (
	PERM_CREATE_EVENT = "event-create"
	PERM_UPDATE_EVENT = "event-update"
	PERM_DELETE_EVENT = "event-delete"
)

var (
	ErrNoSession      = errors.New("No session found")
	ErrSessionExpired = errors.New("Session has expired")
)

type SessionToken []byte
type SessionID string

type Session struct {
	ID        SessionID `json:"id"`
	UserID    int64     `json:"userId"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type sessionRepository interface {
	InsertSession(ctx context.Context, sessionID string, userID int64, expiry time.Time) (storage.Session, error)
	FindSession(ctx context.Context, sessionID string) (storage.Session, storage.User, error)
	DeleteSession(ctx context.Context, sessionID string) error
	UpdateSessionExpiry(ctx context.Context, sessionID string, newExpiry time.Time) (storage.Session, error)
}

type permissionRepository interface {
	HasPermission(ctx context.Context, userID int64, perm string) error
}

type authService struct {
	sessionRepo sessionRepository
	permRepo    permissionRepository
	userRepo    userRepository
}

func NewAuthService(sessionRepo sessionRepository, userRepo userRepository, permRepo permissionRepository) *authService {
	return &authService{
		sessionRepo: sessionRepo,
		userRepo:    userRepo,
		permRepo:    permRepo,
	}
}

func (s Session) HasExpired() bool {
	now := time.Now()

	return now.After(s.ExpiresAt) || now.Equal(s.ExpiresAt)
}

func (s authService) GenerateSessionToken() (SessionToken, error) {
	bytes := make([]byte, 20)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}

	encoder := base32.StdEncoding.WithPadding(base32.NoPadding)
	token := encoder.EncodeToString(bytes)
	return []byte(token), nil
}

// Returns whether or not a session should be extended.
func (s Session) ShouldExtend() bool {
	now := time.Now()

	return now.After(s.ExpiresAt.Add(-SESSION_EXTEND_LIFETIME_BUFFER))
}

// Generates a session id based on the session token.
func (t SessionToken) GenerateSessionID() string {
	hash := sha256.Sum256(t)

	sessionId := hex.EncodeToString(hash[:])
	return sessionId
}

func (s authService) CreateSession(ctx context.Context, token SessionToken, userID int64) (Session, error) {
	sessionID := token.GenerateSessionID()

	repoSession, err := s.sessionRepo.InsertSession(ctx, sessionID, userID, time.Now().Add(SESSION_LIFETIME))
	if err != nil {
		return Session{}, err
	}

	session := Session{
		ID:        SessionID(repoSession.ID),
		UserID:    repoSession.UserID,
		ExpiresAt: repoSession.ExpiresAt,
	}

	return session, nil
}

func (s authService) ValidateSessionToken(ctx context.Context, token SessionToken) (Session, User, error) {
	sessionID := token.GenerateSessionID()

	repoSession, repoUser, err := s.sessionRepo.FindSession(ctx, sessionID)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return Session{}, User{}, ErrNoSession
		}

		return Session{}, User{}, err
	}

	session := Session{
		ID:        SessionID(repoSession.ID),
		UserID:    repoSession.UserID,
		ExpiresAt: repoSession.ExpiresAt,
	}

	user := User{
		ID:        repoUser.ID,
		FirstName: repoUser.FirstName,
		LastName:  repoUser.LastName,
		Email:     repoUser.Email,
	}

	if session.HasExpired() {
		err := s.sessionRepo.DeleteSession(ctx, sessionID)
		if err != nil {
			return Session{}, User{}, err
		}

		return Session{}, User{}, ErrSessionExpired
	}

	if session.ShouldExtend() {
		newExpiry := time.Now().Add(SESSION_LIFETIME)
		_, err := s.sessionRepo.UpdateSessionExpiry(ctx, sessionID, newExpiry)
		if err != nil {
			return Session{}, User{}, err
		}

		session.ExpiresAt = newExpiry
	}

	return session, user, nil
}

func (s authService) InvalidateSession(ctx context.Context, sessionID SessionID) error {
	err := s.sessionRepo.DeleteSession(ctx, string(sessionID))
	if err != nil {
		return err
	}

	return nil
}

func (s authService) SetSessionTokenCookie(w http.ResponseWriter, token SessionToken, expiresAt time.Time) {
	cookie := &http.Cookie{
		Name:     SESSION_COOKIE_NAME,
		Value:    string(token),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Expires:  expiresAt,
		Path:     "/",
	}

	http.SetCookie(w, cookie)
}

func (s authService) DeleteSessionTokenCookie(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:     SESSION_COOKIE_NAME,
		Value:    "",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   0,
		Path:     "/",
	}

	http.SetCookie(w, cookie)
}

func (s authService) HasPermission(ctx context.Context, userID int64, perm string) error {
	err := s.permRepo.HasPermission(ctx, userID, perm)
	if err != nil {
		return Errorf(ERRUNAUTHORIZED, "Not authorized")
	}

	return nil
}

func (s authService) Login(ctx context.Context, w http.ResponseWriter, r *http.Request, email string, password []byte) (User, error) {
	repoUser, err := s.userRepo.FindUserByEmail(ctx, email)
	if err != nil {
		return User{}, Errorf(ERRUNAUTHORIZED, err.Error())
	}

	userPasswordHash, err := s.userRepo.FindUserPasswordHash(ctx, repoUser.ID)
	if err != nil {
		return User{}, Errorf(ERRUNAUTHORIZED, err.Error())
	}

	err = bcrypt.CompareHashAndPassword(userPasswordHash, password)
	if err != nil {
		return User{}, Errorf(ERRUNAUTHORIZED, "Invalid credentials")
	}

	token, err := s.GenerateSessionToken()
	if err != nil {
		return User{}, Errorf(ERRUNAUTHORIZED, err.Error())
	}

	cookie, err := r.Cookie(SESSION_COOKIE_NAME)
	if err != nil && !errors.Is(err, http.ErrNoCookie) {
		return User{}, Errorf(ERRINTERNAL, err.Error())
	}

	prevToken := SessionToken(cookie.Value)

	prevSession, _, err := s.ValidateSessionToken(ctx, prevToken)
	if err == nil {
		err := s.InvalidateSession(ctx, prevSession.ID)
		if err != nil {
			return User{}, nil
		}
	}

	session, err := s.CreateSession(ctx, token, repoUser.ID)
	if err != nil {
		return User{}, Errorf(ERRUNAUTHORIZED, err.Error())
	}

	s.SetSessionTokenCookie(w, token, session.ExpiresAt)

	user := User{
		ID:        repoUser.ID,
		FirstName: repoUser.FirstName,
		LastName:  repoUser.LastName,
		Email:     repoUser.Email,
	}

	return user, nil
}

func (s authService) Register(ctx context.Context, w http.ResponseWriter, user User, password password.Password, passwordConfirm password.Password) (User, error) {
	err := user.Validate()
	if err != nil {
		return User{}, err
	}

	passwordErrors := password.Validate()
	if passwordErrors != nil {
		return User{}, Errorf(ERRINVALID, fmt.Sprint(passwordErrors))
	}

	if !password.Equals(passwordConfirm) {
		return User{}, Errorf(ERRINVALID, "Passwords do not match")
	}

	passwordHash, err := password.Hash()
	if err != nil {
		return User{}, err
	}

	repoUser := storage.User{
		Email:        user.Email,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		PasswordHash: passwordHash,
	}

	repoUser, err = s.userRepo.InsertUser(ctx, repoUser)
	if err != nil {
		return User{}, err
	}

	user.ID = repoUser.ID

	return user, nil
}
