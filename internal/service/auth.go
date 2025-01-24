package service

import (
	"github.com/mattismoel/konnekt/internal/domain/auth"
	"github.com/mattismoel/konnekt/internal/domain/user"
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
