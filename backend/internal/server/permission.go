package server

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (s Server) handleListPermissions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		q, err := NewListQueryFromURL(r.URL.Query())
		if err != nil {
			writeError(w, err)
			return
		}

		result, err := s.authService.ListPermissions(ctx, q)
		if err != nil {
			writeError(w, err)
			return
		}

		writeJSON(w, http.StatusOK, result)
	}
}

func (s Server) handleListMemberPermissions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		memberID, err := strconv.Atoi(chi.URLParam(r, "memberID"))
		if err != nil {
			writeError(w, err)
			return
		}

		perms, err := s.authService.MemberPermissions(ctx, int64(memberID))
		if err != nil {
			writeError(w, err)
			return
		}

		writeJSON(w, http.StatusOK, perms)
	}
}
