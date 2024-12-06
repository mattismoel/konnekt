package service_test

import (
	"testing"
	"time"

	"github.com/mattismoel/konnekt"
	"github.com/mattismoel/konnekt/internal/service"
)

type eventUpdaterFunc func(service.Event) service.Event

var now = time.Now()

func TestEventEquals(t *testing.T) {
	var genres = []string{"Rock", "Punk", "Hip Hop", "R&B"}

	var baseAddress = service.Address{
		Country:     "Denmark",
		City:        "Odense",
		Street:      "Postenvej",
		HouseNumber: "18A",
	}

	var baseEvent = service.Event{
		ID:          1,
		Title:       "Base Title",
		Description: "Base Description",
		FromDate:    now,
		ToDate:      now.Add(2 * time.Hour),
		Address:     baseAddress,
		Genres:      genres[:2],
	}

	type updater func(service.Event) service.Event

	type test struct {
		aUpdater   updater
		bUpdater   updater
		wantEquals bool
	}

	tests := map[string]test{
		"Equal": {
			aUpdater:   nil,
			bUpdater:   nil,
			wantEquals: true,
		},
		"Title Differ": {
			aUpdater: nil,
			bUpdater: func(e service.Event) service.Event {
				e.Title = "Other Title"
				return e
			},
			wantEquals: false,
		},
		"Description Differ": {
			aUpdater: nil,
			bUpdater: func(e service.Event) service.Event {
				e.Description = "Other Description"
				return e
			},
			wantEquals: false,
		},
		"FromDate Differ": {
			aUpdater: nil,
			bUpdater: func(e service.Event) service.Event {
				e.FromDate = baseEvent.FromDate.Add(1 * time.Hour)
				return e
			},
			wantEquals: false,
		},
		"ToDate Differ": {
			aUpdater: nil,
			bUpdater: func(e service.Event) service.Event {
				e.ToDate = baseEvent.ToDate.Add(2 * time.Hour)
				return e
			},
			wantEquals: false,
		},
		"Addresses Differ": {
			aUpdater: nil,
			bUpdater: func(e service.Event) service.Event {
				e.Address = service.Address{
					Country:     "Sweden",
					City:        "Stockholm",
					Street:      "Other Street",
					HouseNumber: "19B",
				}
				return e
			},
			wantEquals: false,
		},
		"Genres Differ": {
			aUpdater: nil,
			bUpdater: func(e service.Event) service.Event {
				e.Genres = genres[2:]
				return e
			},
			wantEquals: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			a, b := baseEvent, baseEvent

			if tt.aUpdater != nil {
				a = tt.aUpdater(a)
			}

			if tt.bUpdater != nil {
				b = tt.bUpdater(b)
			}

			gotEquals := a.Equals(b)

			if gotEquals != tt.wantEquals {
				t.Fatalf("got %v, want %v", gotEquals, tt.wantEquals)
			}
		})
	}
}

func TestEventValidate(t *testing.T) {
	var genres = []string{"Rock", "Punk", "Hip Hop", "R&B"}

	var baseAddress = service.Address{
		Country:     "Denmark",
		City:        "Odense",
		Street:      "Postenvej",
		HouseNumber: "18A",
	}

	var baseEvent = service.Event{
		ID:          1,
		Title:       "Base Title",
		Description: "Base Description",
		FromDate:    now,
		ToDate:      now.Add(2 * time.Hour),
		Address:     baseAddress,
		Genres:      genres[:2],
	}

	type test struct {
		updater eventUpdaterFunc
		err     error
	}

	tests := map[string]test{
		"Valid event": {
			updater: nil,
			err:     nil,
		},
		"No Title": {
			updater: func(e service.Event) service.Event {
				e.Title = " "
				return e
			},
			err: service.ErrNoTitle,
		},
		"No Description": {
			updater: func(e service.Event) service.Event {
				e.Description = " "
				return e
			},
			err: service.ErrNoDescription,
		},
		"Zero FromDate": {
			updater: func(e service.Event) service.Event {
				e.FromDate = time.Time{}
				return e
			},
			err: service.ErrZeroFromDay,
		},
		"Zero ToDate": {
			updater: func(e service.Event) service.Event {
				e.ToDate = time.Time{}
				return e
			},
			err: service.ErrZeroToDay,
		},
		"No Address": {
			updater: func(e service.Event) service.Event {
				e.Address = service.Address{}
				return e
			},
			err: service.ErrEmptyAddress,
		},
		"Nil Genres": {
			updater: func(e service.Event) service.Event {
				e.Genres = nil
				return e
			},
			err: service.ErrNoGenres,
		},
		"Empty Genres": {
			updater: func(e service.Event) service.Event {
				e.Genres = []string{}
				return e
			},
			err: service.ErrNoGenres,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			event := baseEvent

			if tt.updater != nil {
				event = tt.updater(event)
			}

			err := event.Validate()

			if konnekt.ErrorCode(err) != konnekt.ErrorCode(tt.err) {
				t.Fatalf("got %v, want %v", err, tt.err)
			}
		})
	}
}
