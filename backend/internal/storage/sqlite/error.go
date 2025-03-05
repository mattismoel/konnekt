package sqlite

import "errors"

var (
	ErrUserAlreadyExists = errors.New("User already exists")
	ErrNotFound          = errors.New("Resource not found")
)
