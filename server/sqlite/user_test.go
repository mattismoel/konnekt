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

func TestDeleteUser(t *testing.T) {
	type test struct {
		id       int64
		wantCode string
	}

	tests := map[string]test{
		"Valid ID": {
			id:       1,
			wantCode: "",
		},
		"Invalid ID": {
			id:       999,
			wantCode: "",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			repo, dsn := MustOpenRepo(t)
			defer MustCloseRepo(t, repo, dsn)

			MustCreateUser(t, context.Background(), repo, baseUser,
				[]byte("Password123!"), []byte("Password123!"))

			service := sqlite.NewUserService(repo)

			err := service.DeleteUser(context.Background(), tt.id)
			code := konnekt.ErrorCode(err)
			if code != tt.wantCode {
				t.Fatalf("got code %q, want code %q, error: %v", code, tt.wantCode, err)
			}

			_, err = service.FindUsers(context.Background(), konnekt.UserFilter{Email: &baseUser.Email})
			if err != nil && konnekt.ErrorCode(err) != konnekt.ERRNOTFOUND {
				t.Fatal(err)
			}
		})
	}
}

func TestFindUserByID(t *testing.T) {
	type test struct {
		id       int64
		wantUser konnekt.User
		wantCode string
	}

	tests := map[string]test{
		"Valid ID": {
			id:       1,
			wantUser: baseUser,
			wantCode: "",
		},
		"Invalid ID": {
			id:       999,
			wantUser: konnekt.User{},
			wantCode: konnekt.ERRNOTFOUND,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			repo, dsn := MustOpenRepo(t)
			defer MustCloseRepo(t, repo, dsn)

			service := sqlite.NewUserService(repo)

			MustCreateUser(t, context.Background(), repo, baseUser, []byte("Password123!"), []byte("Password123!"))

			user, err := service.FindUserByID(context.Background(), tt.id)

			code := konnekt.ErrorCode(err)
			if code != tt.wantCode {
				t.Fatalf("got code %q, want code %q, error: %v", code, tt.wantCode, err)
			}

			if !user.Equals(tt.wantUser) {
				t.Fatalf("got %+v, want %+v", user, tt.wantUser)
			}
		})
	}
}

func MustCreateUser(t testing.TB, ctx context.Context, repo *sqlite.Repository, user konnekt.User, password []byte, passwordConfirm []byte) {
	t.Helper()

	service := sqlite.NewUserService(repo)
	_, err := service.CreateUser(ctx, user, password, passwordConfirm)
	if err != nil {
		t.Fatal(err)
	}
}
