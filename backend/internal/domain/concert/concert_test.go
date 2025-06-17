package concert_test

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/mattismoel/konnekt/internal/domain/artist"
	"github.com/mattismoel/konnekt/internal/domain/concert"
)

func TestWithID(t *testing.T) {
	type test struct {
		id      int64
		wantID  int64
		wantErr error
	}

	tests := map[string]test{
		"Valid ID": {
			id:      10,
			wantID:  10,
			wantErr: nil,
		},
		"Negative ID": {
			id:      -1,
			wantID:  0,
			wantErr: concert.ErrInvalidID,
		},
		"Zero ID": {
			id:      0,
			wantID:  0,
			wantErr: concert.ErrInvalidID,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c, err := concert.NewConcert(
				concert.WithID(tt.id),
			)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got %v, want %v\n", err, tt.wantErr)
			}

			if c.ID != tt.wantID {
				t.Fatalf("got id %d, want %d\n", c.ID, tt.wantID)
			}
		})
	}
}

func TestWithArtist(t *testing.T) {
	exampleArtist, err := artist.NewArtist(
		artist.WithID(1),
		artist.WithName("Example Artist"),
		artist.WithDescription("Artist description..."),
		artist.WithGenres(
			artist.Genre{ID: 1, Name: "Rock"},
			artist.Genre{ID: 2, Name: "Pop"},
			artist.Genre{ID: 3, Name: "RnB"},
		),
		artist.WithPreviewURL("https://example.com"),
		artist.WithSocials("https://example.com/1", "https://example.com/2"),
	)

	if err != nil {
		t.Fatalf("Could not create example artist: %v", err)
	}

	t.Run("Create concert with artist", func(t *testing.T) {
		c, err := concert.NewConcert(
			concert.WithArtist(exampleArtist),
		)

		if err != nil {
			t.Fatalf("Could not create concert: %v", err)
		}

		if !reflect.DeepEqual(c.Artist, exampleArtist) {
			t.Fatalf("got artist:\n%+v\n, want:\n%+v\n", c.Artist, exampleArtist)
		}
	})
}

func TestWithDate(t *testing.T) {
	type test struct {
		from    time.Time
		to      time.Time
		wantErr error
	}

	tests := map[string]test{
		"Valid date relationship": {
			from:    time.Now(),
			to:      time.Now().Add(30 * time.Minute),
			wantErr: nil,
		},
		"To date before from date": {
			from:    time.Now(),
			to:      time.Now().Add(-30 * time.Minute),
			wantErr: concert.ErrInvalidDateRelationship,
		},
		"From date after to date": {
			from:    time.Now().Add(30 * time.Minute),
			to:      time.Now(),
			wantErr: concert.ErrInvalidDateRelationship,
		},
		"Zero From date": {
			from:    time.Time{},
			to:      time.Now().Add(-30 * time.Minute),
			wantErr: concert.ErrInvalidDate,
		},
		"Zero To date": {
			from:    time.Now(),
			to:      time.Time{},
			wantErr: concert.ErrInvalidDate,
		},
		"Zero dates": {
			from:    time.Time{},
			to:      time.Time{},
			wantErr: concert.ErrInvalidDate,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := concert.NewConcert(
				concert.WithFrom(tt.from),
				concert.WithTo(tt.to),
			)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got %v, want %v\n", err, tt.wantErr)
			}
		})
	}
}
