package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/mattismoel/konnekt/internal/domain/concert"
	"github.com/mattismoel/konnekt/internal/domain/event"
)

const COVER_IMAGE_WIDTH = 2048

var (
	ErrEventNoExist = APIError{Message: "Event does not exist", Status: http.StatusNotFound}
)

func (s Server) handleListEvents() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		query, err := NewListQueryFromURL(r.URL.Query())
		if err != nil {
			writeError(w, err)
			return
		}

		result, err := s.eventService.List(ctx, query)
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

		eventID, err := paramID("eventID", r)
		if err != nil {
			writeError(w, err)
			return
		}

		e, err := s.eventService.ByID(ctx, eventID)
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
		Title       string              `json:"title"`
		Description string              `json:"description"`
		ImageURL    string              `json:"imageUrl"`
		TicketURL   string              `json:"ticketUrl"`
		VenueID     int64               `json:"venueId"`
		Concerts    []createConcertLoad `json:"concerts"`
		IsPublic    bool                `json:"isPublic"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var load createEventLoad

		ctx := r.Context()

		requestMember, err := s.memberFromRequest(ctx, w, r)
		if err != nil {
			writeError(w, err)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&load); err != nil {
			writeError(w, err)
			return
		}

		concerts := make([]concert.Concert, 0)

		for _, loadConcert := range load.Concerts {
			a, err := s.artistService.ByID(ctx, loadConcert.ArtistID)
			if err != nil {
				writeError(w, err)
				return
			}

			c, err := concert.NewConcert(
				concert.WithArtist(a),
				concert.WithFrom(loadConcert.From),
				concert.WithTo(loadConcert.To),
			)

			if err != nil {
				writeError(w, err)
				return
			}

			fmt.Printf("CONCERT TO INSERT: %+v\n", c)

			concerts = append(concerts, c)
		}

		v, err := s.venueService.ByID(ctx, load.VenueID)
		if err != nil {
			writeError(w, err)
			return
		}

		e, err := event.NewEvent(
			event.WithTitle(load.Title),
			event.WithDescription(load.Description),
			event.WithTicketURL(load.TicketURL),
			event.WithVenue(v),
			event.WithImageURL(load.ImageURL),
			event.WithConcerts(concerts...),
			event.WithIsPublic(load.IsPublic),
			event.WithCreatedBy(requestMember.ID),
			event.WithUpdatedBy(requestMember.ID),
		)

		if err != nil {
			writeError(w, err)
			return
		}

		eventID, err := s.eventService.Create(ctx, e)
		if err != nil {
			writeError(w, err)
			return
		}

		createdEvent, err := s.eventService.ByID(ctx, eventID)
		if err != nil {
			writeError(w, err)
			return
		}

		writeJSON(w, http.StatusOK, createdEvent)
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
		TicketURL   string              `json:"ticketURL"`
		ImageURL    string              `json:"imageUrl"`
		Concerts    []updateConcertLoad `json:"concerts"`
		VenueID     int64               `json:"venueId"`
		IsPublic    bool                `json:"isPublic"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		eventID, err := paramID("eventID", r)
		if err != nil {
			writeError(w, err)
			return
		}

		var load updateEventLoad

		err = json.NewDecoder(r.Body).Decode(&load)
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

		// concerts := make([]service.UpdateConcert, 0)
		//
		// for _, c := range load.Concerts {
		// 	concerts = append(concerts, service.UpdateConcert{
		// 		ArtistID: c.ArtistID,
		// 		From:     c.From,
		// 		To:       c.To,
		// 	})
		// }

		// Return if event does not exist.
		prevEvent, err := s.eventService.ByID(ctx, eventID)
		if err != nil {
			writeError(w, err)
			return
		}

		venue, err := s.venueService.ByID(ctx, load.VenueID)
		if err != nil {
			writeError(w, err)
			return
		}

		concerts := make([]concert.Concert, 0)
		for _, loadConcert := range load.Concerts {
			artist, err := s.artistService.ByID(ctx, loadConcert.ArtistID)
			if err != nil {
				writeError(w, err)
				return
			}

			c, err := concert.NewConcert(
				concert.WithID(eventID),
				concert.WithArtist(artist),
				concert.WithFrom(loadConcert.From),
				concert.WithTo(loadConcert.To),
			)

			if err != nil {
				writeError(w, err)
				return
			}

			concerts = append(concerts, c)
		}

		e, err := event.NewEvent(
			event.WithID(eventID),
			event.WithTitle(load.Title),
			event.WithDescription(load.Description),
			event.WithTicketURL(load.TicketURL),
			event.WithConcerts(concerts...),
			event.WithVenue(venue),
			event.WithIsPublic(load.IsPublic),
			event.WithUpdatedBy(requestMember.ID),
		)

		if err != nil {
			writeError(w, err)
			return
		}

		// If there is a cover image URL update, set it.
		if strings.TrimSpace(load.ImageURL) != "" {
			if err := s.objectStore.Delete(ctx, prevEvent.ImageURL); err != nil {
				writeError(w, err)
				return
			}

			if err := e.WithCfgs(event.WithImageURL(load.ImageURL)); err != nil {
				writeError(w, err)
				return
			}
		}

		e, err = s.eventService.Update(ctx, eventID, e)
		if err != nil {
			writeError(w, err)
			return
		}

		writeJSON(w, http.StatusOK, e)
	}
}

func (s Server) handleUploadEventImage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("image")
		if err != nil {
			writeError(w, err)
			return
		}

		defer file.Close()

		ctx := r.Context()

		url, err := s.eventService.UploadImage(ctx, file)
		if err != nil {
			writeError(w, err)
			return
		}

		writeText(w, http.StatusOK, url)
	}
}

func (s Server) handleDeleteEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		eventID, err := paramID("eventID", r)
		if err != nil {
			writeError(w, err)
			return
		}

		err = s.eventService.Delete(ctx, eventID)
		if err != nil {
			writeError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
