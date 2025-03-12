package event

import (
	"errors"
	"github.com/mattismoel/konnekt/internal/domain/concert"
	"github.com/mattismoel/konnekt/internal/domain/venue"
	"net/http"
	"net/url"
	"strings"
)

var (
	ErrInvalidID            = errors.New("Event ID must be a positive integer")
	ErrEmptyTitle           = errors.New("Event title must not be empty")
	ErrEmptyDescription     = errors.New("Event description must not be empty")
	ErrInvalidImageURL      = errors.New("Event image URL must be valid")
	ErrImageURLInaccessible = errors.New("Image URL must be accessible")
)

type Event struct {
	ID          int64             `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	ImageURL    string            `json:"imageUrl"`
	Venue       venue.Venue       `json:"venue"`
	Concerts    []concert.Concert `json:"concerts"`
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

func WithID(id int64) CfgFunc {
	return func(e *Event) error {
		if id <= 0 {
			return ErrInvalidID
		}

		e.ID = id
		return nil
	}
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

func WithImageURL(u string) CfgFunc {
	return func(e *Event) error {
		url, err := url.ParseRequestURI(u)
		if err != nil {
			return ErrInvalidImageURL
		}

		resp, err := http.Get(url.String())
		if err != nil {
			return ErrImageURLInaccessible
		}

		// Check whether page is accessible to the end user.
		if resp.StatusCode < 200 || resp.StatusCode >= 400 {
			return ErrImageURLInaccessible
		}

		e.ImageURL = url.String()

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
