package server

import (
	"context"
	"net/http"

	konnekt "github.com/mattismoel/konnekt/backend"
	"github.com/mattismoel/konnekt/backend/api"
	"github.com/mattismoel/konnekt/backend/urlutil"
)

type EventRepo interface {
	InsertEvent(context.Context, konnekt.CreateEvent) (int64, error)
	EventByID(context.Context, int64) (konnekt.Event, error)
	ListEvents(context.Context, api.ListRequest) (api.ListResponse[konnekt.Event], error)
	Delete(context.Context, int64) error
	Update(context.Context, api.UpdateRequest[konnekt.UpdateEvent]) error
}

func (s Server) handleListEvents() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		lr := api.NewListRequest(r)

		result, err := s.EventRepo.ListEvents(ctx, lr)
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

func (s Server) handleCreateEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var load konnekt.CreateEvent
		if err := api.ReadJSON(r.Body, &load); err != nil {
			api.WriteError(w, r, api.BadRequestError(r, "Invalid request body"))
			return
		}

		eventID, err := s.EventRepo.InsertEvent(ctx, load)
		if err != nil {
			api.WriteError(w, r, err)
			return
		}

		createdEvent, err := s.EventRepo.EventByID(ctx, eventID)
		if err != nil {
			api.WriteError(w, r, err)
			return
		}

		if err := api.WriteJSON(w, createdEvent, int(http.StatusCreated)); err != nil {
			api.WriteError(w, r, err)
			return
		}
	}
}

func (s Server) handleGetEventByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		eventID, err := urlutil.PathInt(r, "eventID")
		if err != nil {
			api.WriteError(w, r, api.BadRequestError(r, "Inavlid event ID"))
			return
		}

		event, err := s.EventRepo.EventByID(ctx, int64(eventID))
		if err != nil {
			api.WriteError(w, r, err)
			return
		}

		if err := api.WriteJSON(w, event, int(http.StatusOK)); err != nil {
			api.WriteError(w, r, err)
			return
		}
	}
}
