package service

import (
	"context"
	"fmt"
	"time"

	"github.com/mattismoel/konnekt/internal/domain/artist"
	"github.com/mattismoel/konnekt/internal/domain/concert"
	"github.com/mattismoel/konnekt/internal/domain/event"
	"github.com/mattismoel/konnekt/internal/domain/venue"
)

type EventService struct {
	eventRepo  event.Repository
	artistRepo artist.Repository
	venueRepo  venue.Repository
}

func NewEventService(eventRepo event.Repository, artistRepo artist.Repository, venueRepo venue.Repository) (*EventService, error) {
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


func (s EventService) Create(ctx context.Context, load CreateEvent) (int64, error) {
	e, err := event.NewEvent(load.Title, load.Description, load.CoverImageURL)
	if err != nil {
		return 0, err
	}

	venue, err := s.venueRepo.ByID(ctx, load.VenueID)
	if err != nil {
		return 0, err
	}

	e.WithVenue(venue)

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

	e.WithConcerts(concerts...)

	eventID, err := s.eventRepo.Insert(ctx, e)
	if err != nil {
		return 0, err
	}

	return eventID, nil
}

type EventListResult struct {
	Page       int
	PerPage    int
	PageCount  int
	TotalCount int
	Events     []event.Event
}

func (s EventService) List(ctx context.Context, page, perPage int) (EventListResult, error) {
	offset := (page - 1) * perPage

	events, totalCount, err := s.eventRepo.List(ctx, perPage, offset)
	if err != nil {
		return EventListResult{}, err
	}

	fmt.Printf("Total: %d, Page: %d, PerPage: %d\n", totalCount, page, perPage)

	pageCount := (totalCount + perPage - 1) / perPage

	return EventListResult{
		Page:       page,
		PerPage:    perPage,
		TotalCount: totalCount,
		PageCount:  pageCount,
		Events:     events,
	}, nil
}
