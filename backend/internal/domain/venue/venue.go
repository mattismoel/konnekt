package venue

import (
	"strings"

	"github.com/mattismoel/konnekt/internal/query"
)

type Venue struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	CountryCode string `json:"countryCode"`
	City        string `json:"city"`
}

type Query struct {
	query.ListQuery
}

func NewVenue(name, countryCode, city string) (Venue, error) {
	return Venue{
		Name:        strings.TrimSpace(name),
		CountryCode: strings.TrimSpace(countryCode),
		City:        strings.TrimSpace(city),
	}, nil
}
