package api

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
)

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Cause   string `json:"cause"`
	Path    string `json:"path"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Message, e.Cause)
}

func InternalServerError(r *http.Request, cause string) Error {
	return Error{
		Status:  http.StatusInternalServerError,
		Message: "Internal Server Error",
		Cause:   cause,
		Path:    r.URL.Path,
	}
}

func NotFoundError(r *http.Request, cause string) Error {
	return Error{
		Status:  http.StatusNotFound,
		Message: "Resource not found",
		Cause:   cause,
		Path:    r.URL.Path,
	}
}

func BadRequestError(r *http.Request, cause string) Error {
	return Error{
		Status:  http.StatusBadRequest,
		Message: "Bad Request",
		Cause:   cause,
		Path:    r.URL.Path,
	}
}

func UnauthorisedError(r *http.Request) Error {
	return Error{
		Status:  http.StatusUnauthorized,
		Message: "Unauthorized",
		Cause:   "User is unauthorised for the request",
		Path:    r.URL.Path,
	}
}

func WriteError(w http.ResponseWriter, r *http.Request, err error) error {
	var apiError Error
	if errors.As(err, &apiError) {
		if err := WriteJSON(w, apiError, apiError.Status); err != nil {
			return err
		}
		return nil
	}

	slog.Error("Server error", "error", err.Error())
	err = WriteJSON(w, InternalServerError(r, "Something went wrong"), int(http.StatusInternalServerError))
	if err != nil {
		return err
	}

	return nil
}
