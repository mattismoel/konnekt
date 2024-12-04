package konnekt_test

import (
	"testing"

	"github.com/mattismoel/konnekt"
)

type addressUpdaterFunc func(konnekt.Address) konnekt.Address

var baseAddress = konnekt.Address{
	Country:     "Denmark",
	City:        "Odense",
	Street:      "Postenvej",
	HouseNumber: "18B",
}

func TestAddressValidate(t *testing.T) {
	type test struct {
		updater  addressUpdaterFunc
		wantCode string
	}

	tests := map[string]test{
		"Valid Address": {
			updater:  nil,
			wantCode: "",
		},
		"No Country": {
			updater: func(a konnekt.Address) konnekt.Address {
				a.Country = " "
				return a
			},
			wantCode: konnekt.ERRINVALID,
		},
		"No City": {
			updater: func(a konnekt.Address) konnekt.Address {
				a.City = " "
				return a
			},
			wantCode: konnekt.ERRINVALID,
		},
		"No Street": {
			updater: func(a konnekt.Address) konnekt.Address {
				a.Street = " "
				return a
			},
			wantCode: konnekt.ERRINVALID,
		},
		"No House Number": {
			updater: func(a konnekt.Address) konnekt.Address {
				a.HouseNumber = " "
				return a
			},
			wantCode: konnekt.ERRINVALID,
		},
		"Empty": {
			updater: func(a konnekt.Address) konnekt.Address {
				return konnekt.Address{}
			},
			wantCode: konnekt.ERRINVALID,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			address := baseAddress

			if tt.updater != nil {
				address = tt.updater(address)
			}

			err := address.Validate()
			code := konnekt.ErrorCode(err)

			if code != tt.wantCode {
				t.Fatalf("got code %q, want code %q, error: %v", code, tt.wantCode, err)
			}
		})
	}
}

func TestAddressEquals(t *testing.T) {
	type test struct {
		aUpdater   addressUpdaterFunc
		bUpdater   addressUpdaterFunc
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
			bUpdater: func(a konnekt.Address) konnekt.Address {
				a.Country = "Sweden"
				return a
			},
			wantEquals: false,
		},
		"City differ": {
			aUpdater: nil,
			bUpdater: func(a konnekt.Address) konnekt.Address {
				a.City = "Stockholm"
				return a
			},
			wantEquals: false,
		},
		"Street differ": {
			aUpdater: nil,
			bUpdater: func(a konnekt.Address) konnekt.Address {
				a.Street = "Otherstreet"
				return a
			},
			wantEquals: false,
		},
		"House Number differ": {
			aUpdater: nil,
			bUpdater: func(a konnekt.Address) konnekt.Address {
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
