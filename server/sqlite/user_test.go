package sqlite_test

import (
	"context"
	"testing"

	"github.com/mattismoel/konnekt"
	"github.com/mattismoel/konnekt/internal/password"
	"github.com/mattismoel/konnekt/sqlite"
)

type userUpdater func(konnekt.User) konnekt.User

var baseUser = konnekt.User{
	ID:        1,
	Email:     "test@mail.com",
	FirstName: "John",
	LastName:  "Doe",
}

func TestCreateUser(t *testing.T) {
	type test struct {
		updater         userUpdater
		password        password.Password
		passwordConfirm password.Password
		errCode         string
	}

	tests := map[string]test{
		"Valid load": {
			updater:         nil,
			password:        []byte("Password123!"),
			passwordConfirm: []byte("Password123!"),
			errCode:         "",
		},
		"Invalid email": {
			updater: func(u konnekt.User) konnekt.User {
				u.Email = "invalid-email.com"
				return u
			},
			password:        []byte("Password123!"),
			passwordConfirm: []byte("Password123!"),
			errCode:         konnekt.ERRINVALID,
		},
		"No email": {
			updater: func(u konnekt.User) konnekt.User {
				u.Email = " "
				return u
			},
			password:        []byte("Password123!"),
			passwordConfirm: []byte("Password123!"),
			errCode:         konnekt.ERRINVALID,
		},
		"No First Name": {
			updater: func(u konnekt.User) konnekt.User {
				u.FirstName = " "
				return u
			},
			password:        []byte("Password123!"),
			passwordConfirm: []byte("Password123!"),
			errCode:         konnekt.ERRINVALID,
		},
		"No Last Name": {
			updater: func(u konnekt.User) konnekt.User {
				u.LastName = " "
				return u
			},
			password:        []byte("Password123!"),
			passwordConfirm: []byte("Password123!"),
			errCode:         konnekt.ERRINVALID,
		},
		"Empty password": {
			updater:         nil,
			password:        []byte(""),
			passwordConfirm: []byte("Password123!"),
			errCode:         konnekt.ERRINVALID,
		},
		"Non-matching passwords": {
			updater:         nil,
			password:        []byte("password!!!!"),
			passwordConfirm: []byte("Password123!"),
			errCode:         konnekt.ERRINVALID,
		},
		"Nil password": {
			updater:         nil,
			password:        nil,
			passwordConfirm: []byte("Password123!"),
			errCode:         konnekt.ERRINVALID,
		},
		"Empty passwords": {
			updater:         nil,
			password:        []byte(""),
			passwordConfirm: []byte(""),
			errCode:         konnekt.ERRINVALID,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			repo, dsn := MustOpenRepo(t)
			defer MustCloseRepo(t, repo, dsn)

			service := sqlite.NewUserService(repo)

			user := baseUser
			if tt.updater != nil {
				user = tt.updater(user)
			}

			_, err := service.CreateUser(context.Background(), user, tt.password, tt.passwordConfirm)

			code := konnekt.ErrorCode(err)

			if code != tt.errCode {
				t.Fatalf("got %q, want %q, error: %v", code, tt.errCode, err)
			}
		})
	}
}
