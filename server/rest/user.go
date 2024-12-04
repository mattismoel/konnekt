package rest

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mattismoel/konnekt"
	"github.com/mattismoel/konnekt/internal/password"
)

type UserService interface {
	FindUserByID(context.Context, int64) (konnekt.User, error)
	FindUsers(context.Context, konnekt.UserFilter) ([]konnekt.User, error)

	CreateUser(context.Context, konnekt.User, password.Password, password.Password) (konnekt.User, error)

	UpdateUser(context.Context, int64, konnekt.UpdateUser) (konnekt.User, error)
	DeleteUser(context.Context, int64) error
}

func (s server) createUserRoutes() http.Handler {
	r := chi.NewRouter()

	r.Post("/", s.handleCreateUser())

	return r
}

type CreateUserLoad struct {
	konnekt.User
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

func (s server) handleCreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var load CreateUserLoad

		err := readJSON(r, &load)
		if err != nil {
			Error(w, r, err)
			return
		}

		user, err := s.userService.CreateUser(r.Context(), konnekt.User{
			Email:     load.Email,
			FirstName: load.FirstName,
			LastName:  load.LastName,
		}, []byte(load.Password), []byte(load.PasswordConfirm))

		if err != nil {
			Error(w, r, err)
			return
		}

		writeJSON(w, http.StatusOK, user)
	}
}
