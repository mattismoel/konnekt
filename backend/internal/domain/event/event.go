package event

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/mattismoel/konnekt/internal/domain/concert"
	"github.com/mattismoel/konnekt/internal/domain/venue"
)

var (
	ErrEmptyTitle           = errors.New("Event title must not be empty")
	ErrEmptyDescription     = errors.New("Event description must not be empty")
	ErrInvalidCoverImageURL = errors.New("Event cover image URL must be valid")
)

type Event struct {
	ID            int64             `json:"id"`
	Title         string            `json:"title"`
	Description   string            `json:"description"`
	CoverImageURL string            `json:"coverImageUrl"`
	Venue         venue.Venue       `json:"venue"`
	Concerts      []concert.Concert `json:"concerts"`
}

type Query struct {
	Limit   int
	Page    int
	PerPage int

	From time.Time
	To   time.Time
}

func (q Query) Offset() int {
	return (q.Page - 1) * q.PerPage
}

func NewEvent(title string, description string, coverImageURL string) (Event, error) {
	resp, err := http.Get(coverImageURL)
	if err != nil {
		return Event{}, ErrInvalidCoverImageURL
	}

	// Check whether page is accessible to the end user.
	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return Event{}, ErrInvalidCoverImageURL
	}

	return Event{
		Title:         strings.TrimSpace(title),
		Description:   strings.TrimSpace(description),
		CoverImageURL: strings.TrimSpace(coverImageURL),
		Concerts:      make([]concert.Concert, 0),
	}, nil
}

func (e *Event) WithConcerts(concerts ...concert.Concert) {
	e.Concerts = append(e.Concerts, concerts...)
}

func (e *Event) WithVenue(v venue.Venue) {
	e.Venue = v
}
