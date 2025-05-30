package service

import (
	"context"
	"image"
	"io"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/mattismoel/konnekt/internal/domain/artist"
	"github.com/mattismoel/konnekt/internal/domain/concert"
	"github.com/mattismoel/konnekt/internal/domain/event"
	"github.com/mattismoel/konnekt/internal/domain/venue"
	"github.com/mattismoel/konnekt/internal/object"
	"github.com/mattismoel/konnekt/internal/query"
	"github.com/nfnt/resize"
)

const EVENT_COVER_IMAGE_WIDTH_PX = 2048

type EventService struct {
	eventRepo   event.Repository
	artistRepo  artist.Repository
	venueRepo   venue.Repository
	objectStore object.Store
}

func NewEventService(
	eventRepo event.Repository,
	artistRepo artist.Repository,
	venueRepo venue.Repository,
	objectStore object.Store,
) (*EventService, error) {
	return &EventService{
		eventRepo:   eventRepo,
		artistRepo:  artistRepo,
		venueRepo:   venueRepo,
		objectStore: objectStore,
	}, nil
}

type CreateEvent struct {
	Title       string
	Description string
	TicketURL   string
	ImageURL    string
	VenueID     int64
	Concerts    []CreateConcert
	IsPublic    bool
}

type CreateConcert struct {
	ArtistID int64
	From     time.Time
	To       time.Time
}

func (s EventService) ByID(ctx context.Context, eventID int64) (event.Event, error) {
	e, err := s.eventRepo.ByID(ctx, eventID)
	if err != nil {
		return event.Event{}, err
	}

	return e, nil
}

func (s EventService) Create(ctx context.Context, load CreateEvent) (event.Event, error) {
	venue, err := s.venueRepo.ByID(ctx, load.VenueID)
	if err != nil {
		return event.Event{}, err
	}

	concerts := make([]concert.Concert, 0)
	for _, c := range load.Concerts {
		artist, err := s.artistRepo.ByID(ctx, c.ArtistID)
		if err != nil {
			return event.Event{}, err
		}

		c, err := concert.NewConcert(
			concert.WithArtist(artist),
			concert.WithFrom(c.From),
			concert.WithTo(c.To),
		)

		if err != nil {
			return event.Event{}, err
		}

		concerts = append(concerts, c)
	}

	e, err := event.NewEvent(
		event.WithTitle(load.Title),
		event.WithDescription(load.Description),
		event.WithTicketURL(load.TicketURL),
		event.WithVenue(venue),
		event.WithImageURL(load.ImageURL),
		event.WithConcerts(concerts...),
		event.WithIsPublic(load.IsPublic),
	)

	if err != nil {
		return event.Event{}, err
	}

	eventID, err := s.eventRepo.Insert(ctx, *e)
	if err != nil {
		return event.Event{}, err
	}

	createdEvent, err := s.eventRepo.ByID(ctx, eventID)
	if err != nil {
		return event.Event{}, err
	}

	return createdEvent, nil
}

type UpdateConcert struct {
	ArtistID int64
	From     time.Time
	To       time.Time
}

type UpdateEvent struct {
	Title       string
	Description string
	TicketURL   string
	ImageURL    string
	VenueID     int64
	Concerts    []UpdateConcert
	IsPublic    bool
}

func (s EventService) Update(ctx context.Context, eventID int64, load UpdateEvent) (event.Event, error) {
	// Return if event does not exist.
	prevEvent, err := s.eventRepo.ByID(ctx, eventID)
	if err != nil {
		return event.Event{}, err
	}

	venue, err := s.venueRepo.ByID(ctx, load.VenueID)
	if err != nil {
		return event.Event{}, err
	}

	concerts := make([]concert.Concert, 0)
	for _, c := range load.Concerts {
		artist, err := s.artistRepo.ByID(ctx, c.ArtistID)
		if err != nil {
			return event.Event{}, err
		}

		concert, err := concert.NewConcert(
			concert.WithID(eventID),
			concert.WithArtist(artist),
			concert.WithFrom(c.From),
			concert.WithTo(c.To),
		)

		if err != nil {
			return event.Event{}, err
		}

		concerts = append(concerts, concert)
	}

	e, err := event.NewEvent(
		event.WithID(eventID),
		event.WithTitle(load.Title),
		event.WithDescription(load.Description),
		event.WithTicketURL(load.TicketURL),
		event.WithConcerts(concerts...),
		event.WithVenue(venue),
		event.WithIsPublic(load.IsPublic),
	)

	// If there is a cover image URL update, set it.
	if strings.TrimSpace(load.ImageURL) != "" {
		url, err := url.Parse(prevEvent.ImageURL)
		if err != nil {
			return event.Event{}, err
		}

		if err := s.objectStore.Delete(ctx, url.Path); err != nil {
			return event.Event{}, err
		}

		if err := e.WithCfgs(event.WithImageURL(load.ImageURL)); err != nil {
			return event.Event{}, err
		}
	}

	if err != nil {
		return event.Event{}, err
	}

	err = s.eventRepo.Update(ctx, eventID, *e)
	if err != nil {
		return event.Event{}, err
	}

	updatedEvent, err := s.eventRepo.ByID(ctx, eventID)
	if err != nil {
		return event.Event{}, err
	}

	return updatedEvent, nil
}

func (s EventService) UploadImage(ctx context.Context, r io.Reader) (string, error) {
	img, _, err := image.Decode(r)

	fileName := createRandomImageFileName("jpeg")

	if img.Bounds().Max.X > EVENT_COVER_IMAGE_WIDTH_PX {
		img = resize.Resize(EVENT_COVER_IMAGE_WIDTH_PX, 0, img, resize.Lanczos2)

	}

	formattedImg, err := formatJPEG(img)
	if err != nil {
		return "", err
	}

	url, err := s.objectStore.Upload(ctx, path.Join("/events", fileName), formattedImg)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (s EventService) List(ctx context.Context, q query.ListQuery) (query.ListResult[event.Event], error) {
	result, err := s.eventRepo.List(ctx, q)
	if err != nil {
		return query.ListResult[event.Event]{}, err
	}

	return result, nil
}

func (s EventService) Delete(ctx context.Context, eventID int64) error {
	e, err := s.eventRepo.ByID(ctx, eventID)
	if err != nil {
		return err
	}

	url, err := url.Parse(e.ImageURL)
	if err != nil {
		return err
	}

	err = s.objectStore.Delete(ctx, url.Path)
	if err != nil {
		return err
	}

	err = s.eventRepo.Delete(ctx, eventID)
	if err != nil {
		return err
	}

	return nil
}
