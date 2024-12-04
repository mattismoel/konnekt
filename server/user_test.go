package konnekt_test

import (
	"testing"

	"github.com/mattismoel/konnekt"
)

type userUpdaterFunc func(konnekt.User) konnekt.User

var baseUser = konnekt.User{
	ID:        1,
	Email:     "test@mail.com",
	FirstName: "John",
	LastName:  "Doe",
}

func TestUserValidate(t *testing.T) {
	type test struct {
		updater  userUpdaterFunc
		wantCode string
	}

	tests := map[string]test{
		"Valid user": {
			updater:  nil,
			wantCode: "",
		},
		"Negative ID": {
			updater: func(u konnekt.User) konnekt.User {
				u.ID = -1
				return u
			},
			wantCode: konnekt.ERRINVALID,
		},
		"Empty email": {
			updater: func(u konnekt.User) konnekt.User {
				u.Email = ""
				return u
			},
			wantCode: konnekt.ERRINVALID,
		},
		"Invalid email": {
			updater: func(u konnekt.User) konnekt.User {
				u.Email = ""
				return u
			},
			wantCode: konnekt.ERRINVALID,
		},
		"No first name": {
			updater: func(u konnekt.User) konnekt.User {
				u.FirstName = ""
				return u
			},
			wantCode: konnekt.ERRINVALID,
		},
		"No last name": {
			updater: func(u konnekt.User) konnekt.User {
				u.LastName = ""
				return u
			},
			wantCode: konnekt.ERRINVALID,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			user := baseUser

			if tt.updater != nil {
				user = tt.updater(user)
			}

			err := user.Validate()

			code := konnekt.ErrorCode(err)

			if code != tt.wantCode {
				t.Fatalf("got code %q, want code %q, error: %v", code, tt.wantCode, err)
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
			bUpdater: func(u konnekt.User) konnekt.User {
				u.Email = "other@mail.com"
				return u
			},
			wantEquals: false,
		},
		"First Name differ": {
			aUpdater: nil,
			bUpdater: func(u konnekt.User) konnekt.User {
				u.FirstName = "Sophie"
				return u
			},
			wantEquals: false,
		},
		"Last name differ": {
			aUpdater: nil,
			bUpdater: func(u konnekt.User) konnekt.User {
				u.LastName = "Johnson"
				return u
			},
			wantEquals: false,
		},
		"ID differ": {
			aUpdater: nil,
			bUpdater: func(u konnekt.User) konnekt.User {
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
