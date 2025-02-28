package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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

func (s Server) handleGetArtistByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		artistID, err := strconv.Atoi(chi.URLParam(r, "artistID"))
		if err != nil {
			writeError(w, err)
			return
		}

		artist, err := s.artistService.ByID(ctx, int64(artistID))
		if err != nil {
			writeError(w, err)
			return
		}

		writeJSON(w, http.StatusOK, artist)
	}
}

func (s Server) handleCreateArtist() http.HandlerFunc {
	type createArtistLoad struct {
		Name        string   `json:"name"`
		Description string   `json:"description"`
		ImageURL    string   `json:"imageUrl"`
		GenreIDs    []int64  `json:"genreIds"`
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
			Socials:     load.Socials,
		})

		if err != nil {
			writeError(w, err)
			return
		}
	}
}

func (s Server) handleDeleteArtist() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		artistID, err := strconv.Atoi(chi.URLParam(r, "artistID"))
		if err != nil {
			writeError(w, err)
			return
		}

		err = s.artistService.Delete(ctx, int64(artistID))
		if err != nil {
			writeError(w, err)
			return
		}
	}
}

func (s Server) handleListGenres() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		result, err := s.artistService.ListGenres(ctx, service.GenreListQuery{
			ListQuery: NewListQueryFromRequest(r),
		})

		if err != nil {
			writeError(w, err)
			return
		}

		writeJSON(w, http.StatusOK, result)
	}
}
