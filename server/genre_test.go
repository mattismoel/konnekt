package konnekt_test

import (
	"testing"

	"github.com/mattismoel/konnekt"
)

type genreUpdaterFunc func(konnekt.Genre) konnekt.Genre

var baseGenres = []konnekt.Genre{
	{ID: 1, Name: "Rock"},
	{ID: 2, Name: "Punk"},
	{ID: 3, Name: "Indie"},
	{ID: 4, Name: "Pop"},
}

func TestGenreEquals(t *testing.T) {
	type test struct {
		aUpdater  genreUpdaterFunc
		bUpdater  genreUpdaterFunc
		wantEqual bool
	}

	tests := map[string]test{
		"Equal": {
			aUpdater:  nil,
			bUpdater:  nil,
			wantEqual: true,
		},
		"Not Equal": {
			aUpdater: nil,
			bUpdater: func(g konnekt.Genre) konnekt.Genre {
				return baseGenres[1]
			},
			wantEqual: false,
		},
		"Different Names": {
			aUpdater: nil,
			bUpdater: func(g konnekt.Genre) konnekt.Genre {
				g.Name = "Other Genre name"
				return g
			},
			wantEqual: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			a, b := baseGenres[0], baseGenres[0]

			if tt.aUpdater != nil {
				a = tt.aUpdater(a)
			}

			if tt.bUpdater != nil {
				b = tt.bUpdater(b)
			}

			gotEqual := a.Equals(b)

			if gotEqual != tt.wantEqual {
				t.Fatalf("got %v, want %v", gotEqual, tt.wantEqual)
			}
		})
	}
}

func TestGenreValid(t *testing.T) {
	type test struct {
		updater genreUpdaterFunc
		err     error
	}

	tests := map[string]test{
		"Valid genre": {
			updater: nil,
			err:     nil,
		},
		"Negative ID": {
			updater: func(g konnekt.Genre) konnekt.Genre {
				g.ID = -1
				return g
			},
			err: konnekt.Errorf(konnekt.ERRINVALID, "Error"),
		},
		"No Name": {
			updater: func(g konnekt.Genre) konnekt.Genre {
				g.Name = " "
				return g
			},
			err: konnekt.Errorf(konnekt.ERRINVALID, "Error"),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			genre := baseGenres[0]

			if tt.updater != nil {
				genre = tt.updater(genre)
			}

			err := genre.Validate()

			if konnekt.ErrorCode(err) != konnekt.ErrorCode(tt.err) {
				t.Fatalf("got %q, want %q", err, tt.err)
			}
		})
	}
}
