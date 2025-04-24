package member

import "errors"

var (
	ErrNotFound      = errors.New("Member not found")
	ErrAlreadyExists = errors.New("Member already exists")
)
