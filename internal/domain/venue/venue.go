package venue

import "strings"

type Venue struct {
	ID          int64
	Name        string
	CountryCode string
	City        string
}

func NewVenue(name, countryCode, city string) (Venue, error) {
	return Venue{
		Name:        strings.TrimSpace(name),
		CountryCode: strings.TrimSpace(countryCode),
		City:        strings.TrimSpace(city),
	}, nil
}
