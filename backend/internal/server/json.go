package server

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
