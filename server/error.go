package konnekt

import (
	"errors"
	"fmt"
)

const (
	ERRCONFLICT       = "conflict"
	ERRINTERNAL       = "internal"
	ERRINVALID        = "invalid"
	ERRNOTFOUND       = "not_found"
	ERRNOTIMPLEMENTED = "not_implemented"
	ERRUNAUTHORIZED   = "unauthorized"
)

type Error struct {
	Code    string
	Message string
}

func (e Error) Error() string {
	return fmt.Sprintf("error: code=%s message=%s", e.Code, e.Message)
}

func ErrorCode(err error) string {
	var e Error

	if err == nil {
		return ""
	}

	if errors.As(err, &e) {
		return e.Code
	}

	return ERRINTERNAL
}

func ErrorMessage(err error) string {
	var e Error

	if err == nil {
		return ""
	}

	if errors.As(err, &e) {
		return e.Message
	}

	return "Internal error"
}

func Errorf(code string, format string, args ...any) Error {
	return Error{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
	}
}
