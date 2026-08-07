package event_test

import (
	"errors"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/mattismoel/konnekt/internal/domain/artist"
	"github.com/mattismoel/konnekt/internal/domain/concert"
	"github.com/mattismoel/konnekt/internal/domain/event"
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
			wantErr: event.ErrInvalidID,
		},
		"Zero ID": {
			id:      0,
			wantID:  0,
			wantErr: event.ErrInvalidID,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			e, err := event.NewEvent(
				event.WithID(tt.id),
			)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got %v, want %v\n", err, tt.wantErr)
			}

			if e.ID != tt.wantID {
				t.Fatalf("got id %d, want id %d\n", e.ID, tt.wantID)
			}
		})
	}
}

func TestWithTitle(t *testing.T) {
	type test struct {
		title     string
		wantTitle string
		wantErr   error
	}

	tests := map[string]test{
		"Valid Title": {
			title:     "Example Title",
			wantTitle: "Example Title",
			wantErr:   nil,
		},
		"Empty title": {
			title:     "",
			wantTitle: "",
			wantErr:   event.ErrInvalidTitle,
		},
		"Space-only title": {
			title:     " ",
			wantTitle: "",
			wantErr:   event.ErrInvalidTitle,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			e, err := event.NewEvent(
				event.WithTitle(tt.title),
			)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got %v, want %v\n", err, tt.wantErr)
			}

			if e.Title != tt.wantTitle {
				t.Fatalf("got title %q, want title %q\n", e.Title, tt.wantTitle)
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
		"Valid Description": {
			description:     "Example Desc...",
			wantDescription: "Example Desc...",
			wantErr:         nil,
		},
		"Empty description": {
			description:     "",
			wantDescription: "",
			wantErr:         event.ErrInvalidDescription,
		},
		"Space-only description": {
			description:     " ",
			wantDescription: "",
			wantErr:         event.ErrInvalidDescription,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			e, err := event.NewEvent(
				event.WithDescription(tt.description),
			)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got %v, want %v\n", err, tt.wantErr)
			}

			if e.Description != tt.wantDescription {
				t.Fatalf("got description %q, want description %q\n", e.Description, tt.wantDescription)
			}
		})
	}
}

func TestWithImageURL(t *testing.T) {
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
		"Inaccessible URL": {
			url:     "https://example.com/inaccessible",
			wantURL: "",
			wantErr: event.ErrImageURLInaccessible,
		},
		"Empty URL": {
			url:     "",
			wantURL: "",
			wantErr: event.ErrInvalidImageURL,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			e, err := event.NewEvent(
				event.WithImageURL(tt.url),
			)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got %v, want %v\n", err, tt.wantErr)
			}

			if e.ImageURL != tt.wantURL {
				t.Fatalf("got image URL %q, want image URL %q\n", e.ImageURL, tt.wantURL)
			}
		})
	}
}

func TestWithTicketURL(t *testing.T) {
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
		"Inaccessible URL": {
			url:     "https://some-weird-non-existant-domain.com/no-exist",
			wantURL: "",
			wantErr: event.ErrTicketURLInaccessible,
		},
		"Empty URL": {
			url:     "",
			wantURL: "",
			wantErr: event.ErrTicketURLInvalid,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			e, err := event.NewEvent(
				event.WithTicketURL(tt.url),
			)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got %v, want %v\n", err, tt.wantErr)
			}

			if e.TicketURL != tt.wantURL {
				t.Fatalf("got ticket URL %q, want ticket URL %q\n", e.TicketURL, tt.wantURL)
			}
		})
	}
}

func TestWithConcerts(t *testing.T) {
	type test struct {
		concerts     []concert.Concert
		wantConcerts []concert.Concert
		wantErr      error
	}

	exampleConcerts := []concert.Concert{
		{
			ID:   1,
			From: time.Now(),
			To:   time.Now().Add(30 * time.Minute),
			Artist: artist.Artist{
				ID:          1,
				Name:        "Artist 1",
				Description: "Artist 1 Description...",
				ImageURL:    "https://example.com",
				PreviewURL:  "https://example.com",
				Genres: []artist.Genre{
					{ID: 1, Name: "Rock"},
					{ID: 2, Name: "Punk"},
				},
				Socials: []artist.Social{
					"https://example.com/1",
					"https://example.com/2",
				},
			},
		},
	}
	tests := map[string]test{
		"Valid concerts": {
			concerts:     exampleConcerts,
			wantConcerts: exampleConcerts,
			wantErr:      nil,
		},
		"Empty concerts": {
			concerts:     make([]concert.Concert, 0),
			wantConcerts: make([]concert.Concert, 0),
			wantErr:      event.ErrNoConcerts,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			e, err := event.NewEvent(
				event.WithConcerts(tt.concerts...),
			)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got %v, want %v\n", err, tt.wantErr)
			}

			if err == nil && !cmp.Equal(e.Concerts, tt.wantConcerts) {
				t.Fatalf("concerts mismatch: %s\n", cmp.Diff(tt.wantConcerts, e.Concerts))
			}
		})
	}
}

func TestWithVenue(t *testing.T) {
	exampleVenue, err := venue.NewVenue(
		venue.WithID(1),
		venue.WithCity("Odense"),
		venue.WithCountryCode("DK"),
	)

	if err != nil {
		t.Fatalf("Could not create test venue: %v", err)
	}

	t.Run("Event with venue", func(t *testing.T) {
		e, err := event.NewEvent(
			event.WithVenue(exampleVenue),
		)

		if err != nil {
			t.Fatalf("got %v, want nil", err)
		}

		if !cmp.Equal(e.Venue, exampleVenue) {
			t.Fatalf("got %+v, want  %+v\n", e.Venue, exampleVenue)
		}
	})
}

func TestWithIsPublic(t *testing.T) {
	type test struct {
		isPublic bool
	}

	tests := map[string]bool{
		"Public (true)":  true,
		"Public (false)": false,
	}

	for name, isPublic := range tests {
		t.Run(name, func(t *testing.T) {
			e, err := event.NewEvent(
				event.WithIsPublic(isPublic),
			)

			if err != nil {
				t.Fatalf("got %v, want nil", err)
			}

			if !e.IsPublic == isPublic {
				t.Fatalf("got %v, want %v", e.IsPublic, isPublic)
			}
		})
	}
}
