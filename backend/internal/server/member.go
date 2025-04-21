package server

import (
	"net/http"
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
