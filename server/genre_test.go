package konnekt_test

import (
	"testing"

	"github.com/mattismoel/konnekt"
)

var baseGenres = []konnekt.Genre{
	{ID: 1, Name: "Rock"},
	{ID: 2, Name: "Punk"},
	{ID: 3, Name: "Indie"},
	{ID: 4, Name: "Pop"},
}

func TestGenreEquals(t *testing.T) {
	type test struct {
		a         konnekt.Genre
		b         konnekt.Genre
		wantEqual bool
	}

	tests := map[string]test{
		"Equal": {
			a:         baseGenres[0],
			b:         baseGenres[0],
			wantEqual: true,
		},
		"Not Equal": {
			a:         baseGenres[0],
			b:         baseGenres[1],
			wantEqual: false,
		},
		"Different Names": {
			a: baseGenres[0],
			b: konnekt.Genre{
				ID:   baseGenres[0].ID,
				Name: "Other Genre Name",
			},
			wantEqual: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			gotEqual := tt.a.Equals(tt.b)

			if gotEqual != tt.wantEqual {
				t.Fatalf("got %v, want %v", gotEqual, tt.wantEqual)
			}
		})
	}
}

func TestGenreValid(t *testing.T) {
	type test struct {
		g   konnekt.Genre
		err error
	}

	tests := map[string]test{
		"Valid genre": {
			g:   baseGenres[0],
			err: nil,
		},
		"Negative ID": {
			g:   konnekt.Genre{ID: -1, Name: "Rock"},
			err: konnekt.Errorf(konnekt.ERRINVALID, "Error"),
		},
		"No Name": {
			g:   konnekt.Genre{ID: 1, Name: " "},
			err: konnekt.Errorf(konnekt.ERRINVALID, "Error"),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			err := tt.g.Validate()

			if konnekt.ErrorCode(err) != konnekt.ErrorCode(tt.err) {
				t.Fatalf("got %q, want %q", err, tt.err)
			}
		})
	}
}
