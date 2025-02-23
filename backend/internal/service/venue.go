package service

import (
	"github.com/mattismoel/konnekt/internal/domain/venue"
)

type VenueService struct {
	venueRepo venue.Repository
}

func NewVenueService(venueRepo venue.Repository) *VenueService {
	return &VenueService{venueRepo: venueRepo}
}

