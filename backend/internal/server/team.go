package server

import (
	"net/http"

	"github.com/mattismoel/konnekt/backend/api"
	"github.com/mattismoel/konnekt/backend/urlutil"
)

func (s Server) handleGetMemberTeams() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		memberID, err := urlutil.PathInt(r, "memberID")
		if err != nil {
			api.WriteError(w, r, api.BadRequestError(r, "Invalid ID"))
			return
		}

		result, err := s.AuthRepo.MemberTeams(ctx, int64(memberID))
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
