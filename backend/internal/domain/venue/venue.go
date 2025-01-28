package venue

import "strings"

type Venue struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	CountryCode string `json:"countryCode"`
	City        string `json:"city"`
}

func NewVenue(name, countryCode, city string) (Venue, error) {
	return Venue{
		Name:        strings.TrimSpace(name),
		CountryCode: strings.TrimSpace(countryCode),
		City:        strings.TrimSpace(city),
	}, nil
}
