package server

import (
	"net/http"

	"github.com/mattismoel/konnekt/backend/api"
	"github.com/mattismoel/konnekt/backend/urlutil"
)

func (s Server) handleGetTeamPermissions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		teamID, err := urlutil.PathInt(r, "teamID")
		if err != nil {
			api.WriteError(w, r, api.BadRequestError(r, "Invalid ID"))
			return
		}

		result, err := s.AuthRepo.TeamPermissions(ctx, int64(teamID))
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
