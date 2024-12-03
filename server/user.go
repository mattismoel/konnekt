package konnekt

import (
	"net/mail"
	"strings"
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

func (u User) Validate() error {
	if u.ID < 0 {
		return Errorf(ERRINVALID, "ID must be a non-negative integer")
	}

	if strings.TrimSpace(u.FirstName) == "" {
		return Errorf(ERRINVALID, "First name must be set")
	}

	if strings.TrimSpace(u.LastName) == "" {
		return Errorf(ERRINVALID, "Last name must be set")
	}

	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return Errorf(ERRINVALID, "Invalid email")
	}

	return nil
}
