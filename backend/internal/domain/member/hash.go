package member

import "golang.org/x/crypto/bcrypt"

type PasswordHash []byte

// Checks if the password hash matches with the given password.
func (h PasswordHash) Matches(password []byte) error {
	err := bcrypt.CompareHashAndPassword(h, password)
	if err != nil {
		return err
	}

	return nil
}
