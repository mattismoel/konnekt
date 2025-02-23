package server

import (
	"encoding/json"
	"net/http"

	"github.com/mattismoel/konnekt/internal/service"
)

func (s Server) handleListArtists() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		q := NewListQueryFromRequest(r)

		result, err := s.artistService.List(ctx, service.ArtistListQuery{
			ListQuery: q,
		})

		if err != nil {
			writeError(w, err)
			return
		}

		writeJSON(w, http.StatusOK, result)
	}
}

func (s Server) handleCreateArtist() http.HandlerFunc {
	type createArtistLoad struct {
		Name        string   `json:"name"`
		Description string   `json:"description"`
		ImageURL    string   `json:"imageUrl"`
		GenreIDs    []int64  `json:"genreIDs"`
		Socials     []string `json:"socials"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var load createArtistLoad

		err := json.NewDecoder(r.Body).Decode(&load)
		if err != nil {
			writeError(w, err)
			return
		}

		ctx := r.Context()

		_, err = s.artistService.Create(ctx, service.CreateArtist{
			Name:        load.Name,
			Description: load.Description,
			ImageURL:    load.ImageURL,
			GenreIDs:    load.GenreIDs,
		})

		if err != nil {
			writeError(w, err)
			return
		}
	}
}
