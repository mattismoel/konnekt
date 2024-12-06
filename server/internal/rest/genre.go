package rest

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mattismoel/konnekt/internal/service"
)

type GenreService interface {
	CreateGenre(ctx context.Context, name string) (service.Genre, error)
	FindGenres(ctx context.Context, filter service.GenreFilter) ([]service.Genre, error)
	DeleteGenre(ctx context.Context, id int64) error
}

func (s server) createGenreRoutes() http.Handler {
	r := chi.NewRouter()

	r.Get("/", s.handleGetGenres())
	r.Post("/", s.handleCreateGenre())
	r.Delete("/{id}", s.handleDeleteGenre())

	return r
}

func (s server) handleGetGenres() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		genres, err := s.genreService.FindGenres(r.Context(), service.GenreFilter{})
		if err != nil {
			Error(w, r, err)
			return
		}

		writeJSON(w, http.StatusOK, genres)
	}
}

func (s server) handleCreateGenre() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var load struct {
			Name string `json:"name"`
		}

		err := readJSON(r, &load)
		if err != nil {
			Error(w, r, err)
			return
		}

		_, err = s.genreService.CreateGenre(r.Context(), load.Name)
		if err != nil {
			Error(w, r, err)
			return
		}
	}
}

func (s server) handleDeleteGenre() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			Error(w, r, err)
			return
		}

		err = s.genreService.DeleteGenre(r.Context(), int64(id))
		if err != nil {
			Error(w, r, err)
			return
		}
	}
}
