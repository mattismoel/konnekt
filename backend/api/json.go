package api

import (
	"encoding/json"
	"io"
	"net/http"
)

func ReadJSON[T any](r io.Reader, dst *T) error {
	if err := json.NewDecoder(r).Decode(dst); err != nil {
		return err
	}

	return nil
}

func WriteJSON[T any](w http.ResponseWriter, v T, status int) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return err
	}

	return nil
}
