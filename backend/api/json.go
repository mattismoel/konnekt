package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ReadJSON[T any](r io.Reader, dst *T) error {
	if err := json.NewDecoder(r).Decode(dst); err != nil {
		return fmt.Errorf("Could not read JSON body: %v", err)
	}

	return nil
}

func WriteJSON[T any](w http.ResponseWriter, v T, status int) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("Could not write JSON: %v", err)
	}

	return nil
}
