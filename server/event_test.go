package konnekt_test

import (
	"testing"
	"time"

	"github.com/mattismoel/konnekt"
)

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
		a          konnekt.Event
		b          konnekt.Event
		wantEquals bool
	}

	tests := map[string]test{
		"Equal": {
			a:          baseEvent,
			b:          baseEvent,
			wantEquals: true,
		},
		"Title Differ": {
			a: baseEvent,
			b: konnekt.Event{
				ID:          baseEvent.ID,
				Title:       "Other Title",
				Description: baseEvent.Description,
				FromDate:    baseEvent.FromDate,
				ToDate:      baseEvent.ToDate,
				Address:     baseEvent.Address,
				Genres:      baseEvent.Genres,
			},
			wantEquals: false,
		},
		"Description Differ": {
			a: baseEvent,
			b: konnekt.Event{
				ID:          baseEvent.ID,
				Title:       baseEvent.Title,
				Description: "Other Description",
				FromDate:    baseEvent.FromDate,
				ToDate:      baseEvent.ToDate,
				Address:     baseEvent.Address,
				Genres:      baseEvent.Genres,
			},
			wantEquals: false,
		},
		"FromDate Differ": {
			a: baseEvent,
			b: konnekt.Event{
				ID:          baseEvent.ID,
				Title:       baseEvent.Title,
				Description: baseEvent.Description,
				FromDate:    baseEvent.FromDate.Add(1 * time.Hour),
				ToDate:      baseEvent.ToDate,
				Address:     baseEvent.Address,
				Genres:      baseEvent.Genres,
			},
			wantEquals: false,
		},
		"ToDate Differ": {
			a: baseEvent,
			b: konnekt.Event{
				ID:          baseEvent.ID,
				Title:       baseEvent.Title,
				Description: baseEvent.Description,
				FromDate:    baseEvent.FromDate,
				ToDate:      baseEvent.ToDate.Add(2 * time.Hour),
				Address:     baseEvent.Address,
				Genres:      baseEvent.Genres,
			},
			wantEquals: false,
		},
		"Addresses Differ": {
			a: baseEvent,
			b: konnekt.Event{
				ID:          baseEvent.ID,
				Title:       baseEvent.Title,
				Description: baseEvent.Description,
				FromDate:    baseEvent.FromDate,
				ToDate:      baseEvent.ToDate,
				Address: konnekt.Address{
					Country:     "Sweden",
					City:        "Stockholm",
					Street:      "Other Street",
					HouseNumber: "19B",
				},
				Genres: baseEvent.Genres,
			},
			wantEquals: false,
		},
		"Genres Differ": {
			a: baseEvent,
			b: konnekt.Event{
				ID:          baseEvent.ID,
				Title:       baseEvent.Title,
				Description: baseEvent.Description,
				FromDate:    baseEvent.FromDate,
				ToDate:      baseEvent.ToDate,
				Address:     baseEvent.Address,
				Genres: []konnekt.Genre{
					{ID: 5, Name: "New Age"},
				},
			},
			wantEquals: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			gotEquals := tt.a.Equals(tt.b)

			if gotEquals != tt.wantEquals {
				t.Fatalf("got %v, want %v", gotEquals, tt.wantEquals)
			}
		})
	}
}

func TestEventValidate(t *testing.T) {
	type test struct {
		e   konnekt.Event
		err error
	}

	tests := map[string]test{
		"Valid event": {
			e:   baseEvent,
			err: nil,
		},
		"Negative ID": {
			e: konnekt.Event{
				ID:          -1,
				Title:       baseEvent.Title,
				Description: baseEvent.Description,
				FromDate:    baseEvent.FromDate,
				ToDate:      baseEvent.ToDate,
				Address:     baseEvent.Address,
				Genres:      baseGenres,
			},
			err: konnekt.Errorf(konnekt.ERRINVALID, ""),
		},
		"No Title": {
			e: konnekt.Event{
				ID:          baseEvent.ID,
				Title:       " ",
				Description: baseEvent.Description,
				FromDate:    baseEvent.FromDate,
				ToDate:      baseEvent.ToDate,
				Address:     baseEvent.Address,
				Genres:      baseGenres,
			},
			err: konnekt.Errorf(konnekt.ERRINVALID, ""),
		},
		"No Description": {
			e: konnekt.Event{
				ID:          baseEvent.ID,
				Title:       baseEvent.Title,
				Description: " ",
				FromDate:    baseEvent.FromDate,
				ToDate:      baseEvent.ToDate,
				Address:     baseEvent.Address,
				Genres:      baseGenres,
			},
			err: konnekt.Errorf(konnekt.ERRINVALID, ""),
		},
		"Zero FromDate": {
			e: konnekt.Event{
				ID:          baseEvent.ID,
				Title:       baseEvent.Title,
				Description: baseEvent.Description,
				FromDate:    time.Time{},
				ToDate:      baseEvent.ToDate,
				Address:     baseEvent.Address,
				Genres:      baseGenres,
			},
			err: konnekt.Errorf(konnekt.ERRINVALID, ""),
		},
		"Zero ToDate": {
			e: konnekt.Event{
				ID:          baseEvent.ID,
				Title:       baseEvent.Title,
				Description: baseEvent.Description,
				FromDate:    baseEvent.FromDate,
				ToDate:      time.Time{},
				Address:     baseEvent.Address,
				Genres:      baseGenres,
			},
			err: konnekt.Errorf(konnekt.ERRINVALID, ""),
		},
		"No Address": {
			e: konnekt.Event{
				ID:          baseEvent.ID,
				Title:       baseEvent.Title,
				Description: baseEvent.Description,
				FromDate:    baseEvent.FromDate,
				ToDate:      baseEvent.ToDate,
				Address:     konnekt.Address{},
				Genres:      baseGenres,
			},
			err: konnekt.Errorf(konnekt.ERRINVALID, ""),
		},
		"Nil Genres": {
			e: konnekt.Event{
				ID:          baseEvent.ID,
				Title:       baseEvent.Title,
				Description: baseEvent.Description,
				FromDate:    baseEvent.FromDate,
				ToDate:      baseEvent.ToDate,
				Address:     baseEvent.Address,
			},
			err: konnekt.Errorf(konnekt.ERRINVALID, ""),
		},
		"Empty Genres": {
			e: konnekt.Event{
				ID:          baseEvent.ID,
				Title:       baseEvent.Title,
				Description: baseEvent.Description,
				FromDate:    baseEvent.FromDate,
				ToDate:      baseEvent.ToDate,
				Address:     baseEvent.Address,
				Genres:      []konnekt.Genre{},
			},
			err: konnekt.Errorf(konnekt.ERRINVALID, ""),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			err := tt.e.Validate()

			if konnekt.ErrorCode(err) != konnekt.ErrorCode(tt.err) {
				t.Fatalf("got %v, want %v", err, tt.err)
			}
		})
	}
}
