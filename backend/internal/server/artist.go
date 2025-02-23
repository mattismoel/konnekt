package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/mattismoel/konnekt/internal/service"
)

func (s Server) handleListArtists() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		page, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil || page <= 0 {
			page = DEFAULT_PAGE
		}

		perPage, err := strconv.Atoi(r.URL.Query().Get("perPage"))
		if err != nil || perPage <= 0 {
			perPage = DEFAULT_PER_PAGE
		}

		if perPage > MAX_PER_PAGE {
			perPage = MAX_PER_PAGE
		}

		result, err := s.artistService.List(ctx, service.ArtistQuery{
			Page:    page,
			PerPage: perPage,
		})

		if err != nil {
			writeError(w, err)
			return
		}

		writeJSON(w, http.StatusOK, ListReponse{
			Page:       result.Page,
			PerPage:    result.PerPage,
			TotalCount: result.Page,
			PageCount:  result.PageCount,
			Records:    result.Artists,
		})
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
