package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/mattismoel/konnekt/internal/service"
	"github.com/mattismoel/konnekt/internal/storage/mock"
)

type userUpdaterFunc func(service.User) service.User

var baseUser = service.User{
	ID:        1,
	Email:     "test@mail.com",
	FirstName: "John",
	LastName:  "Doe",
}

func TestUserValidate(t *testing.T) {
	type test struct {
		updater userUpdaterFunc
		err     error
	}

	tests := map[string]test{
		"Valid user": {
			updater: nil,
			err:     nil,
		},
		"Empty email": {
			updater: func(u service.User) service.User {
				u.Email = ""
				return u
			},
			err: service.ErrInvalidEmail,
		},
		"Invalid email": {
			updater: func(u service.User) service.User {
				u.Email = ""
				return u
			},
			err: service.ErrInvalidEmail,
		},
		"No first name": {
			updater: func(u service.User) service.User {
				u.FirstName = ""
				return u
			},
			err: service.ErrInvalidFirstName,
		},
		"No last name": {
			updater: func(u service.User) service.User {
				u.LastName = ""
				return u
			},
			err: service.ErrInvalidLastName,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			user := baseUser

			if tt.updater != nil {
				user = tt.updater(user)
			}

			err := user.Validate()

			if !errors.Is(err, tt.err) {
				t.Fatalf("got %v, want %v", err, tt.err)
			}
		})
	}
}

func TestUsersEqual(t *testing.T) {
	type test struct {
		aUpdater   userUpdaterFunc
		bUpdater   userUpdaterFunc
		wantEquals bool
	}

	tests := map[string]test{
		"Equal": {
			aUpdater:   nil,
			bUpdater:   nil,
			wantEquals: true,
		},
		"Email differ": {
			aUpdater: nil,
			bUpdater: func(u service.User) service.User {
				u.Email = "other@mail.com"
				return u
			},
			wantEquals: false,
		},
		"First Name differ": {
			aUpdater: nil,
			bUpdater: func(u service.User) service.User {
				u.FirstName = "Sophie"
				return u
			},
			wantEquals: false,
		},
		"Last name differ": {
			aUpdater: nil,
			bUpdater: func(u service.User) service.User {
				u.LastName = "Johnson"
				return u
			},
			wantEquals: false,
		},
		"ID differ": {
			aUpdater: nil,
			bUpdater: func(u service.User) service.User {
				u.ID = 999
				return u
			},
			wantEquals: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			a, b := baseUser, baseUser

			if tt.aUpdater != nil {
				a = tt.aUpdater(a)
			}

			if tt.bUpdater != nil {
				b = tt.bUpdater(b)
			}

			isEqual := a.Equals(b)

			if isEqual != tt.wantEquals {
				t.Fatalf("got %v, want %v", isEqual, tt.wantEquals)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	user := service.User{
		Email:     "test-email@mail.com",
		FirstName: "Peter",
		LastName:  "Parker",
	}

	type updater func(service.User) service.User

	type test struct {
		updater         updater
		wantUpdater     updater
		password        []byte
		passwordConfirm []byte
		err             error
	}

	tests := map[string]test{
		"Valid user": test{
			updater: func(u service.User) service.User {
				return user
			},
			err: nil,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			mock := mock.NewMockStorage()
			service := service.NewUserService(mock)

			user := tt.updater(user)

			u, err := service.CreateUser(context.Background(), user, tt.password, tt.passwordConfirm)
			if !errors.Is(err, tt.err) {
				t.Fatalf("got %v, want %v", err, tt.err)
			}

		})
	}
}
