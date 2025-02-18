package service

import (
	"context"

	"github.com/mattismoel/konnekt/internal/domain/user"
)

type UserService struct {
	userRepo user.Repository
}

func NewUserService(userRepo user.Repository) (*UserService, error) {
	return &UserService{
		userRepo: userRepo,
	}, nil
}

func (srv UserService) ByID(ctx context.Context, userID int64) (user.User, error) {
	u, err := srv.userRepo.ByID(ctx, userID)
	if err != nil {
		return user.User{}, err
	}

	return u, nil
}
