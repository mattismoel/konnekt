package sqlite_test

import (
	"context"
	"testing"

	"github.com/mattismoel/konnekt"
	"github.com/mattismoel/konnekt/internal/password"
	"github.com/mattismoel/konnekt/sqlite"
)

var baseUser = konnekt.User{
	ID:        1,
	Email:     "test@mail.com",
	FirstName: "John",
	LastName:  "Doe",
}

func TestCreateUser(t *testing.T) {
	type test struct {
		u               konnekt.User
		password        password.Password
		passwordConfirm password.Password
		errCode         string
	}

	tests := map[string]test{
		"Valid load": {
			u:               baseUser,
			password:        []byte("Password123!"),
			passwordConfirm: []byte("Password123!"),
			errCode:         "",
		},
		"Invalid email": {
			u: konnekt.User{
				Email:     "invalid-email",
				FirstName: baseUser.FirstName,
				LastName:  baseUser.LastName,
			},
			password:        []byte("Password123!"),
			passwordConfirm: []byte("Password123!"),
			errCode:         konnekt.ERRINVALID,
		},
		"No email": {
			u: konnekt.User{
				Email:     "",
				FirstName: baseUser.FirstName,
				LastName:  baseUser.LastName,
			},
			password:        []byte("Password123!"),
			passwordConfirm: []byte("Password123!"),
			errCode:         konnekt.ERRINVALID,
		},
		"No First Name": {
			u: konnekt.User{
				Email:     baseUser.Email,
				FirstName: "",
				LastName:  baseUser.LastName,
			},
			password:        []byte("Password123!"),
			passwordConfirm: []byte("Password123!"),
			errCode:         konnekt.ERRINVALID,
		},
		"No Last Name": {
			u: konnekt.User{
				Email:     baseUser.Email,
				FirstName: baseUser.FirstName,
				LastName:  "",
			},
			password:        []byte("Password123!"),
			passwordConfirm: []byte("Password123!"),
			errCode:         konnekt.ERRINVALID,
		},
		"Empty password": {
			u:               baseUser,
			password:        []byte(""),
			passwordConfirm: []byte("Password123!"),
			errCode:         konnekt.ERRINVALID,
		},
		"Non-matching passwords": {
			u:               baseUser,
			password:        []byte("password!!!!"),
			passwordConfirm: []byte("Password123!"),
			errCode:         konnekt.ERRINVALID,
		},
		"Nil password": {
			u:               baseUser,
			password:        nil,
			passwordConfirm: []byte("Password123!"),
			errCode:         konnekt.ERRINVALID,
		},
		"Empty passwords": {
			u:               baseUser,
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

			_, err := service.CreateUser(context.Background(), tt.u, tt.password, tt.passwordConfirm)

			code := konnekt.ErrorCode(err)

			if code != tt.errCode {
				t.Fatalf("got %q, want %q, error: %v", code, tt.errCode, err)
			}
		})
	}
}
