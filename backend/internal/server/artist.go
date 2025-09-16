package server

import (
	"context"
	"errors"
	"net/http"

	konnekt "github.com/mattismoel/konnekt/backend"
	"github.com/mattismoel/konnekt/backend/api"
	"github.com/mattismoel/konnekt/backend/urlutil"
)

type ArtistRepo interface {
	InsertArtist(context.Context, konnekt.CreateArtist) (int64, error)
	ArtistByID(context.Context, int64) (konnekt.Artist, error)
	ListArtists(context.Context, api.ListRequest) (api.ListResponse[konnekt.Artist], error)
	UpdateArtist(context.Context, int64, api.UpdateRequest[konnekt.UpdateArtist]) error
	DeleteArtist(context.Context, int64) error

	InsertGenre(context.Context, konnekt.CreateGenre) (int64, error)
	GenreByID(context.Context, int64) (konnekt.Genre, error)
}

func (s Server) handleCreateArtist() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var load konnekt.CreateArtist
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

		artist, err := s.ArtistRepo.ArtistByID(ctx, int64(artistID))
		if err != nil {
			if errors.Is(err, konnekt.ErrResourceNotFound) {
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

		var load api.UpdateRequest[konnekt.UpdateArtist]

		if err := api.ReadJSON(r.Body, &load); err != nil {
			api.WriteError(w, r, err)
			return
		}

		if err := s.ArtistRepo.UpdateArtist(ctx, int64(artistID), load); err != nil {
			if errors.Is(err, konnekt.ErrResourceNotFound) {
				api.WriteError(w, r, api.NotFoundError(r, "Artist not found"))
				return
			}

			api.WriteError(w, r, err)
			return
		}

		updatedArtist, err := s.ArtistRepo.ArtistByID(ctx, int64(artistID))
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

		if err := s.ArtistRepo.DeleteArtist(ctx, int64(artistID)); err != nil {
			api.WriteError(w, r, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
