package service

import (
	"context"
	"fmt"

	"github.com/mattismoel/konnekt/internal/domain/venue"
	"github.com/mattismoel/konnekt/internal/query"
)

type VenueService struct {
	venueRepo venue.Repository
}

func NewVenueService(venueRepo venue.Repository) *VenueService {
	return &VenueService{venueRepo: venueRepo}
}

func (s VenueService) List(ctx context.Context, q venue.Query) (query.ListResult[venue.Venue], error) {
	result, err := s.venueRepo.List(ctx, q)
	if err != nil {
		return query.ListResult[venue.Venue]{}, fmt.Errorf("Could not list venues: %v", err)
	}

	return result, nil
}

func (s VenueService) Create(ctx context.Context, v venue.Venue) (int64, error) {
	venueID, err := s.venueRepo.Insert(ctx, v)
	if err != nil {
		return 0, fmt.Errorf("Could not insert venue into repository: %v", err)
	}

	return venueID, nil
}

func (s VenueService) Delete(ctx context.Context, venueID int64) error {
	err := s.venueRepo.Delete(ctx, venueID)
	if err != nil {
		return fmt.Errorf("Could not delete venue from repository %d: %v", venueID, err)
	}

	return nil
}

func (s VenueService) Update(ctx context.Context, venueID int64, v venue.Venue) (venue.Venue, error) {
	if err := s.venueRepo.Update(ctx, venueID, v); err != nil {
		return venue.Venue{}, fmt.Errorf("Could not update venue %d: %v", venueID, err)
	}

	updatedVenue, err := s.venueRepo.ByID(ctx, venueID)
	if err != nil {
		return venue.Venue{}, fmt.Errorf("Could not get updated venue: %v", err)
	}

	return updatedVenue, nil
}

func (s VenueService) ByID(ctx context.Context, venueID int64) (venue.Venue, error) {
	v, err := s.venueRepo.ByID(ctx, venueID)
	if err != nil {
		return venue.Venue{}, fmt.Errorf("Could not get venue by id %d: %v", venueID, err)
	}

	return v, nil
}
