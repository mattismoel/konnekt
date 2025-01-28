package auth

import (
	"bytes"
	"errors"
)

var (
	ErrPasswordsNoMatch = errors.New("Passwords do not match")
)

func DoPasswordsMatch(password, passwordConfirm []byte) error {
	if !bytes.Equal(password, passwordConfirm) {
		return ErrPasswordsNoMatch
	}

	return nil
}
