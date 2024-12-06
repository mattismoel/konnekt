package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/mattismoel/konnekt"
	"github.com/mattismoel/konnekt/internal/storage"
)

const DATE_PRECISION = time.Minute

var (
	ErrNoTitle       = errors.New("No title")
	ErrNoDescription = errors.New("No description")
	ErrZeroFromDay   = errors.New("No from-date")
	ErrZeroToDay     = errors.New("No to-date")
	ErrNoGenres      = errors.New("No genres")
)

type Event struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	FromDate    time.Time `json:"fromDate"`
	ToDate      time.Time `json:"toDate"`
	Address     Address   `json:"address"`
	Genres      []string  `json:"genres"`
}

type EventFilter struct {
	ID         *int64    `json:"id"`
	ArtistName *string   `json:"artistName"`
	MinDate    time.Time `json:"minDate"`
	MaxDate    time.Time `json:"maxDate"`

	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type eventRepository interface {
	InsertEvent(ctx context.Context, baseEvent storage.BaseEvent, address storage.Address, genres []string) (storage.Event, error)
	DeleteEvent(ctx context.Context, id int64) error
	UpdateEvent(ctx context.Context, id int64, event storage.Event) (storage.Event, error)
	FindEventByID(ctx context.Context, id int64) (storage.Event, error)
	FindEvents(ctx context.Context, filter EventFilter) ([]storage.Event, error)
}

type eventService struct {
	repo eventRepository
}

func NewEventService(repo eventRepository) *eventService {
	return &eventService{repo: repo}
}

type CreateEventRequest struct {
	Title       string
	Description string
	FromDate    time.Time
	ToDate      time.Time
	Address     CreateAddressRequest
	Genres      []string
}

type CreateAddressRequest struct {
	Country     string
	City        string
	Street      string
	HouseNumber string
}

func (s eventService) CreateEvent(ctx context.Context, event Event) (Event, error) {
	err := event.Validate()
	if err != nil {
		return Event{}, konnekt.Errorf(konnekt.ERRINVALID, err.Error())
	}

	repoBaseEvent := storage.BaseEvent{
		Title:       event.Title,
		Description: event.Description,
		FromDate:    event.FromDate,
		ToDate:      event.ToDate,
	}

	fmt.Println(repoBaseEvent)

	repoAddress := storage.Address{
		Country:     event.Address.Country,
		City:        event.Address.City,
		Street:      event.Address.Street,
		HouseNumber: event.Address.HouseNumber,
	}

	repoEvent, err := s.repo.InsertEvent(ctx, repoBaseEvent, repoAddress, event.Genres)
	if err != nil {
		return Event{}, err
	}

	event.ID = repoEvent.ID

	return event, nil
}

func (s eventService) DeleteEvent(ctx context.Context, id int64) error {
	err := s.repo.DeleteEvent(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s eventService) FindEventByID(ctx context.Context, id int64) (Event, error) {
	repoEvent, err := s.repo.FindEventByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Event{}, konnekt.Errorf(konnekt.ERRNOTFOUND, "No events found")
		}

		return Event{}, err
	}

	genres := []string{}

	for _, genre := range repoEvent.Genres {
		genres = append(genres, genre.Name)
	}

	event := Event{
		ID:          repoEvent.ID,
		Title:       repoEvent.Title,
		Description: repoEvent.Description,
		FromDate:    repoEvent.FromDate,
		ToDate:      repoEvent.ToDate,
		Address: Address{
			Country:     repoEvent.Address.Country,
			City:        repoEvent.Address.City,
			Street:      repoEvent.Address.Street,
			HouseNumber: repoEvent.Address.HouseNumber,
		},
		Genres: genres,
	}

	return event, nil
}

func (s eventService) FindEvents(ctx context.Context, filter EventFilter) ([]Event, error) {
	repoEvents, err := s.repo.FindEvents(ctx, filter)
	if err != nil {
		return nil, err
	}

	if len(repoEvents) == 0 || repoEvents == nil {
		return []Event{}, konnekt.Errorf(konnekt.ERRNOTFOUND, "No events found")
	}

	events := []Event{}

	for _, repoEvent := range repoEvents {
		events = append(events, repoEventToEvent(repoEvent))
	}

	return events, err
}

func (s eventService) UpdateEvent(ctx context.Context, id int64, update Event) (Event, error) {
	repoEvent, err := s.repo.FindEventByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Event{}, konnekt.Errorf(konnekt.ERRNOTFOUND, "Event not found")
		}

		return Event{}, err
	}

	if err = update.Validate(); err != nil {
		return Event{}, err
	}

	repoEvent, err = s.repo.UpdateEvent(ctx, id, repoEvent)
	if err != nil {
		return Event{}, err
	}

	update.ID = repoEvent.ID

	return update, nil
}

func (e Event) Validate() error {
	if strings.TrimSpace(e.Title) == "" {
		return fmt.Errorf("Title must be set")
	}

	if strings.TrimSpace(e.Description) == "" {
		return fmt.Errorf("Description must be set")
	}

	if e.FromDate.IsZero() {
		return fmt.Errorf("FromDate must be set")
	}

	if e.ToDate.IsZero() {
		return fmt.Errorf("ToDate must be set")
	}

	if err := e.Address.Validate(); err != nil {
		return err
	}

	if e.Genres == nil || len(e.Genres) == 0 {
		return fmt.Errorf("Genre count must be at least 1")
	}

	return nil
}

func (e Event) Equals(a Event) bool {
	if e.Title != a.Title {
		return false
	}

	if e.Description != a.Description {
		return false
	}

	if !e.FromDate.Truncate(DATE_PRECISION).Equal(a.FromDate.Truncate(DATE_PRECISION)) {
		return false
	}

	if !e.ToDate.Truncate(DATE_PRECISION).Equal(a.ToDate.Truncate(DATE_PRECISION)) {
		return false
	}

	if !e.Address.Equals(a.Address) {
		return false
	}

	if !slices.Equal(e.Genres, a.Genres) {
		return false
	}

	return true
}

func repoEventToEvent(e storage.Event) Event {
	genres := []string{}

	for _, genre := range e.Genres {
		genres = append(genres, genre.Name)
	}

	return Event{
		ID:          e.ID,
		Title:       e.Title,
		Description: e.Description,
		FromDate:    e.FromDate,
		ToDate:      e.ToDate,
		Address: Address{
			Country:     e.Address.Country,
			City:        e.Address.City,
			Street:      e.Address.Street,
			HouseNumber: e.Address.HouseNumber,
		},
		Genres: genres,
	}
}
