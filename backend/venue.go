package konnekt

import (
	"context"

	"github.com/mattismoel/konnekt/backend/api"
)

type Venue struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	City    string `json:"city"`
	Country string `json:"country"`
}

type CreateVenue struct {
	Name      string `json:"name"`
	City      string `json:"city"`
	Country   string `json:"country"`
	CreatedBy int64  `json:"-"`
}

type VenueRepo interface {
	InsertVenue(context.Context, CreateVenue) (int64, error)
	VenueByID(context.Context, int64) (Venue, error)
	ListVenues(context.Context, api.ListRequest) (api.ListResponse[Venue], error)
}
