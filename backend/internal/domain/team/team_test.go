package team_test

import (
	"errors"
	"testing"

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
			wantErr: team.ErrTeamIDInvalid,
		},
		"Zero ID": {
			id:      0,
			wantID:  0,
			wantErr: team.ErrTeamIDInvalid,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			tm, err := team.NewTeam(
				team.WithID(tt.id),
			)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got %v, want %v", err, tt.wantErr)
			}

			if tm.ID != tt.wantID {
				t.Fatalf("got %d, want %d", tm.ID, tt.wantID)
			}
		})
	}
}

func TestWithName(t *testing.T) {
	type test struct {
		name     string
		wantName string
		wantErr  error
	}

	tests := map[string]test{
		"Valid Name": {
			name:     "booking",
			wantName: "booking",
			wantErr:  nil,
		},
		"Empty name": {
			name:     "",
			wantName: "",
			wantErr:  team.ErrTeamNameInvalid,
		},
		"Space-only name": {
			name:     " ",
			wantName: "",
			wantErr:  team.ErrTeamNameInvalid,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			tm, err := team.NewTeam(
				team.WithName(tt.name),
			)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got %v, want %v", err, tt.wantErr)
			}

			if tm.Name != tt.wantName {
				t.Fatalf("got %q, want %q", tm.Name, tt.wantName)
			}
		})
	}
}

func TestWithDisplayName(t *testing.T) {
	type test struct {
		dn      string
		wantDN  string
		wantErr error
	}

	tests := map[string]test{
		"Valid display name": {
			dn:      "Booking",
			wantDN:  "Booking",
			wantErr: nil,
		},
		"Empty display name": {
			dn:      "",
			wantDN:  "",
			wantErr: team.ErrTeamDisplayNameInvalid,
		},
		"Space-only display name": {
			dn:      " ",
			wantDN:  "",
			wantErr: team.ErrTeamDisplayNameInvalid,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			tm, err := team.NewTeam(
				team.WithDisplayName(tt.dn),
			)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got %v, want %v", err, tt.wantErr)
			}

			if tm.DisplayName != tt.wantDN {
				t.Fatalf("got %q, want %q", tm.DisplayName, tt.wantDN)
			}
		})
	}
}

func TestWithDescription(t *testing.T) {
	type test struct {
		description     string
		wantDescription string
		wantErr         error
	}

	tests := map[string]test{
		"Valid description": {
			description:     "Example description...",
			wantDescription: "Example description...",
			wantErr:         nil,
		},
		"Empty description": {
			description:     "",
			wantDescription: "",
			wantErr:         team.ErrTeamDescriptionInvalid,
		},
		"Space-only description": {
			description:     " ",
			wantDescription: "",
			wantErr:         team.ErrTeamDescriptionInvalid,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			tm, err := team.NewTeam(
				team.WithDescription(tt.description),
			)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got %v, want %v", err, tt.wantErr)
			}

			if tm.Description != tt.wantDescription {
				t.Fatalf("got %q, want %q", tm.Description, tt.wantDescription)
			}
		})
	}
}
