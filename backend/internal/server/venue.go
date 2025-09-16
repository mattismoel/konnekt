package server

import (
	"net/http"

	konnekt "github.com/mattismoel/konnekt/backend"
	"github.com/mattismoel/konnekt/backend/api"
)

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
