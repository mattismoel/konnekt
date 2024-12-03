package konnekt

import "context"

type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type UserService interface {
	FindUserByID(context.Context, int64) (User, error)
	FindUsers(context.Context, UserFilter) ([]User, error)

	CreateUser(context.Context, User) (int64, error)

	UpdateUser(context.Context, UpdateUser) (User, error)
	DeleteUser(context.Context, int64) error
}

type UserFilter struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`

	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type UpdateUser struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
