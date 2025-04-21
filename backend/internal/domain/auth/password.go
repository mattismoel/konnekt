package auth

import (
	"bytes"
	"errors"
	"unicode/utf8"
)

const (
	MINIMUM_PASSWORD_LENGTH = 8  // Minimum length of a members password.
	MAXIMUM_PASSWORD_LENGTH = 24 // Maximum length of a members password.
)

var (
	ErrPasswordTooShort = errors.New("Password must be at least 8 characters long")
	ErrPasswordTooLong  = errors.New("Password must be max 24 characters long")

	ErrPasswordsNoMatch = errors.New("Passwords do not match")
)

type Password []byte

func (p1 Password) Matches(p2 Password) error {
	if !bytes.Equal(p1, p2) {
		return ErrPasswordsNoMatch
	}

	return nil
}

// Checks whether a Password is valid for use.
func (p Password) Validate(b []byte) error {
	passLength := utf8.RuneCount(b)

	if passLength < MINIMUM_PASSWORD_LENGTH {
		return ErrPasswordTooShort
	}

	if passLength > MAXIMUM_PASSWORD_LENGTH {
		return ErrPasswordTooLong
	}

	return nil
}
