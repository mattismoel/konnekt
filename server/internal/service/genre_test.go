package service_test

import (
	"testing"

	"github.com/mattismoel/konnekt"
	"github.com/mattismoel/konnekt/internal/service"
)

type genreUpdaterFunc func(service.Genre) service.Genre

var baseGenres = []service.Genre{"Rock", "Punk", "Indie", "Pop"}

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
			bUpdater: func(g service.Genre) service.Genre {
				return baseGenres[1]
			},
			wantEqual: false,
		},
		"Different Names": {
			aUpdater: nil,
			bUpdater: func(g service.Genre) service.Genre {
				g = "Other Genre name"
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
		"No Name": {
			updater: func(g service.Genre) service.Genre {
				g = " "
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
