package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mattismoel/konnekt/internal/service"
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

func (s Server) handleCreateRole() http.HandlerFunc {
	type CreateRoleLoad struct {
		Name        string `json:"name"`
		DisplayName string `json:"displayName"`
		Description string `json:"description"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var load CreateRoleLoad

		err := json.NewDecoder(r.Body).Decode(&load)
		if err != nil {
			writeError(w, err)
			return
		}

		role, err := s.authService.CreateRole(ctx, service.CreateRole{
			Name:        load.Name,
			DisplayName: load.DisplayName,
			Description: load.Description,
		})

		if err != nil {
			writeError(w, err)
			return
		}

		writeJSON(w, http.StatusOK, role)
	}
}
