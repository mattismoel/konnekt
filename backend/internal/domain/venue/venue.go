package venue

import (
	"errors"
	"strings"

	"github.com/mattismoel/konnekt/internal/query"
)

var (
	ErrInvalidID          = errors.New("Venue ID must be valid")
	ErrInvalidName        = errors.New("Venue name must be valid")
	ErrInvalidCity        = errors.New("City must be valid")
	ErrInvalidCountryCode = errors.New("Country code must be valid")
)

type cfgFunc func(v *Venue) error

type Venue struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	CountryCode string `json:"countryCode"`
	City        string `json:"city"`
}
type Query struct {
	query.ListQuery
}

func NewVenue(cfgs ...cfgFunc) (Venue, error) {
	v := Venue{}

	err := v.WithCfgs(cfgs...)
	if err != nil {
		return Venue{}, err
	}

	return v, nil
}

func (v *Venue) WithCfgs(cfgs ...cfgFunc) error {
	for _, cfg := range cfgs {
		if err := cfg(v); err != nil {
			return err
		}
	}

	return nil
}

func WithID(id int64) cfgFunc {
	return func(v *Venue) error {
		if id <= 0 {
			return ErrInvalidID
		}

		v.ID = id
		return nil
	}
}

func WithName(name string) cfgFunc {
	return func(v *Venue) error {
		if strings.TrimSpace(name) == "" {
			return ErrInvalidName
		}

		v.Name = name
		return nil
	}
}

func WithCity(city string) cfgFunc {
	return func(v *Venue) error {
		if strings.TrimSpace(city) == "" {
			return ErrInvalidCity
		}

		v.City = city
		return nil
	}
}

func WithCountryCode(countryCode string) cfgFunc {
	return func(v *Venue) error {
		if strings.TrimSpace(countryCode) == "" {
			return ErrInvalidCountryCode
		}

		v.CountryCode = countryCode
		return nil
	}
}
