package service

import (
	"context"
	"errors"
	"time"

	"github.com/mattismoel/konnekt/internal/domain/auth"
	"github.com/mattismoel/konnekt/internal/domain/user"
	"golang.org/x/crypto/bcrypt"
)

const (
	SESSION_LIFETIME       = 30 * 24 * time.Hour // 30 day expiry.
	SESSION_REFRESH_BUFFER = 15 * 24 * time.Hour // 15 day refresh buffer.
)

var (
	ErrNoSession = errors.New("No such session")
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

func (srv AuthService) Register(ctx context.Context, email string, password []byte, passwordConfirm []byte, firstName, lastName string) (auth.SessionToken, time.Time, error) {
	// Return if user already exists.
	_, err := srv.userRepo.ByEmail(ctx, email)
	if err != nil {
		return "", time.Time{}, user.ErrAlreadyExists
	}

	if err := auth.DoPasswordsMatch(password, passwordConfirm); err != nil {
		return "", time.Time{}, err
	}

	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", time.Time{}, err
	}

	userID, err := srv.userRepo.Insert(ctx, email, firstName, lastName, hash)
	if err != nil {
		return "", time.Time{}, err
	}

	token, expiry, err := srv.createSession(ctx, userID)
	if err != nil {
		return "", time.Time{}, err
	}

	return token, expiry, nil
}

func (srv AuthService) validateUser(ctx context.Context, email string, password []byte) (user.User, error) {
	// Return early if user does not exist.
	usr, err := srv.userRepo.ByEmail(ctx, email)
	if err != nil {
		return user.User{}, err
	}

	hash, err := srv.userRepo.PasswordHash(ctx, usr.ID)
	if err != nil {
		return user.User{}, err
	}

	if err := hash.Matches(password); err != nil {
		return user.User{}, err
	}

	return usr, err
}

func (srv AuthService) clearUserSession(ctx context.Context, userID int64) error {
	err := srv.authRepo.DeleteUserSession(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}

func (srv AuthService) createSession(ctx context.Context, userID int64) (auth.SessionToken, time.Time, error) {
	token, err := auth.NewSessionToken()
	if err != nil {
		return "", time.Time{}, err
	}

	session := auth.NewSession(token, userID, SESSION_LIFETIME)

	err = srv.authRepo.InsertSession(ctx, session)
	if err != nil {
		return "", time.Time{}, err
	}

	return token, session.ExpiresAt, nil
}

func (srv AuthService) Login(ctx context.Context, email string, password []byte) (auth.SessionToken, time.Time, error) {
	usr, err := srv.validateUser(ctx, email, password)
	if err != nil {
		return "", time.Time{}, err
	}

	err = srv.clearUserSession(ctx, usr.ID)
	if err != nil {
		return "", time.Time{}, err
	}

	token, expiry, err := srv.createSession(ctx, usr.ID)
	if err != nil {
		return "", time.Time{}, err
	}

	return token, expiry, nil
}
