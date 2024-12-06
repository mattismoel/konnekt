package service_test

import (
	"errors"
	"testing"

	"github.com/mattismoel/konnekt/internal/service"
)

var baseAddress = service.Address{
	Country:     "Denmark",
	City:        "Odense",
	Street:      "Postenvej",
	HouseNumber: "18B",
}

func TestAddressValidate(t *testing.T) {
	type updater func(service.Address) service.Address

	type test struct {
		updater updater
		err     error
	}

	tests := map[string]test{
		"Valid Address": {
			updater: nil,
			err:     nil,
		},
		"No Country": {
			updater: func(a service.Address) service.Address {
				a.Country = " "
				return a
			},
			err: service.ErrInvalidCountry,
		},
		"No City": {
			updater: func(a service.Address) service.Address {
				a.City = " "
				return a
			},
			err: service.ErrInvalidCity,
		},
		"No Street": {
			updater: func(a service.Address) service.Address {
				a.Street = " "
				return a
			},
			err: service.ErrInvalidStreet,
		},
		"No House Number": {
			updater: func(a service.Address) service.Address {
				a.HouseNumber = " "
				return a
			},
			err: service.ErrInvalidHouseNumber,
		},
		"Empty": {
			updater: func(a service.Address) service.Address {
				return service.Address{}
			},
			err: service.ErrEmptyAddress,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			address := baseAddress

			if tt.updater != nil {
				address = tt.updater(address)
			}

			err := address.Validate()

			if !errors.Is(err, tt.err) {
				t.Fatalf("got %v, want %v", err, tt.err)
			}
		})
	}
}

func TestAddressEquals(t *testing.T) {
	type updater func(service.Address) service.Address

	type test struct {
		aUpdater   updater
		bUpdater   updater
		wantEquals bool
	}

	tests := map[string]test{
		"Equal": {
			aUpdater:   nil,
			bUpdater:   nil,
			wantEquals: true,
		},
		"Country differ": {
			aUpdater: nil,
			bUpdater: func(a service.Address) service.Address {
				a.Country = "Sweden"
				return a
			},
			wantEquals: false,
		},
		"City differ": {
			aUpdater: nil,
			bUpdater: func(a service.Address) service.Address {
				a.City = "Stockholm"
				return a
			},
			wantEquals: false,
		},
		"Street differ": {
			aUpdater: nil,
			bUpdater: func(a service.Address) service.Address {
				a.Street = "Otherstreet"
				return a
			},
			wantEquals: false,
		},
		"House Number differ": {
			aUpdater: nil,
			bUpdater: func(a service.Address) service.Address {
				a.HouseNumber = "1C"
				return a
			},
			wantEquals: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			a, b := baseAddress, baseAddress

			if tt.aUpdater != nil {
				a = tt.aUpdater(a)
			}

			if tt.bUpdater != nil {
				b = tt.bUpdater(b)
			}

			gotEquals := a.Equals(b)

			if gotEquals != tt.wantEquals {
				t.Fatalf("got %v, want %v", gotEquals, tt.wantEquals)
			}
		})
	}
}
