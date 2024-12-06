package service

import (
	"context"
	"errors"
	"net/mail"
	"strings"

	"github.com/mattismoel/konnekt/internal/storage"
)

var (
	ErrInvalidEmail     = errors.New("Invalid email")
	ErrInvalidFirstName = errors.New("Invalid first name")
	ErrInvalidLastName  = errors.New("Invalid last name")
)

type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type UserFilter struct {
	ID        *int64  `json:"id"`
	Email     *string `json:"email"`
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`

	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type userRepository interface {
	InsertUser(ctx context.Context, user storage.User) (storage.User, error)
	DeleteUser(ctx context.Context, id int64) error
	UpdateUser(ctx context.Context, id int64, update storage.User) (storage.User, error)
	FindUserByEmail(ctx context.Context, email string) (storage.User, error)
	FindUserPasswordHash(ctx context.Context, userID int64) ([]byte, error)
}

type userService struct {
	repo userRepository
}

func NewUserService(repo userRepository) *userService {
	return &userService{repo: repo}
}

func (u User) Validate() error {
	if strings.TrimSpace(u.FirstName) == "" {
		return ErrInvalidFirstName
	}

	if strings.TrimSpace(u.LastName) == "" {
		return ErrInvalidLastName
	}

	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return ErrInvalidEmail
	}

	return nil
}

func (u1 User) Equals(u2 User) bool {
	if strings.TrimSpace(u1.Email) != strings.TrimSpace(u2.Email) {
		return false
	}

	if strings.TrimSpace(u1.FirstName) != strings.TrimSpace(u2.FirstName) {
		return false
	}

	if strings.TrimSpace(u1.LastName) != strings.TrimSpace(u2.LastName) {
		return false
	}

	if u1.ID != u2.ID {
		return false
	}

	return true
}
