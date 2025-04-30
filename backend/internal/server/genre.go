package server

import (
	"encoding/json"
	"net/http"
)

func (s Server) handleCreateGenre() http.HandlerFunc {
	type createGenreLoad struct {
		Name string `json:"name"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var load createGenreLoad

		err := json.NewDecoder(r.Body).Decode(&load)
		if err != nil {
			writeError(w, err)
			return
		}

		ctx := r.Context()

		_, err = s.artistService.CreateGenre(ctx, load.Name)
		if err != nil {
			writeError(w, err)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
