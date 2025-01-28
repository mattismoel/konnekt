package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/mattismoel/konnekt/internal/service"
)

func (s Server) handleListEvents() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		events, err := s.eventService.List(ctx)
		if err != nil {
			writeError(w, err)
			return
		}

		err = writeJSON(w, http.StatusOK, events)
		if err != nil {
			writeError(w, err)
			return
		}

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
		VenueID       int64               `json:"venueID"`
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
			Title:         load.Title,
			Description:   load.Description,
			CoverImageURL: load.CoverImageURL,
			VenueID:       load.VenueID,
			Concerts:      concerts,
		})

		if err != nil {
			switch {
			default:
				writeError(w, err)
			}
		}
	}
}
