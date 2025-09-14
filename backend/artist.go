package konnekt

import (
	"context"
	"errors"
	"net/http"

	"github.com/mattismoel/konnekt/backend/api"
	"github.com/mattismoel/konnekt/backend/mask"
	"github.com/mattismoel/konnekt/backend/urlutil"
)

type Artist struct {
	ID          ID       `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	ImageURL    string   `json:"imageUrl"`
	PreviewURL  string   `json:"previewUrl"`
	Socials     []Social `json:"social"`
	Genres      []Genre  `json:"genres"`
}

type CreateArtist struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	ImageURL    string   `json:"imageUrl"`
	PreviewURL  string   `json:"previewUrl"`
	Socials     []Social `json:"socials"`
	GenreIDs    []ID     `json:"genreIds"`
	CreatedBy   int64    `json:"-"`
}

type UpdateArtist struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	ImageURL    string   `json:"imageUrl"`
	PreviewURL  string   `json:"previewUrl"`
	Socials     []Social `json:"socials"`
	GenreIDs    []ID     `json:"genres"`
}

type ArtistRepo interface {
	InsertArtist(context.Context, CreateArtist) (ID, error)
	ArtistByID(context.Context, ID) (Artist, error)
	ListArtists(context.Context, api.ListRequest) (api.ListResponse[Artist], error)
	UpdateArtist(context.Context, ID, api.UpdateRequest[UpdateArtist]) error
	DeleteArtist(context.Context, ID) error

	InsertGenre(context.Context, CreateGenre) (ID, error)
	GenreByID(context.Context, ID) (Genre, error)
}

type Social string

func (s Server) handleCreateArtist() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var load CreateArtist
		if err := api.ReadJSON(r.Body, &load); err != nil {
			api.WriteError(w, r, err)
			return
		}

		load.CreatedBy = 1

		insertedID, err := s.ArtistRepo.InsertArtist(ctx, load)
		if err != nil {
			api.WriteError(w, r, err)
			return
		}

		artist, err := s.ArtistRepo.ArtistByID(ctx, insertedID)
		if err != nil {
			api.WriteError(w, r, err)
			return
		}

		if err := api.WriteJSON(w, artist, int(http.StatusOK)); err != nil {
			api.WriteError(w, r, err)
			return
		}
	}
}

func (s Server) handleGetArtistByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		artistID, err := urlutil.PathInt(r, "artistID")
		if err != nil {
			api.WriteError(w, r, err)
			return
		}

		artist, err := s.ArtistRepo.ArtistByID(ctx, ID(artistID))
		if err != nil {
			if errors.Is(err, ErrResourceNotFound) {
				api.WriteError(w, r, api.NotFoundError(r, "Could not find artist"))
				return
			}
			api.WriteError(w, r, err)
			return
		}

		if err := api.WriteJSON(w, artist, http.StatusOK); err != nil {
			api.WriteError(w, r, err)
			return
		}
	}
}

func (s Server) handleUpdateArtist() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		artistID, err := urlutil.PathInt(r, "artistID")
		if err != nil {
			api.WriteError(w, r, err)
			return
		}

		var load api.UpdateRequest[UpdateArtist]

		if err := api.ReadJSON(r.Body, &load); err != nil {
			api.WriteError(w, r, err)
			return
		}

		if err := s.ArtistRepo.UpdateArtist(ctx, ID(artistID), load); err != nil {
			if errors.Is(err, ErrResourceNotFound) {
				api.WriteError(w, r, api.NotFoundError(r, "Artist not found"))
				return
			}

			api.WriteError(w, r, err)
			return
		}

		updatedArtist, err := s.ArtistRepo.ArtistByID(ctx, ID(artistID))
		if err != nil {
			api.WriteError(w, r, err)
			return
		}

		if err := api.WriteJSON(w, updatedArtist, int(http.StatusOK)); err != nil {
			api.WriteError(w, r, err)
			return
		}
	}
}

func (s Server) handleListArtists() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		lr := api.NewListRequest(r)

		result, err := s.ArtistRepo.ListArtists(ctx, lr)
		if err != nil {
			api.WriteError(w, r, err)
			return
		}

		if err := api.WriteJSON(w, result, int(http.StatusOK)); err != nil {
			api.WriteError(w, r, err)
			return
		}
	}
}

func (s Server) handleDeleteArtist() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		artistID, err := urlutil.PathInt(r, "artistID")
		if err != nil {
			api.WriteError(w, r, api.BadRequestError(r, "Invalid artist ID path value"))
			return
		}

		if err := s.ArtistRepo.DeleteArtist(ctx, ID(artistID)); err != nil {
			api.WriteError(w, r, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (ua UpdateArtist) Fields() mask.FieldMap {
	return mask.FieldMap{
		"name":        ua.Name,
		"description": ua.Description,
		"image_url":   ua.ImageURL,
		"preview_url": ua.PreviewURL,
		"socials":     ua.Socials,
		"genres":      ua.GenreIDs,
	}
}
