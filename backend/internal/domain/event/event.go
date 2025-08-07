package event

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/mattismoel/konnekt/internal/cfg"
	"github.com/mattismoel/konnekt/internal/domain/concert"
	"github.com/mattismoel/konnekt/internal/domain/venue"
)

var (
	ErrInvalidID            = errors.New("Event ID must be a positive integer")
	ErrInvalidTitle         = errors.New("Event title must not be valid")
	ErrInvalidDescription   = errors.New("Event description must be valid")
	ErrInvalidImageURL      = errors.New("Event image URL must be valid")
	ErrImageURLInaccessible = errors.New("Image URL must be accessible")

	ErrTicketURLInvalid      = errors.New("Ticket URL must be valid")
	ErrTicketURLInaccessible = errors.New("Ticket URL must be accessible")
	ErrNoConcerts            = errors.New("Event must have at least one concert")
)

type Event struct {
	ID          int64             `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	TicketURL   string            `json:"ticketUrl"`
	ImageURL    string            `json:"imageUrl"`
	Venue       venue.Venue       `json:"venue"`
	Concerts    []concert.Concert `json:"concerts"`
	IsPublic    bool              `json:"isPublic"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedBy int64     `json:"createdBy"`
	UpdatedBy int64     `json:"updatedBy"`
}

func (e *Event) WithCfgs(cfgs ...cfg.Func[Event]) error {
	for _, cfg := range cfgs {
		if err := cfg(e); err != nil {
			return err
		}
	}

	return nil
}

func NewEvent(cfgs ...cfg.Func[Event]) (Event, error) {
	e := Event{
		Concerts: make([]concert.Concert, 0),
		IsPublic: false,
	}

	if err := e.WithCfgs(cfgs...); err != nil {
		return Event{}, err
	}

	return e, nil
}

func WithID(id int64) cfg.Func[Event] {
	return func(e *Event) error {
		if id <= 0 {
			return ErrInvalidID
		}

		e.ID = id
		return nil
	}
}

func WithTitle(title string) cfg.Func[Event] {
	return func(e *Event) error {
		title = strings.TrimSpace(title)

		if title == "" {
			return ErrInvalidTitle
		}

		e.Title = title
		return nil
	}
}

func WithDescription(description string) cfg.Func[Event] {
	return func(e *Event) error {
		description = strings.TrimSpace(description)

		if description == "" {
			return ErrInvalidDescription
		}

		e.Description = description

		return nil
	}
}

func WithTicketURL(u string) cfg.Func[Event] {
	return func(e *Event) error {
		url, err := url.ParseRequestURI(u)
		if err != nil {
			return ErrTicketURLInvalid
		}

		resp, err := http.Get(url.String())
		if err != nil {
			return ErrTicketURLInaccessible
		}

		if !(resp.StatusCode >= 200) || !(resp.StatusCode < 400) {
			return ErrTicketURLInaccessible
		}

		e.TicketURL = u
		return nil
	}
}

func WithImageURL(u string) cfg.Func[Event] {
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

func WithConcerts(concerts ...concert.Concert) cfg.Func[Event] {
	return func(e *Event) error {
		if len(concerts) <= 0 {
			return ErrNoConcerts
		}

		e.Concerts = concerts
		return nil
	}
}

func WithVenue(v venue.Venue) cfg.Func[Event] {
	return func(e *Event) error {
		e.Venue = v
		return nil
	}
}

func WithIsPublic(isPublic bool) cfg.Func[Event] {
	return func(e *Event) error {
		e.IsPublic = isPublic
		return nil
	}
}

func WithCreatedBy(memberID int64) cfg.Func[Event] {
	return func(e *Event) error {
		e.CreatedBy = memberID
		return nil
	}
}

func WithUpdatedBy(memberID int64) cfg.Func[Event] {
	return func(e *Event) error {
		e.UpdatedBy = memberID
		return nil
	}
}
