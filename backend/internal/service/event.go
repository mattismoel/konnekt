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
	VenueID     int64
	Concerts    []CreateConcert
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

func (s EventService) Create(ctx context.Context, load CreateEvent) (int64, error) {
	venue, err := s.venueRepo.ByID(ctx, load.VenueID)
	if err != nil {
		return 0, err
	}

	concerts := make([]concert.Concert, 0)
	for _, loadConcert := range load.Concerts {
		artist, err := s.artistRepo.ByID(ctx, loadConcert.ArtistID)
		if err != nil {
			return 0, err
		}

		c, err := concert.NewConcert(artist, loadConcert.From, loadConcert.To)
		if err != nil {
			return 0, err
		}

		concerts = append(concerts, c)
	}

	e, err := event.NewEvent(
		event.WithTitle(load.Title),
		event.WithDescription(load.Description),
		event.WithCoverImageURL(load.CoverImageURL),
		event.WithVenue(venue),
		event.WithConcerts(concerts...),
	)

	if err != nil {
		return 0, err
	}

	eventID, err := s.eventRepo.Insert(ctx, *e)
	if err != nil {
		return 0, err
	}

	return eventID, nil
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
