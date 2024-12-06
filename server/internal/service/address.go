package service

import (
	"errors"
	"strings"
)

var (
	ErrInvalidCountry     = errors.New("Invalid country")
	ErrInvalidCity        = errors.New("Invalid city")
	ErrInvalidStreet      = errors.New("Invalid street")
	ErrInvalidHouseNumber = errors.New("Invalid street number")
	ErrEmptyAddress       = errors.New("Address is empty")
)

type Address struct {
	Country     string `json:"country"`
	City        string `json:"city"`
	Street      string `json:"street"`
	HouseNumber string `json:"houseNumber"`
}

func (a Address) Validate() error {
	emptyAddr := Address{}

	if a == emptyAddr {
		return ErrEmptyAddress
	}

	if strings.TrimSpace(a.Country) == "" {
		return ErrInvalidCountry
	}

	if strings.TrimSpace(a.City) == "" {
		return ErrInvalidCity
	}

	if strings.TrimSpace(a.Street) == "" {
		return ErrInvalidStreet
	}
	if strings.TrimSpace(a.HouseNumber) == "" {
		return ErrInvalidHouseNumber
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
