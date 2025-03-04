package event

import (
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/mattismoel/konnekt/internal/domain/concert"
	"github.com/mattismoel/konnekt/internal/domain/venue"
)

var (
	ErrEmptyTitle                = errors.New("Event title must not be empty")
	ErrEmptyDescription          = errors.New("Event description must not be empty")
	ErrInvalidCoverImageURL      = errors.New("Event cover image URL must be valid")
	ErrCoverImageURLInaccessible = errors.New("Cover image URL must be accessible")
)

type Event struct {
	ID            int64             `json:"id"`
	Title         string            `json:"title"`
	Description   string            `json:"description"`
	CoverImageURL string            `json:"coverImageUrl"`
	Venue         venue.Venue       `json:"venue"`
	Concerts      []concert.Concert `json:"concerts"`
}

type CfgFunc func(e *Event) error

func (e *Event) WithCfgs(cfgs ...CfgFunc) error {
	for _, cfg := range cfgs {
		if err := cfg(e); err != nil {
			return err
		}
	}

	return nil
}

func NewEvent(cfgs ...CfgFunc) (*Event, error) {
	e := &Event{
		Concerts: make([]concert.Concert, 0),
	}

	if err := e.WithCfgs(cfgs...); err != nil {
		return &Event{}, err
	}

	return e, nil
}

func WithTitle(title string) CfgFunc {
	return func(e *Event) error {
		title = strings.TrimSpace(title)

		if title == "" {
			return ErrEmptyTitle
		}

		e.Title = title
		return nil
	}
}

func WithDescription(description string) CfgFunc {
	return func(e *Event) error {
		description = strings.TrimSpace(description)

		if description == "" {
			return ErrEmptyDescription
		}

		e.Description = description

		return nil
	}
}

func WithCoverImageURL(coverImageURL string) CfgFunc {
	return func(e *Event) error {
		url, err := url.ParseRequestURI(coverImageURL)
		if err != nil {
			return ErrInvalidCoverImageURL
		}

		resp, err := http.Get(url.String())
		if err != nil {
			return ErrCoverImageURLInaccessible
		}

		// Check whether page is accessible to the end user.
		if resp.StatusCode < 200 || resp.StatusCode >= 400 {
			return ErrCoverImageURLInaccessible
		}

		e.CoverImageURL = url.String()

		return nil
	}
}

func WithConcerts(concerts ...concert.Concert) CfgFunc {
	return func(e *Event) error {
		e.Concerts = append(e.Concerts, concerts...)
		return nil
	}
}

func WithVenue(v venue.Venue) CfgFunc {
	return func(e *Event) error {
		e.Venue = v
		return nil
	}
}
