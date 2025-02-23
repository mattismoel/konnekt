package server

import (
	"net/http"

	"github.com/mattismoel/konnekt/internal/service"
)

func (s Server) handleListVenues() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		result, err := s.venueService.List(ctx, service.VenueListQuery{
			ListQuery: NewListQueryFromRequest(r),
		})

		if err != nil {
			writeError(w, err)
			return
		}

		writeJSON(w, http.StatusOK, result)
	}
}
