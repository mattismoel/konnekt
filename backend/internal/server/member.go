package server

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (s Server) handleListMembers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query, err := NewListQueryFromURL(r.URL.Query())
		if err != nil {
			writeError(w, err)
			return
		}

		ctx := r.Context()

		result, err := s.memberService.List(ctx, query)
		if err != nil {
			writeError(w, err)
			return
		}

		err = writeJSON(w, http.StatusOK, result)
		if err != nil {
			writeError(w, err)
			return
		}
	}
}

func (s Server) handleApproveMember() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		memberID, err := strconv.Atoi(chi.URLParam(r, "memberID"))
		if err != nil {
			writeError(w, err)
			return
		}

		ctx := r.Context()

		err = s.memberService.Approve(ctx, int64(memberID))
		if err != nil {
			writeError(w, err)
			return
		}
	}
}
