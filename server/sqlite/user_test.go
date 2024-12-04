package sqlite_test

import (
	"context"
	"testing"

	"github.com/mattismoel/konnekt"
	"github.com/mattismoel/konnekt/internal/password"
	"github.com/mattismoel/konnekt/internal/ptr"
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

func TestUpdateUser(t *testing.T) {
	baseUser := konnekt.User{
		ID:        1,
		Email:     "test@mail.com",
		FirstName: "John",
		LastName:  "Doe",
	}

	baseUpdate := konnekt.UpdateUser{
		Email:     ptr.From("new@mail.com"),
		FirstName: ptr.From("Sophie"),
		LastName:  ptr.From("Johnson"),
	}

	type userUpdateUpdater func(konnekt.UpdateUser) konnekt.UpdateUser

	type test struct {
		id              int64
		updater         userUpdateUpdater
		wantCode        string
		wantUserUpdater userUpdater
	}

	tests := map[string]test{
		"Valid load": {
			id:       1,
			updater:  nil,
			wantCode: "",
			wantUserUpdater: func(u konnekt.User) konnekt.User {
				u.Email = *baseUpdate.Email
				u.FirstName = *baseUpdate.FirstName
				u.LastName = *baseUpdate.LastName
				return u
			},
		},
		"Email update": {
			id: 1,
			updater: func(u konnekt.UpdateUser) konnekt.UpdateUser {
				return konnekt.UpdateUser{Email: baseUpdate.Email}
			},
			wantCode: "",
			wantUserUpdater: func(u konnekt.User) konnekt.User {
				u.Email = *baseUpdate.Email
				return u
			},
		},
		"First name update": {
			id: 1,
			updater: func(u konnekt.UpdateUser) konnekt.UpdateUser {
				return konnekt.UpdateUser{FirstName: baseUpdate.FirstName}
			},
			wantCode: "",
			wantUserUpdater: func(u konnekt.User) konnekt.User {
				u.FirstName = *baseUpdate.FirstName
				return u
			},
		},
		"Last name update": {
			id: 1,
			updater: func(u konnekt.UpdateUser) konnekt.UpdateUser {
				return konnekt.UpdateUser{LastName: baseUpdate.LastName}
			},
			wantCode: "",
			wantUserUpdater: func(u konnekt.User) konnekt.User {
				u.LastName = *baseUpdate.LastName
				return u
			},
		},
		"Empty update": {
			id: 1,
			updater: func(u konnekt.UpdateUser) konnekt.UpdateUser {
				return konnekt.UpdateUser{}
			},
			wantCode: "",
			wantUserUpdater: func(u konnekt.User) konnekt.User {
				return baseUser
			},
		},
		"Empty email": {
			id: 1,
			updater: func(u konnekt.UpdateUser) konnekt.UpdateUser {
				return konnekt.UpdateUser{Email: ptr.From("")}
			},
			wantCode: konnekt.ERRINVALID,
			wantUserUpdater: func(u konnekt.User) konnekt.User {
				return konnekt.User{}
			},
		},
		"Empty first name": {
			id: 1,
			updater: func(u konnekt.UpdateUser) konnekt.UpdateUser {
				return konnekt.UpdateUser{FirstName: ptr.From("")}
			},
			wantCode: konnekt.ERRINVALID,
			wantUserUpdater: func(u konnekt.User) konnekt.User {
				return konnekt.User{}
			},
		},
		"Empty last name": {
			id: 1,
			updater: func(u konnekt.UpdateUser) konnekt.UpdateUser {
				return konnekt.UpdateUser{LastName: ptr.From("")}
			},
			wantCode: konnekt.ERRINVALID,
			wantUserUpdater: func(u konnekt.User) konnekt.User {
				return konnekt.User{}
			},
		},
		"Invalid email": {
			id: 1,
			updater: func(u konnekt.UpdateUser) konnekt.UpdateUser {
				return konnekt.UpdateUser{Email: ptr.From("invalid-email.com")}
			},
			wantCode: konnekt.ERRINVALID,
			wantUserUpdater: func(u konnekt.User) konnekt.User {
				return konnekt.User{}
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			repo, dsn := MustOpenRepo(t)
			defer MustCloseRepo(t, repo, dsn)

			service := sqlite.NewUserService(repo)
			MustCreateUser(t, context.Background(), repo, baseUser, []byte("Password123!"), []byte("Password123!"))

			update := baseUpdate
			if tt.updater != nil {
				update = tt.updater(update)
			}

			user, err := service.UpdateUser(context.Background(), tt.id, update)

			code := konnekt.ErrorCode(err)

			if code != tt.wantCode {
				t.Fatalf("got code %q, want code %q, error: %v", code, tt.wantCode, err)
			}

			wantUser := baseUser
			if tt.wantUserUpdater != nil {
				wantUser = tt.wantUserUpdater(wantUser)
			}

			if !user.Equals(wantUser) {
				t.Fatalf("got %+v, want %+v", user, wantUser)
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
