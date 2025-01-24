package server

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
)

const (
	SERVER_ERR_MESSAGE = "Something went wrong"
)

type APIError struct {
	Message string `json:"message"`
	Status  int    `json:"-"`
}

func newAPIError(msg string, status int) APIError {
	return APIError{
		Message: msg,
		Status:  status,
	}
}

func (e APIError) Error() string {
	return fmt.Sprintf("%s", e.Message)
}

func writeError(w http.ResponseWriter, err error) {
	var apiErr APIError
	if ok := errors.As(err, &apiErr); ok {
		writeJSON(w, apiErr.Status, apiErr)
		return
	}

	writeJSON(w, http.StatusInternalServerError, APIError{
		Message: SERVER_ERR_MESSAGE,
	})

	slog.Error("Interal server error", "error", err.Error())

	return

}
