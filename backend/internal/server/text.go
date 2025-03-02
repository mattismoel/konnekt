package server

import "net/http"

func writeText(w http.ResponseWriter, status int, text string) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "text/plain")

	_, err := w.Write([]byte(text))
	if err != nil {
		return err
	}

	return nil
}
