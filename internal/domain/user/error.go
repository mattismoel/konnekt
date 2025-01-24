package user

import "errors"

var (
	ErrNotFound      = errors.New("User not found")
	ErrAlreadyExists = errors.New("User already exists")
)
