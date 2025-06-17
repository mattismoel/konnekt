package venue_test

import (
	"errors"
	"testing"

	"github.com/mattismoel/konnekt/internal/domain/venue"
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
			wantErr: venue.ErrInvalidID,
		},
		"Zero ID": {
			id:      0,
			wantID:  0,
			wantErr: venue.ErrInvalidID,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			v, err := venue.NewVenue(
				venue.WithID(tt.id),
			)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got %v, want %v", err, tt.wantErr)
			}

			if v.ID != tt.wantID {
				t.Fatalf("got %d, want %d", v.ID, tt.wantID)
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
		"Valid name": {
			name:     "Posten",
			wantName: "Posten",
			wantErr:  nil,
		},
		"Empty name": {
			name:     "",
			wantName: "",
			wantErr:  venue.ErrInvalidName,
		},
		"Space-only name": {
			name:     " ",
			wantName: "",
			wantErr:  venue.ErrInvalidName,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			v, err := venue.NewVenue(
				venue.WithName(tt.name),
			)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got %v, want %v", err, tt.wantErr)
			}

			if v.Name != tt.wantName {
				t.Fatalf("got %q, want %q", v.Name, tt.wantName)
			}
		})
	}
}

func TestWithCity(t *testing.T) {
	type test struct {
		city     string
		wantCity string
		wantErr  error
	}

	tests := map[string]test{
		"Valid city": {
			city:     "Odense",
			wantCity: "Odense",
			wantErr:  nil,
		},
		"Empty city": {
			city:     "",
			wantCity: "",
			wantErr:  venue.ErrInvalidCity,
		},
		"Space-only city": {
			city:     " ",
			wantCity: "",
			wantErr:  venue.ErrInvalidCity,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			v, err := venue.NewVenue(
				venue.WithCity(tt.city),
			)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got %v, want %v", err, tt.wantErr)
			}

			if v.City != tt.wantCity {
				t.Fatalf("got %q, want %q", v.City, tt.wantCity)
			}
		})
	}
}

func TestWithCountry(t *testing.T) {
	type test struct {
		country     string
		wantCountry string
		wantErr     error
	}

	tests := map[string]test{
		"Valid country": {
			country:     "DK",
			wantCountry: "DK",
			wantErr:     nil,
		},
		"Empty country": {
			country:     "",
			wantCountry: "",
			wantErr:     venue.ErrInvalidCountryCode,
		},
		"Space-only country": {
			country:     " ",
			wantCountry: "",
			wantErr:     venue.ErrInvalidCountryCode,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			v, err := venue.NewVenue(
				venue.WithCountryCode(tt.country),
			)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got %v, want %v", err, tt.wantErr)
			}

			if v.CountryCode != tt.wantCountry {
				t.Fatalf("got %q, want %q", v.CountryCode, tt.wantCountry)
			}
		})
	}
}
