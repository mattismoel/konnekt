package server

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (s Server) handleListUserRoles() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userID, err := strconv.Atoi(chi.URLParam(r, "userID"))
		if err != nil {
			writeError(w, err)
			return
		}

		roles, err := s.authService.UserRoles(ctx, int64(userID))
		if err != nil {
			writeError(w, err)
			return
		}

		writeJSON(w, http.StatusOK, roles)
	}
}
