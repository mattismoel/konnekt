package server

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (s Server) handleListRoles() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		query, err := NewListQueryFromURL(r.URL.Query())
		if err != nil {
			writeError(w, err)
			return
		}

		result, err := s.authService.ListRoles(ctx, query)
		if err != nil {
			writeError(w, err)
			return
		}

		writeJSON(w, http.StatusOK, result)
	}
}
