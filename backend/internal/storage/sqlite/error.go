package sqlite

import "errors"

var (
	ErrAlreadyExists = errors.New("Entity already exists")
	ErrNotFound      = errors.New("Resource not found")
)
