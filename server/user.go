package konnekt

import (
	"net/mail"
	"strings"

	"github.com/mattismoel/konnekt"
)

type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
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
