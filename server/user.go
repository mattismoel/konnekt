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

func (u User) Validate() error {
	if u.ID < 0 {
		return konnekt.Errorf(konnekt.ERRINVALID, "ID must be a non-negative integer")
	}

	if strings.TrimSpace(u.FirstName) == "" {
		return konnekt.Error(konnekt.ERRINVALID, "First name must be set")
	}

	if strings.TrimSpace(u.LastName) == "" {
		return konnekt.Error(konnekt.ERRINVALID, "Last name must be set")
	}

	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return konnekt.Error(konnekt.ERRINVALID, "Invalid email")
	}

	return nil
}
