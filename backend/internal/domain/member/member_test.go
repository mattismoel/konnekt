package member_test

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mattismoel/konnekt/internal/domain/member"
	"github.com/mattismoel/konnekt/internal/domain/team"
)

func TestWithID(t *testing.T) {
	type test struct {
		id      int64
		wantID  int64
		wantErr error
	}

	tests := map[string]test{
		"Valid ID": {
			id:      1,
			wantID:  1,
			wantErr: nil,
		},
		"Negative ID": {
			id:      -1,
			wantID:  0,
			wantErr: member.ErrIDInvalid,
		},
		"Zero ID": {
			id:      0,
			wantID:  0,
			wantErr: member.ErrIDInvalid,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			m, err := member.NewMember(
				member.WithID(tt.id),
			)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got %v, want %v", err, tt.wantErr)
			}

			if m.ID != tt.wantID {
				t.Fatalf("got id %d, want id %d", m.ID, tt.wantID)
			}
		})
	}
}

func TestWithFirstName(t *testing.T) {
	type test struct {
		fn      string
		wantFN  string
		wantErr error
	}

	tests := map[string]test{
		"Valid First Name": {
			fn:      "John",
			wantFN:  "John",
			wantErr: nil,
		},
		"Empty First Name": {
			fn:      "",
			wantFN:  "",
			wantErr: member.ErrFirstNameInvalid,
		},
		"Space-only first name": {
			fn:      " ",
			wantFN:  "",
			wantErr: member.ErrFirstNameInvalid,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			m, err := member.NewMember(
				member.WithFirstName(tt.fn),
			)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got %v, want %v", err, tt.wantErr)
			}

			if m.FirstName != tt.wantFN {
				t.Fatalf("got %q, want %q", m.FirstName, tt.wantFN)
			}
		})
	}
}

func TestWithLastName(t *testing.T) {
	type test struct {
		ln      string
		wantLN  string
		wantErr error
	}

	tests := map[string]test{
		"Valid Last Name": {
			ln:      "Doe",
			wantLN:  "Doe",
			wantErr: nil,
		},
		"Empty Last Name": {
			ln:      "",
			wantLN:  "",
			wantErr: member.ErrLastNameInvalid,
		},
		"Space-only last name": {
			ln:      " ",
			wantLN:  "",
			wantErr: member.ErrLastNameInvalid,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			m, err := member.NewMember(
				member.WithLastName(tt.ln),
			)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got %v, want %v", err, tt.wantErr)
			}

			if m.LastName != tt.wantLN {
				t.Fatalf("got %q, want %q", m.LastName, tt.wantLN)
			}
		})
	}
}

func TestWithEmail(t *testing.T) {
	type test struct {
		email     string
		wantEmail string
		wantErr   error
	}

	tests := map[string]test{
		"Valid email": {
			email:     "example@mail.com",
			wantEmail: "example@mail.com",
			wantErr:   nil,
		},
		"Invalid email (no @)": {
			email:     "example.mail.com",
			wantEmail: "",
			wantErr:   member.ErrEmailInvalid,
		},
		"Invalid email (no end)": {
			email:     "example@",
			wantEmail: "",
			wantErr:   member.ErrEmailInvalid,
		},
		"Empty email": {
			email:     "",
			wantEmail: "",
			wantErr:   member.ErrEmailInvalid,
		},
		"Space-only email": {
			email:     " ",
			wantEmail: "",
			wantErr:   member.ErrEmailInvalid,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			m, err := member.NewMember(
				member.WithEmail(tt.email),
			)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got %v, want %v", err, tt.wantErr)
			}

			if m.Email != tt.wantEmail {
				t.Fatalf("got %q, want %q", m.Email, tt.wantEmail)
			}
		})
	}
}

func TestWithTeams(t *testing.T) {
	exampleTeams := team.TeamCollection{
		{ID: 1, Name: "booking", DisplayName: "Booking", Description: "Booking Description..."},
		{ID: 2, Name: "public-relations", DisplayName: "PR", Description: "PR Description..."},
	}

	t.Run("Member with teams", func(t *testing.T) {
		m, err := member.NewMember(
			member.WithTeams(exampleTeams),
		)

		if err != nil {
			t.Fatalf("got %v, want nil", err)
		}

		if !cmp.Equal(m.Teams, exampleTeams) {
			t.Fatalf("slice mismatch: %s", cmp.Diff(exampleTeams, m.Teams))
		}
	})
}

func TestWithProfilePictureURL(t *testing.T) {
	type test struct {
		url     string
		wantURL string
		wantErr error
	}

	tests := map[string]test{
		"Valid URL": {
			url:     "https://example.com",
			wantURL: "https://example.com",
			wantErr: nil,
		},
		"Empty URL": {
			url:     "",
			wantURL: "",
			wantErr: member.ErrProfileImageURLInvalid,
		},
		"Space-only URL": {
			url:     " ",
			wantURL: "",
			wantErr: member.ErrProfileImageURLInvalid,
		},
		"Inaccessible URL": {
			url:     "https://example.com/image",
			wantURL: "",
			wantErr: member.ErrProfileImageURLInaccessible,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			m, err := member.NewMember(
				member.WithProfilePictureURL(tt.url),
			)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got %v, want %v", err, tt.wantErr)
			}

			if m.ProfilePictureURL != tt.wantURL {
				t.Fatalf("got %q, want %q", m.ProfilePictureURL, tt.wantURL)
			}
		})
	}
}
