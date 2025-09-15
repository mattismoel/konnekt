package auth

import (
	"bytes"
	"errors"
)

var (
	ErrPasswordNoMatch = errors.New("Passwords must match")
)

// Checks whether the input password is equal to the confirm password.
func ValidatePasswordConfirm(p []byte, confirm []byte) error {
	if !bytes.Equal(p, confirm) {
		return ErrPasswordNoMatch
	}

	return nil
}
