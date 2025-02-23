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

type VenueListQuery struct {
	query.ListQuery
}

func (s VenueService) List(ctx context.Context, q VenueListQuery) (query.ListResult[venue.Venue], error) {
	venues, totalCount, err := s.venueRepo.List(ctx, q.Limit, q.Offset())
	if err != nil {
		return query.ListResult[venue.Venue]{}, err
	}

	return query.ListResult[venue.Venue]{
		Page:       q.Page,
		PerPage:    q.PerPage,
		TotalCount: totalCount,
		PageCount:  q.PageCount(totalCount),
		Records:    venues,
	}, nil
}
