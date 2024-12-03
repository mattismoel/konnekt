package konnekt

import (
	"strings"
)

type Address struct {
	Country     string `json:"country"`
	City        string `json:"city"`
	Street      string `json:"street"`
	HouseNumber string `json:"houseNumber"`
}

type AddressUpdate struct {
	Country     *string `json:"country"`
	City        *string `json:"city"`
	Street      *string `json:"street"`
	HouseNumber *string `json:"houseNumber"`
}

func (a Address) Validate() error {
	if strings.TrimSpace(a.Country) == "" {
		return Errorf(ERRINVALID, "Country must be set")
	}

	if strings.TrimSpace(a.City) == "" {
		return Errorf(ERRINVALID, "City must be set")
	}

	if strings.TrimSpace(a.Street) == "" {
		return Errorf(ERRINVALID, "Street must be set")
	}

	if strings.TrimSpace(a.HouseNumber) == "" {
		return Errorf(ERRINVALID, "House number must be set")
	}

	return nil
}

func (a Address) Equals(b Address) bool {
	if a.City != b.City {
		return false
	}

	if a.Street != b.Street {
		return false
	}

	if a.Country != b.Country {
		return false
	}

	if a.HouseNumber != b.HouseNumber {
		return false
	}

	return true
}
