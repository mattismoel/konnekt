package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mattismoel/konnekt/internal/domain/artist"
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
		PreviewURL  string   `json:"previewUrl"`
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

		artistID, err := s.artistService.Create(ctx, service.CreateArtist{
			Name:        load.Name,
			Description: load.Description,
			ImageURL:    load.ImageURL,
			PreviewURL:  load.PreviewURL,
			GenreIDs:    load.GenreIDs,
			Socials:     load.Socials,
		})

		if err != nil {
			writeError(w, err)
			return
		}

		artist, err := s.artistService.ByID(ctx, artistID)
		if err != nil {
			writeError(w, err)
			return
		}

		writeJSON(w, http.StatusOK, artist)
	}
}

func (s Server) handleUpdateArtist() http.HandlerFunc {
	type updateArtistLoad struct {
		Name        string   `json:"name"`
		Description string   `json:"description"`
		ImageURL    string   `json:"imageUrl"`
		PreviewURL  string   `json:"previewUrl"`
		GenreIDs    []int64  `json:"genreIds"`
		Socials     []string `json:"socials"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var load updateArtistLoad

		err := json.NewDecoder(r.Body).Decode(&load)
		if err != nil {
			writeError(w, err)
			return
		}

		ctx := r.Context()

		artistID, err := strconv.Atoi(chi.URLParam(r, "artistID"))
		if err != nil {
			writeError(w, err)
			return
		}

		a, err := s.artistService.Update(ctx, int64(artistID), service.UpdateArtist{
			Name:        load.Name,
			Description: load.Description,
			PreviewURL:  load.PreviewURL,
			ImageURL:    load.ImageURL,
			GenreIDs:    load.GenreIDs,
			Socials:     load.Socials,
		})

		if err != nil {
			writeError(w, err)
			return
		}

		err = a.WithCfgs(artist.WithID(int64(artistID)))
		if err != nil {
			writeError(w, err)
			return
		}

		writeJSON(w, http.StatusOK, a)
	}
}

func (s Server) handleUploadArtistImage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		file, fileHeader, err := r.FormFile("image")
		if err != nil {
			writeError(w, err)
			return
		}

		defer file.Close()

		url, err := s.artistService.UploadImage(ctx, fileHeader.Filename, file)
		if err != nil {
			writeError(w, err)
			return
		}

		writeText(w, http.StatusOK, url)
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
