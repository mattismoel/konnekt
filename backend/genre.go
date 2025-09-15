package konnekt

import (
	"net/http"

	"github.com/mattismoel/konnekt/backend/api"
)

type Genre string

type CreateGenre struct {
	Name      string `json:"name"`

func (s Server) handleCreateGenre() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var load CreateGenre
		if err := api.ReadJSON(r.Body, &load); err != nil {
			api.WriteError(w, r, err)
			return
		}

		load.CreatedBy = 1

		genreID, err := s.ArtistRepo.InsertGenre(ctx, load)
		if err != nil {
			api.WriteError(w, r, err)
			return
		}

		genre, err := s.ArtistRepo.GenreByID(ctx, genreID)
		if err != nil {
			api.WriteError(w, r, err)
			return
		}

		if err := api.WriteJSON(w, genre, int(http.StatusCreated)); err != nil {
			api.WriteError(w, r, err)
			return
		}
	}
	CreatedBy int64  `json:"-"`
}
