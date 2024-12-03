package konnekt_test

import (
	"testing"

	"github.com/mattismoel/konnekt"
)

var baseAddress = konnekt.Address{
	Country:     "Denmark",
	City:        "Odense",
	Street:      "Postenvej",
	HouseNumber: "18B",
}

func TestAddressValidate(t *testing.T) {
	type test struct {
		a        konnekt.Address
		wantCode string
	}

	tests := map[string]test{
		"Valid Address": {
			a:        baseAddress,
			wantCode: "",
		},
		"No Country": {
			a: konnekt.Address{
				Country:     " ",
				City:        baseAddress.City,
				Street:      baseAddress.Street,
				HouseNumber: baseAddress.HouseNumber,
			},
			wantCode: konnekt.ERRINVALID,
		},
		"No City": {
			a: konnekt.Address{
				Country:     baseAddress.Country,
				City:        " ",
				Street:      baseAddress.Street,
				HouseNumber: baseAddress.HouseNumber,
			},
			wantCode: konnekt.ERRINVALID,
		},
		"No Street": {
			a: konnekt.Address{
				Country:     baseAddress.Country,
				City:        baseAddress.City,
				Street:      " ",
				HouseNumber: baseAddress.HouseNumber,
			},
			wantCode: konnekt.ERRINVALID,
		},
		"No House Number": {
			a: konnekt.Address{
				Country:     baseAddress.Country,
				City:        baseAddress.City,
				Street:      baseAddress.Street,
				HouseNumber: " ",
			},
			wantCode: konnekt.ERRINVALID,
		},
		"Empty": {
			a:        konnekt.Address{},
			wantCode: konnekt.ERRINVALID,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			err := tt.a.Validate()
			code := konnekt.ErrorCode(err)

			if code != tt.wantCode {
				t.Fatalf("got code %q, want code %q, error: %v", code, tt.wantCode, err)
			}
		})
	}
}

func TestAddressEquals(t *testing.T) {
	type test struct {
		a          konnekt.Address
		b          konnekt.Address
		wantEquals bool
	}

	tests := map[string]test{
		"Equal": {
			a:          baseAddress,
			b:          baseAddress,
			wantEquals: true,
		},
		"Country differ": {
			a: baseAddress,
			b: konnekt.Address{
				Country:     "Sweden",
				City:        baseAddress.City,
				Street:      baseAddress.Street,
				HouseNumber: baseAddress.HouseNumber,
			},
			wantEquals: false,
		},
		"City differ": {
			a: baseAddress,
			b: konnekt.Address{
				Country:     baseAddress.Country,
				City:        "Berlin",
				Street:      baseAddress.Street,
				HouseNumber: baseAddress.HouseNumber,
			},
			wantEquals: false,
		},
		"Street differ": {
			a: baseAddress,
			b: konnekt.Address{
				Country:     baseAddress.Country,
				City:        baseAddress.City,
				Street:      "Vegavej",
				HouseNumber: baseAddress.HouseNumber,
			},
			wantEquals: false,
		},
		"House Number differ": {
			a: baseAddress,
			b: konnekt.Address{
				Country:     baseAddress.Country,
				City:        baseAddress.City,
				Street:      baseAddress.Street,
				HouseNumber: "1C",
			},
			wantEquals: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			gotEquals := tt.a.Equals(tt.b)
			if gotEquals != tt.wantEquals {
				t.Fatalf("got %v, want %v", gotEquals, tt.wantEquals)
			}
		})
	}
}
