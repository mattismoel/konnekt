package query_test

import (
	"errors"
	"testing"

	"github.com/mattismoel/konnekt/internal/query"
)

func TestNewFilter(t *testing.T) {
	type test struct {
		cmp       query.Comparator
		wantCmp   query.Comparator
		value     string
		wantValue string
		wantErr   error
	}

	tests := map[string]test{
		"Valid filter": {
			cmp:       query.Equal,
			wantCmp:   query.Equal,
			value:     "2",
			wantValue: "2",
			wantErr:   nil,
		},
		"Empty comparator": {
			cmp:       query.Comparator(""),
			wantCmp:   "",
			value:     "2",
			wantValue: "",
			wantErr:   query.ErrFilterCmpInvalid,
		},
		"Empty value": {
			cmp:       query.Equal,
			wantCmp:   "",
			value:     "",
			wantValue: "",
			wantErr:   query.ErrFilterValueInvalid,
		},
		"Multiple comparator": {
			cmp:       query.Comparator(">=!="),
			wantCmp:   "",
			value:     "2",
			wantValue: "",
			wantErr:   query.ErrFilterCmpInvalid,
		},
		"Multiple values (space separated)": {
			cmp:       query.Equal,
			wantCmp:   "",
			value:     "2 4",
			wantValue: "",
			wantErr:   query.ErrFilterValueInvalid,
		},
		"Multiple values (comma separated)": {
			cmp:       query.Equal,
			wantCmp:   "",
			value:     "2,4",
			wantValue: "",
			wantErr:   query.ErrFilterValueInvalid,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			f, err := query.NewFilter(tt.cmp, tt.value)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got %v, want %v", err, tt.wantErr)
			}

			if f.Cmp != tt.wantCmp {
				t.Fatalf("got cmp %q, want %q", f.Cmp, tt.wantCmp)
			}

			if f.Value != tt.wantValue {
				t.Fatalf("got %q, want %q", f.Value, tt.wantValue)
			}
		})
	}
}

func TestFilterEqual(t *testing.T) {
	type test struct {
		f1        query.Filter
		f2        query.Filter
		wantEqual bool
	}

	tests := map[string]test{
		"Equal filters": {
			f1: query.Filter{
				Cmp:   "=",
				Value: "2",
			},
			f2: query.Filter{
				Cmp:   "=",
				Value: "2",
			},
			wantEqual: true,
		},
		"Cmp differs": {
			f1: query.Filter{
				Cmp:   "<=",
				Value: "2",
			},
			f2: query.Filter{
				Cmp:   "=",
				Value: "2",
			},
			wantEqual: false,
		},
		"Empty cmp (f1)": {
			f1: query.Filter{
				Cmp:   "",
				Value: "2",
			},
			f2: query.Filter{
				Cmp:   "=",
				Value: "2",
			},
			wantEqual: false,
		},
		"Empty cmp (f2)": {
			f1: query.Filter{
				Cmp:   "=",
				Value: "2",
			},
			f2: query.Filter{
				Cmp:   "",
				Value: "2",
			},
			wantEqual: false,
		},
		"Empty value (f1)": {
			f1: query.Filter{
				Cmp:   "=",
				Value: "",
			},
			f2: query.Filter{
				Cmp:   "=",
				Value: "2",
			},
			wantEqual: false,
		},
		"Empty value (f2)": {
			f1: query.Filter{
				Cmp:   "=",
				Value: "2",
			},
			f2: query.Filter{
				Cmp:   "=",
				Value: "",
			},
			wantEqual: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			isEqual := tt.f1.Equals(tt.f2)
			if isEqual != tt.wantEqual {
				t.Fatalf("got %v, want %v", isEqual, tt.wantEqual)
			}
		})
	}
}
