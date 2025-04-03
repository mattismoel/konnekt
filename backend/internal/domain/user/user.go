package user

import (
	"errors"
	"net/mail"
	"strings"
)

var (
	ErrIDInvalid           = errors.New("ID must be a positive integer")
	ErrFirstNameInvalid    = errors.New("First name must be valid and non-empty")
	ErrLastNameInvalid     = errors.New("Last name must be valid and non-empty")
	ErrEmailInvalid        = errors.New("Email must be valid")
	ErrPasswordHashInvalid = errors.New("Password hash must be a non-empty byte array")
)

type User struct {
	ID           int64        `json:"id"`
	FirstName    string       `json:"firstName"`
	LastName     string       `json:"lastName"`
	Email        string       `json:"email"`
	PasswordHash PasswordHash `json:"-"`
}

type cfgFunc func(u *User) error

func NewUser(cfgs ...cfgFunc) (User, error) {
	u := &User{}

	for _, cfg := range cfgs {
		if err := cfg(u); err != nil {
			return User{}, err
		}
	}

	return *u, nil
}

func WithID(id int64) cfgFunc {
	return func(u *User) error {
		if id <= 0 {
			return ErrIDInvalid
		}

		u.ID = id

		return nil
	}
}

func WithFirstName(firstName string) cfgFunc {
	firstName = strings.TrimSpace(firstName)
	return func(u *User) error {
		if firstName == "" {
			return ErrFirstNameInvalid
		}

		u.FirstName = firstName

		return nil
	}
}

func WithLastName(lastName string) cfgFunc {
	lastName = strings.TrimSpace(lastName)

	return func(u *User) error {
		if lastName == "" {
			return ErrLastNameInvalid
		}

		u.LastName = lastName

		return nil
	}
}

func WithEmail(email string) cfgFunc {
	email = strings.TrimSpace(email)

	return func(u *User) error {
		if email == "" {
			return ErrEmailInvalid
		}

		m, err := mail.ParseAddress(email)
		if err != nil {
			return ErrEmailInvalid
		}

		u.Email = m.Address

		return nil
	}
}

func WithPasswordHash(hash []byte) cfgFunc {
	return func(u *User) error {
		if len(hash) <= 0 {
			return ErrPasswordHashInvalid
		}

		u.PasswordHash = hash

		return nil
	}
}
