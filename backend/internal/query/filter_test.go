package query_test

import (
	"errors"
	"testing"

	"github.com/mattismoel/konnekt/internal/query"
)

func TestNewFilter(t *testing.T) {
	type test struct {
		cmp     query.Comparator
		value   string
		wantErr error
	}

	tests := map[string]test{
		"Valid filter": {
			cmp:     query.Equal,
			value:   "2",
			wantErr: nil,
		},
		"Empty comparator": {
			cmp:     query.Comparator(""),
			value:   "2",
			wantErr: query.ErrFilterCmpInvalid,
		},
		"Empty value": {
			cmp:     query.Equal,
			value:   "",
			wantErr: query.ErrFilterValueInvalid,
		},
		"Multiple comparator": {
			cmp:     query.Comparator(">=!="),
			value:   "2",
			wantErr: query.ErrFilterCmpInvalid,
		},
		"Multiple values (space separated)": {
			cmp:     query.Equal,
			value:   "2 4",
			wantErr: query.ErrFilterValueInvalid,
		},
		"Multiple values (comma separated)": {
			cmp:     query.Equal,
			value:   "2,4",
			wantErr: query.ErrFilterValueInvalid,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := query.NewFilter(tt.cmp, tt.value)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got %v, want %v", err, tt.wantErr)
			}
		})
	}
}
