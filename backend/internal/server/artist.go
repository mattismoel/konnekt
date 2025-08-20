package server

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/mattismoel/konnekt/internal/domain/artist"
	"github.com/mattismoel/konnekt/internal/query"
	"github.com/mattismoel/konnekt/internal/service"
)

func (s Server) handleListArtists() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		q, err := NewListQueryFromURL(r.URL.Query())
		if err != nil {
			writeError(w, err)
			return
		}

		result, err := s.artistService.List(ctx, q)

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

		artistID, err := paramID("artistID", r)
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

		requestMember, err := s.memberFromRequest(ctx, w, r)
		if err != nil {
			writeError(w, err)
			return
		}

		filters := make(query.FilterCollection)
		for _, genreID := range load.GenreIDs {
			err := filters.Add("id", query.Filter{
				Cmp:   query.Equal,
				Value: strconv.Itoa(int(genreID)),
			})

			if err != nil {
				writeError(w, err)
				return
			}

		}

		q, err := query.NewListQuery(query.WithFilters(filters))
		if err != nil {
			writeError(w, err)
			return
		}

		genresResult, err := s.artistService.ListGenres(ctx, q)
		if err != nil {
			writeError(w, err)
			return
		}

		socials := make([]artist.Social, 0)
		for _, social := range load.Socials {
			s, err := artist.NewSocial(social)
			if err != nil {
				writeError(w, err)
				return
			}

			socials = append(socials, s)
		}

		a, err := artist.NewArtist(
			artist.WithName(load.Name),
			artist.WithDescription(load.Description),
			artist.WithGenres(genresResult.Records...),
			artist.WithImageURL(load.ImageURL),
			artist.WithSocials(socials...),
			artist.WithCreatedBy(requestMember.ID),
			artist.WithUpdatedBy(requestMember.ID),
		)

		if err != nil {
			writeError(w, err)
			return
		}

		if strings.TrimSpace(load.PreviewURL) != "" {
			err := a.WithCfgs(artist.WithPreviewURL(load.PreviewURL))
			if err != nil {
				writeError(w, err)
				return
			}
		}

		artistID, err := s.artistService.Create(ctx, a)
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

		artistID, err := paramID("artistID", r)
		if err != nil {
			writeError(w, err)
			return
		}

		err = s.artistService.Update(ctx, artistID, service.UpdateArtist{
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

		updatedArtist, err := s.artistService.ByID(ctx, artistID)
		if err != nil {
			writeError(w, err)
			return
		}

		writeJSON(w, http.StatusOK, updatedArtist)
	}
}

func (s Server) handleUploadArtistImage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		file, _, err := r.FormFile("image")
		if err != nil {
			writeError(w, err)
			return
		}

		defer file.Close()

		url, err := s.artistService.UploadImage(ctx, file)
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

		artistID, err := paramID("artistID", r)
		if err != nil {
			writeError(w, err)
			return
		}

		err = s.artistService.Delete(ctx, artistID)
		if err != nil {
			writeError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (s Server) handleListGenres() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		q, err := NewListQueryFromURL(r.URL.Query())
		if err != nil {
			writeError(w, err)
			return
		}

		result, err := s.artistService.ListGenres(ctx, q)
		if err != nil {
			writeError(w, err)
			return
		}

		writeJSON(w, http.StatusOK, result)
	}
}
