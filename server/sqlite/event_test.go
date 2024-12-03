package sqlite_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/mattismoel/konnekt"
	"github.com/mattismoel/konnekt/internal/ptr"
	"github.com/mattismoel/konnekt/sqlite"
)

var now = time.Now()

const BASE_DURATION = 2 * time.Hour

var baseGenres = []konnekt.Genre{
	{ID: 1, Name: "Rock"},
	{ID: 2, Name: "Punk"},
}

var baseAddress = konnekt.Address{
	Country:     "Denmark",
	City:        "Odense",
	Street:      "Postenvej",
	HouseNumber: "18A",
}

var baseEvent = konnekt.Event{
	ID:          1,
	Title:       "Base Title",
	Description: "Base description",
	FromDate:    now,
	ToDate:      now.Add(BASE_DURATION),
	Address:     baseAddress,
	Genres:      baseGenres,
}

func TestCreateEvent(t *testing.T) {
	type test struct {
		event    konnekt.Event
		wantId   int64
		wantCode string
	}

	tests := map[string]test{
		"Valid event": {
			event:    baseEvent,
			wantId:   1,
			wantCode: "",
		},
		"No Title": {
			event: konnekt.Event{
				Title:       "",
				Description: baseEvent.Description,
				FromDate:    baseEvent.FromDate,
				ToDate:      baseEvent.ToDate,
				Address:     baseEvent.Address,
				Genres:      baseEvent.Genres,
			},
			wantId:   0,
			wantCode: konnekt.ERRINVALID,
		},
		"No description": {
			event: konnekt.Event{
				Title:       baseEvent.Title,
				Description: "",
				FromDate:    baseEvent.FromDate,
				ToDate:      baseEvent.ToDate,
				Address:     baseEvent.Address,
				Genres:      baseEvent.Genres,
			},
			wantCode: konnekt.ERRINVALID,
			wantId:   0,
		},
		"Zero FromDate": {
			event: konnekt.Event{
				Title:       baseEvent.Title,
				Description: baseEvent.Description,
				FromDate:    time.Time{},
				ToDate:      baseEvent.ToDate,
				Address:     baseEvent.Address,
				Genres:      baseEvent.Genres,
			},
			wantCode: konnekt.ERRINVALID,
			wantId:   0,
		},
		"Zero ToDate": {
			event: konnekt.Event{
				Title:       baseEvent.Title,
				Description: baseEvent.Description,
				FromDate:    baseEvent.FromDate,
				ToDate:      time.Time{},
				Address:     baseEvent.Address,
				Genres:      baseEvent.Genres,
			},
			wantCode: konnekt.ERRINVALID,
			wantId:   0,
		},
		"No Genres": {
			event: konnekt.Event{
				Title:       baseEvent.Title,
				Description: baseEvent.Description,
				FromDate:    baseEvent.FromDate,
				ToDate:      time.Time{},
				Address:     baseEvent.Address,
				Genres:      []konnekt.Genre{},
			},
			wantCode: konnekt.ERRINVALID,
			wantId:   0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			repo, dsn := MustOpenRepo(t)

			service := sqlite.NewEventService(repo)
			event, err := service.CreateEvent(context.Background(), test.event)

			code := konnekt.ErrorCode(err)

			if code != test.wantCode {
				testError(t, repo, dsn, fmt.Errorf("want code %q, got code %q, error: %v", test.wantCode, code, err.Error()))
				return
			}

			if event.ID != test.wantId {
				testError(t, repo, dsn, fmt.Errorf("want id %d, got id %d", test.wantId, event.ID))
				return
			}
		})
	}
}

func TestGetEventByID(t *testing.T) {
	type test struct {
		id       int64
		wantCode string
	}

	tests := map[string]test{
		"Valid ID": {
			id:       1,
			wantCode: "",
		},
		"Invalid ID": {
			id:       2,
			wantCode: konnekt.ERRNOTFOUND,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			repo, dsn := MustOpenRepo(t)
			defer MustCloseRepo(t, repo, dsn)

			MustCreateEvent(t, context.Background(), repo, baseEvent)

			service := sqlite.NewEventService(repo)

			_, err := service.FindEventByID(context.Background(), tt.id)

			code := konnekt.ErrorCode(err)

			if code != tt.wantCode {
				t.Fatalf("got code %q, want code %q, err: %v", code, tt.wantCode, err)
			}
		})
	}
}

func TestUpdateEvent(t *testing.T) {
	type test struct {
		update    konnekt.EventUpdate
		wantEvent konnekt.Event
		wantCode  string
	}

	updatedEvent := konnekt.EventUpdate{
		Title:       ptr.From("Updated Title"),
		Description: ptr.From("Updated Description"),
		FromDate:    baseEvent.FromDate.Add(1 * time.Hour),
		ToDate:      baseEvent.ToDate.Add(1 * time.Hour),
		Address: &konnekt.AddressUpdate{
			Country:     ptr.From("Updated Country"),
			City:        ptr.From("Updated City"),
			Street:      ptr.From("Updated Street"),
			HouseNumber: ptr.From("Updated House Number"),
		},
		GenreNames: []string{"Pop", "Indie"},
	}

	tests := map[string]test{
		"Full update": {
			update: updatedEvent,
			wantEvent: konnekt.Event{
				ID:          1,
				Title:       *updatedEvent.Title,
				Description: *updatedEvent.Description,
				FromDate:    updatedEvent.FromDate,
				ToDate:      updatedEvent.ToDate,
				Address: konnekt.Address{
					Country:     *updatedEvent.Address.Country,
					City:        *updatedEvent.Address.City,
					Street:      *updatedEvent.Address.Street,
					HouseNumber: *updatedEvent.Address.HouseNumber,
				},
				Genres: []konnekt.Genre{
					{ID: 3, Name: "Pop"},
					{ID: 4, Name: "Indie"},
				},
			},
			wantCode: "",
		},
		"Base Event Update": {
			update: konnekt.EventUpdate{
				Title:       updatedEvent.Title,
				Description: updatedEvent.Description,
				FromDate:    updatedEvent.FromDate,
				ToDate:      updatedEvent.ToDate,
			},
			wantEvent: konnekt.Event{
				ID:          1,
				Title:       *updatedEvent.Title,
				Description: *updatedEvent.Description,
				FromDate:    updatedEvent.FromDate,
				ToDate:      updatedEvent.ToDate,
				Genres:      baseGenres,
				Address:     baseAddress,
			},
			wantCode: "",
		},
		"Genres Update": {
			update: konnekt.EventUpdate{
				GenreNames: updatedEvent.GenreNames,
			},
			wantEvent: konnekt.Event{
				ID:          1,
				Title:       baseEvent.Title,
				Description: baseEvent.Description,
				FromDate:    baseEvent.FromDate,
				ToDate:      baseEvent.ToDate,
				Address:     baseAddress,
				Genres: []konnekt.Genre{
					{ID: 3, Name: "Pop"},
					{ID: 4, Name: "Indie"},
				},
			},
			wantCode: "",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			repo, dsn := MustOpenRepo(t)
			defer MustCloseRepo(t, repo, dsn)

			MustCreateEvent(t, context.Background(), repo, baseEvent)

			service := sqlite.NewEventService(repo)

			event, err := service.UpdateEvent(context.Background(), 1, tt.update)

			code := konnekt.ErrorCode(err)
			if code != tt.wantCode {
				t.Fatalf("got code %q, want code %q, error: %v", code, tt.wantCode, err)
			}

			if !event.Equals(tt.wantEvent) {
				t.Fatalf("got:\n%+v\nwant:\n%+v\n", event, tt.wantEvent)
			}
		})
	}
}

func TestFindEventByID(t *testing.T) {
	type test struct {
		id        int64
		wantEvent konnekt.Event
		wantCode  string
	}

	tests := map[string]test{
		"Valid ID": {
			id:        1,
			wantEvent: baseEvent,
			wantCode:  "",
		},
		"Invalid ID": {
			id:        999,
			wantEvent: konnekt.Event{},
			wantCode:  konnekt.ERRNOTFOUND,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			repo, dsn := MustOpenRepo(t)
			defer MustCloseRepo(t, repo, dsn)

			service := sqlite.NewEventService(repo)

			MustCreateEvent(t, context.Background(), repo, baseEvent)

			event, err := service.FindEventByID(context.Background(), tt.id)

			code := konnekt.ErrorCode(err)

			if code != tt.wantCode {
				t.Fatalf("got code %q, want code %q", code, tt.wantCode)
			}

			if !event.Equals(tt.wantEvent) {
				t.Fatalf("got event %+v, want event %+v", event, tt.wantEvent)
			}
		})
	}
}

func TestFindEvents(t *testing.T) {
	type test struct {
		filter     konnekt.EventFilter
		wantEvents []konnekt.Event
		wantCode   string
	}

	tests := map[string]test{
		"Valid ID": {
			filter:     konnekt.EventFilter{ID: ptr.From(int64(1))},
			wantEvents: []konnekt.Event{baseEvent},
			wantCode:   "",
		},
		"Invalid ID": {
			filter:     konnekt.EventFilter{ID: ptr.From(int64(999))},
			wantEvents: nil,
			wantCode:   "",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			repo, dsn := MustOpenRepo(t)
			defer MustCloseRepo(t, repo, dsn)

			MustCreateEvent(t, context.Background(), repo, baseEvent)

			service := sqlite.NewEventService(repo)

			gotEvents, err := service.FindEvents(context.Background(), tt.filter)
			code := konnekt.ErrorCode(err)

			if code != tt.wantCode {
				t.Fatalf("got code %q, want code %q, error: %v", code, tt.wantCode, err)
			}

			for i, gotEvent := range gotEvents {
				wantEvent := tt.wantEvents[i]
				if !gotEvent.Equals(wantEvent) {
					t.Fatalf("got %+v, want %+v", gotEvent, wantEvent)
				}
			}
		})
	}
}

func MustCreateEvent(tb testing.TB, ctx context.Context, repo *sqlite.Repository, event konnekt.Event) {
	tb.Helper()

	_, err := sqlite.NewEventService(repo).CreateEvent(ctx, event)
	if err != nil {
		tb.Fatal(err)
	}
}
