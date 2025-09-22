package api

import (
	"fmt"
	"net/http"
)

func WriteText(w http.ResponseWriter, txt string, status int) error {
	w.Header().Set("Content-Type", "text/plain")
	if _, err := w.Write([]byte(txt)); err != nil {
		return fmt.Errorf("Could not write text content: %v", err)
	}

	w.WriteHeader(status)
	return nil
}
