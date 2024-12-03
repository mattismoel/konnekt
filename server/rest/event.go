package rest

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mattismoel/konnekt"
)

func (s server) createEventsRoutes() http.Handler {
	r := chi.NewRouter()

	r.Get("/", s.handleGetEvents())
	r.Get("/{id}", s.handleGetEventById())
	r.Post("/", s.handleCreateEvent())
	r.Put("/{id}", s.handleUpdateEvent())
	r.Delete("/{id}", s.handleDeleteEvent())

	return r
}

func (s server) handleGetEvents() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		events, err := s.eventService.FindEvents(r.Context(), konnekt.EventFilter{})
		if err != nil {
			Error(w, r, err)
			return
		}

		writeJSON(w, http.StatusOK, events)
	}
}

func (s server) handleGetEventById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (s server) handleCreateEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var load konnekt.Event

		err := readJSON(r, &load)
		if err != nil {
			Error(w, r, err)
		}

		_, err = s.eventService.CreateEvent(r.Context(), load)
		if err != nil {
			Error(w, r, err)
			return
		}
	}
}

func (s server) handleUpdateEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (s server) handleDeleteEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			Error(w, r, err)
			return
		}

		err = s.eventService.DeleteEvent(r.Context(), int64(id))
		if err != nil {
			Error(w, r, err)
			return
		}
	}
}
