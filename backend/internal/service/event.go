package service

import (
	"context"
	"time"

	"github.com/mattismoel/konnekt/internal/domain/artist"
	"github.com/mattismoel/konnekt/internal/domain/concert"
	"github.com/mattismoel/konnekt/internal/domain/event"
	"github.com/mattismoel/konnekt/internal/domain/venue"
	"github.com/mattismoel/konnekt/internal/object"
	"github.com/mattismoel/konnekt/internal/query"
)

type EventService struct {
	eventRepo  event.Repository
	artistRepo artist.Repository
	venueRepo  venue.Repository
}

func NewEventService(
	eventRepo event.Repository,
	artistRepo artist.Repository,
	venueRepo venue.Repository,
	objectStore object.Store,
) (*EventService, error) {
	return &EventService{
		eventRepo:  eventRepo,
		artistRepo: artistRepo,
		venueRepo:  venueRepo,
	}, nil
}

type CreateEvent struct {
	Title         string
	Description   string
	CoverImageURL string
	VenueID       int64
	Concerts      []CreateConcert
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
		event.WithVenue(venue),
		event.WithCoverImageURL(load.CoverImageURL),
		event.WithConcerts(concerts...),
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
	Title         string
	Description   string
	CoverImageURL string
	VenueID       int64
	Concerts      []UpdateConcert
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
		event.WithConcerts(concerts...),
		event.WithVenue(venue),
		event.WithCoverImageURL(prevEvent.CoverImageURL),
	)

	if err != nil {
		return event.Event{}, err
	}

	err = s.eventRepo.Update(ctx, eventID, *e)
	if err != nil {
		return event.Event{}, err
	}

	return *e, nil
}

type EventListQuery struct {
	query.ListQuery
	From time.Time
	To   time.Time
}

func NewEventListQuery(page, perPage, limit int, from, to time.Time) EventListQuery {
	if to.Before(from) {
		to = from
	}

	return EventListQuery{
		ListQuery: query.NewListQuery(page, perPage, limit),
		From:      from,
		To:        to,
	}
}

func (s EventService) List(ctx context.Context, q EventListQuery) (query.ListResult[event.Event], error) {
	events, totalCount, err := s.eventRepo.List(ctx, q.From, q.To, q.Offset(), q.Limit)
	if err != nil {
		return query.ListResult[event.Event]{}, err
	}

	return query.ListResult[event.Event]{
		Page:       q.Page,
		PerPage:    q.PerPage,
		TotalCount: totalCount,
		PageCount:  q.PageCount(totalCount),
		Records:    events,
	}, nil
}
