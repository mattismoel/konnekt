package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/mattismoel/konnekt/internal/domain/event"
	"github.com/mattismoel/konnekt/internal/service"
)

var (
	ErrEventNoExist = APIError{Message: "Event does not exist", Status: http.StatusNotFound}
)

func (s Server) handleListEvents() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := NewListQueryFromRequest(r)

		fromStr, toStr := r.URL.Query().Get("from_date"), r.URL.Query().Get("to_date")
		from, _ := time.Parse(time.RFC3339, fromStr)
		to, _ := time.Parse(time.RFC3339, toStr)

		ctx := r.Context()

		result, err := s.eventService.List(ctx,
			service.NewEventListQuery(q.Page, q.PerPage, q.Limit, from, to),
		)

		if err != nil {
			writeError(w, err)
			return
		}

		writeJSON(w, http.StatusOK, result)
	}
}

func (s Server) handleEventByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		eventID, err := strconv.Atoi(chi.URLParam(r, "eventID"))
		if err != nil {
			writeError(w, err)
			return
		}

		e, err := s.eventService.ByID(ctx, int64(eventID))
		if err != nil {
			switch {
			case errors.Is(err, event.ErrNoExist):
				writeError(w, newAPIError(err.Error(), http.StatusNotFound))
			default:
				writeError(w, err)
			}
			return
		}

		writeJSON(w, http.StatusOK, e)
	}
}

func (s Server) handleCreateEvent() http.HandlerFunc {
	type createConcertLoad struct {
		ArtistID int64     `json:"artistID"`
		From     time.Time `json:"from"`
		To       time.Time `json:"to"`
	}

	type createEventLoad struct {
		Title         string              `json:"title"`
		Description   string              `json:"description"`
		CoverImageURL string              `json:"coverImageUrl"`
		VenueID       int64               `json:"venueId"`
		Concerts      []createConcertLoad `json:"concerts"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var load createEventLoad

		if err := json.NewDecoder(r.Body).Decode(&load); err != nil {
			writeError(w, err)
			return
		}

		concerts := make([]service.CreateConcert, 0)

		for _, conc := range load.Concerts {
			concerts = append(concerts, service.CreateConcert{
				ArtistID: conc.ArtistID,
				From:     conc.From,
				To:       conc.To,
			})
		}

		_, err := s.eventService.Create(r.Context(), service.CreateEvent{
			Title:       load.Title,
			Description: load.Description,
			VenueID:     load.VenueID,
			Concerts:    concerts,
		})

		if err != nil {
			writeError(w, err)
			return
		}
	}
}

func (s Server) handleUpdateEvent() http.HandlerFunc {
	type updateConcertLoad struct {
		ArtistID int64     `json:"artistId"`
		From     time.Time `json:"from"`
		To       time.Time `json:"to"`
	}

	type updateEventLoad struct {
		Title       string              `json:"title"`
		Description string              `json:"description"`
		Concerts    []updateConcertLoad `json:"concerts"`
		VenueID     int64               `json:"venueId"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s Server) handleSetEventCoverImage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		eventID, err := strconv.Atoi(chi.URLParam(r, "eventID"))
		if err != nil {
			writeError(w, err)
			return
		}

		ctx := r.Context()

		file, fileHeader, err := r.FormFile("image")
		if err != nil {
			writeError(w, err)
			return
		}

		url, err := s.eventService.SetCoverImage(ctx, int64(eventID), fileHeader.Filename, file)
		if err != nil {
			writeError(w, err)
			return
		}

		writeText(w, http.StatusOK, url)
	}
}
