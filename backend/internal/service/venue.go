package service

import (
	"context"

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
		return query.ListResult[venue.Venue]{}, err
	}

	return result, nil
}

type CreateVenue struct {
	Name        string
	City        string
	CountryCode string
}

func (s VenueService) Create(ctx context.Context, load CreateVenue) (int64, error) {
	v, err := venue.NewVenue(load.Name, load.CountryCode, load.City)
	if err != nil {
		return 0, err
	}

	venueID, err := s.venueRepo.Insert(ctx, v)
	if err != nil {
		return 0, err
	}

	return venueID, nil
}

func (s VenueService) Delete(ctx context.Context, venueID int64) error {
	err := s.venueRepo.Delete(ctx, venueID)
	if err != nil {
		return err
	}

	return nil
}

type UpdateVenue struct {
	Name        string
	City        string
	CountryCode string
}

func (s VenueService) Update(ctx context.Context, id int64, load UpdateVenue) (venue.Venue, error) {
	v, err := venue.NewVenue(load.Name, load.CountryCode, load.City)
	if err != nil {
		return venue.Venue{}, err
	}

	if err := s.venueRepo.Update(ctx, id, v); err != nil {
		return venue.Venue{}, err
	}

	updatedVenue, err := s.venueRepo.ByID(ctx, id)
	if err != nil {
		return venue.Venue{}, err
	}

	return updatedVenue, nil
}
