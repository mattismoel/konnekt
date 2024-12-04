package konnekt_test

import (
	"testing"
	"time"

	"github.com/mattismoel/konnekt"
)

type eventUpdaterFunc func(konnekt.Event) konnekt.Event

var now = time.Now()

var baseEvent = konnekt.Event{
	ID:          1,
	Title:       "Base Title",
	Description: "Base Description",
	FromDate:    now,
	ToDate:      now.Add(2 * time.Hour),
	Address:     baseAddress,
	Genres: []konnekt.Genre{
		{ID: 1, Name: "Rock"},
		{ID: 2, Name: "Punk"},
	},
}

func TestEventEquals(t *testing.T) {
	type test struct {
		aUpdater   eventUpdaterFunc
		bUpdater   eventUpdaterFunc
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
			bUpdater: func(e konnekt.Event) konnekt.Event {
				e.Title = "Other Title"
				return e
			},
			wantEquals: false,
		},
		"Description Differ": {
			aUpdater: nil,
			bUpdater: func(e konnekt.Event) konnekt.Event {
				e.Description = "Other Description"
				return e
			},
			wantEquals: false,
		},
		"FromDate Differ": {
			aUpdater: nil,
			bUpdater: func(e konnekt.Event) konnekt.Event {
				e.FromDate = baseEvent.FromDate.Add(1 * time.Hour)
				return e
			},
			wantEquals: false,
		},
		"ToDate Differ": {
			aUpdater: nil,
			bUpdater: func(e konnekt.Event) konnekt.Event {
				e.ToDate = baseEvent.ToDate.Add(2 * time.Hour)
				return e
			},
			wantEquals: false,
		},
		"Addresses Differ": {
			aUpdater: nil,
			bUpdater: func(e konnekt.Event) konnekt.Event {
				e.Address = konnekt.Address{
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
			bUpdater: func(e konnekt.Event) konnekt.Event {
				e.Genres = baseGenres[2:]
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
	type test struct {
		updater eventUpdaterFunc
		err     error
	}

	tests := map[string]test{
		"Valid event": {
			updater: nil,
			err:     nil,
		},
		"Negative ID": {
			updater: func(e konnekt.Event) konnekt.Event {
				e.ID = -1
				return e
			},
			err: konnekt.Errorf(konnekt.ERRINVALID, ""),
		},
		"No Title": {
			updater: func(e konnekt.Event) konnekt.Event {
				e.Title = " "
				return e
			},
			err: konnekt.Errorf(konnekt.ERRINVALID, ""),
		},
		"No Description": {
			updater: func(e konnekt.Event) konnekt.Event {
				e.Description = " "
				return e
			},
			err: konnekt.Errorf(konnekt.ERRINVALID, ""),
		},
		"Zero FromDate": {
			updater: func(e konnekt.Event) konnekt.Event {
				e.FromDate = time.Time{}
				return e
			},
			err: konnekt.Errorf(konnekt.ERRINVALID, ""),
		},
		"Zero ToDate": {
			updater: func(e konnekt.Event) konnekt.Event {
				e.ToDate = time.Time{}
				return e
			},
			err: konnekt.Errorf(konnekt.ERRINVALID, ""),
		},
		"No Address": {
			updater: func(e konnekt.Event) konnekt.Event {
				e.Address = konnekt.Address{}
				return e
			},
			err: konnekt.Errorf(konnekt.ERRINVALID, ""),
		},
		"Nil Genres": {
			updater: func(e konnekt.Event) konnekt.Event {
				e.Genres = nil
				return e
			},
			err: konnekt.Errorf(konnekt.ERRINVALID, ""),
		},
		"Empty Genres": {
			updater: func(e konnekt.Event) konnekt.Event {
				e.Genres = []konnekt.Genre{}
				return e
			},
			err: konnekt.Errorf(konnekt.ERRINVALID, ""),
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
