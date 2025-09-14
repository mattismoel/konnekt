package konnekt

import "errors"

var (
	ErrResourceNotFound = errors.New("Resource does not exist")
	ErrAlreadyExists    = errors.New("Resource already exists")
)
