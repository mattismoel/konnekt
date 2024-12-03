package password

import (
	"bytes"
	"fmt"
	"slices"

	"golang.org/x/crypto/bcrypt"
)

const (
	MIN_LENGTH      = 8
	MAX_LENGTH      = 28
	UPPERCASE_CHARS = "ABCDEFGHIJKLMNOPQRSTUVXYZ"
	LOWERCASE_CHARS = "abcdefghijklmnopqrstuvxyz"
	SPECIAL_CHARS   = "!@#$%^&*()"
)

var (
	ErrTooShort         = fmt.Errorf("Too short. Must be at least %d characters long", MIN_LENGTH)
	ErrTooLong          = fmt.Errorf("Too long. Must be max %d characters long", MAX_LENGTH)
	ErrNoUppercaseChars = fmt.Errorf("Must contain at least one uppercase character")
	ErrNoSpecialChars   = fmt.Errorf("Must contain special character - %s", SPECIAL_CHARS)
)

type Password []byte

// Hashes the password.
func (p Password) Hash() ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(p, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return hash, nil
}

// Returns whether or not the password matches with the given hash.
func (p Password) MatchesHash(hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, p)
	if err != nil {
		return false
	}

	return true
}

// Returns whether or not p1 is equal to p2.
func (p1 Password) Equals(p2 Password) bool {
	return slices.Equal(p1, p2)
}

// Validates that the password is strong based on length and contents.
func (p Password) Validate() []error {
	errors := []error{}

	if len(p) < MIN_LENGTH {
		errors = append(errors, ErrTooShort)
	}

	if len(p) > MAX_LENGTH {
		errors = append(errors, ErrTooLong)
	}

	if !bytes.ContainsAny(p, SPECIAL_CHARS) {
		errors = append(errors, ErrNoSpecialChars)
	}

	if !bytes.ContainsAny(p, UPPERCASE_CHARS) {
		errors = append(errors, ErrNoUppercaseChars)
	}

	if len(errors) <= 0 {
		return nil
	}

	return errors
}
