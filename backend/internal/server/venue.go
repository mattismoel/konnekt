package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mattismoel/konnekt/internal/domain/venue"
	"github.com/mattismoel/konnekt/internal/service"
)

func (s Server) handleListVenues() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		baseQuery, err := NewListQueryFromURL(r.URL.Query())
		if err != nil {
			writeError(w, err)
			return
		}

		result, err := s.venueService.List(ctx, venue.Query{
			ListQuery: baseQuery,
		})

		if err != nil {
			writeError(w, err)
			return
		}

		writeJSON(w, http.StatusOK, result)
	}
}

func (s Server) handleCreateVenue() http.HandlerFunc {
	type createVenueLoad struct {
		Name        string `json:"name"`
		CountryCode string `json:"countryCode"`
		City        string `json:"city"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var load createVenueLoad
		err := json.NewDecoder(r.Body).Decode(&load)
		if err != nil {
			writeError(w, err)
			return
		}

		venueID, err := s.venueService.Create(ctx, service.CreateVenue{
			Name:        load.Name,
			CountryCode: load.CountryCode,
			City:        load.City,
		})

		if err != nil {
			writeError(w, err)
			return
		}

		venue, err := s.venueService.ByID(ctx, venueID)
		if err != nil {
			writeError(w, err)
			return
		}

		if err := writeJSON(w, http.StatusCreated, venue); err != nil {
			writeError(w, err)
			return
		}
	}
}

func (s Server) handleDeleteVenue() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		venueID, err := strconv.Atoi(chi.URLParam(r, "venueID"))
		if err != nil {
			writeError(w, err)
			return
		}

		err = s.venueService.Delete(ctx, int64(venueID))
		if err != nil {
			writeError(w, err)
			return
		}
	}
}

func (s Server) handleUpdateVenue() http.HandlerFunc {
	type UpdateVenueLoad struct {
		Name        string `json:"name"`
		City        string `json:"city"`
		CountryCode string `json:"countryCode"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var load UpdateVenueLoad

		err := json.NewDecoder(r.Body).Decode(&load)
		if err != nil {
			writeError(w, err)
			return
		}

		venueID, err := strconv.Atoi(chi.URLParam(r, "venueID"))
		if err != nil {
			writeError(w, err)
			return
		}

		venue, err := s.venueService.Update(ctx, int64(venueID), service.UpdateVenue{
			Name:        load.Name,
			City:        load.City,
			CountryCode: load.CountryCode,
		})

		if err != nil {
			writeError(w, err)
			return
		}

		writeJSON(w, http.StatusOK, venue)
	}
}

func (s Server) handleVenueByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		venueID, err := paramID("venueID", r)
		if err != nil {
			writeError(w, err)
			return
		}

		ctx := r.Context()

		v, err := s.venueService.ByID(ctx, venueID)
		if err != nil {
			writeError(w, err)
			return
		}

		if err := writeJSON(w, http.StatusOK, v); err != nil {
			writeError(w, err)
			return
		}
	}
}
