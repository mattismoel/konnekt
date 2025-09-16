package server

import (
	"context"
	"net/http"

	konnekt "github.com/mattismoel/konnekt/backend"
	"github.com/mattismoel/konnekt/backend/api"
	"github.com/mattismoel/konnekt/backend/urlutil"
)

type VenueRepo interface {
	InsertVenue(context.Context, konnekt.CreateVenue) (int64, error)
	VenueByID(context.Context, int64) (konnekt.Venue, error)
	ListVenues(context.Context, api.ListRequest) (api.ListResponse[konnekt.Venue], error)
	UpdateVenue(context.Context, int64, api.UpdateRequest[konnekt.UpdateVenue]) error
}

func (s Server) handleCreateVenue() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		requestMemberID, err := s.requestMemberID(r)
		if err != nil {
			api.WriteError(w, r, api.UnauthorisedError(r))
			return
		}

		var load konnekt.CreateVenue
		if err := api.ReadJSON(r.Body, &load); err != nil {
			api.WriteError(w, r, err)
			return
		}

		load.CreatedBy = requestMemberID

		venueID, err := s.VenueRepo.InsertVenue(ctx, load)
		if err != nil {
			api.WriteError(w, r, err)
			return
		}

		venue, err := s.VenueRepo.VenueByID(ctx, venueID)
		if err != nil {
			api.WriteError(w, r, err)
			return
		}

		if err := api.WriteJSON(w, venue, int(http.StatusCreated)); err != nil {
			api.WriteError(w, r, err)
			return
		}
	}
}

func (s Server) handleListVenues() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		lr := api.NewListRequest(r)
		result, err := s.VenueRepo.ListVenues(ctx, lr)
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

func (s Server) handleUpdateVenue() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		venueID, err := urlutil.PathInt(r, "venueID")
		if err != nil {
			api.WriteError(w, r, api.BadRequestError(r, "Invalid venue ID"))
			return
		}

		var load api.UpdateRequest[konnekt.UpdateVenue]
		if err := api.ReadJSON(r.Body, &load); err != nil {
			api.WriteError(w, r, api.BadRequestError(r, "Invalid update body"))
			return
		}

		if err := s.VenueRepo.UpdateVenue(ctx, int64(venueID), load); err != nil {
			api.WriteError(w, r, err)
			return
		}

		updatedVenue, err := s.VenueRepo.VenueByID(ctx, int64(venueID))
		if err != nil {
			api.WriteError(w, r, err)
			return
		}

		if err := api.WriteJSON(w, updatedVenue, http.StatusOK); err != nil {
			api.WriteError(w, r, err)
			return
		}
	}
}
