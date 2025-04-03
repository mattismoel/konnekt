package server

import "net/http"

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
