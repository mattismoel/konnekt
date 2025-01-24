package service

import (
	"github.com/mattismoel/konnekt/internal/domain/auth"
	"github.com/mattismoel/konnekt/internal/domain/user"
	"golang.org/x/crypto/bcrypt"
)

const (
	SESSION_LIFETIME       = 30 * 24 * time.Hour // 30 day expiry.
)
type AuthService struct {
	userRepo user.Repository
	authRepo auth.Repository
}

func NewAuthService(userRepo user.Repository, authRepo auth.Repository) (*AuthService, error) {
	return &AuthService{
		userRepo: userRepo,
		authRepo: authRepo,
	}, nil
}

func (srv AuthService) Register(ctx context.Context, email string, password []byte, passwordConfirm []byte, firstName, lastName string) (auth.Session, auth.SessionToken, error) {
	// Return if user already exists.
	_, err := srv.userRepo.ByEmail(ctx, email)
	if err != nil {
		return auth.Session{}, "", user.ErrAlreadyExists
	}

	if err := auth.DoPasswordsMatch(password, passwordConfirm); err != nil {
		return auth.Session{}, "", err
	}

	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return auth.Session{}, "", err
	}

	userID, err := srv.userRepo.Insert(ctx, email, firstName, lastName, hash)
	if err != nil {
		return auth.Session{}, "", err
	}

	sessionToken, err := auth.NewSessionToken()
	if err != nil {
		return auth.Session{}, "", err
	}

	session := auth.NewSession(sessionToken, userID, SESSION_LIFETIME)

	if err := srv.authRepo.InsertSession(ctx, session); err != nil {
		return auth.Session{}, "", err
	}

	return session, sessionToken, nil
}
