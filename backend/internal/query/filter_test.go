package query_test

import (
	"errors"
	"testing"

	"github.com/mattismoel/konnekt/internal/query"
)

func TestNewFilter(t *testing.T) {
	type test struct {
		key     string
		cmp     query.Comparator
		value   string
		wantErr error
	}

	tests := map[string]test{
		"Valid filter": {
			key:     "key_a",
			cmp:     query.Equal,
			value:   "2",
			wantErr: nil,
		},
		"Empty key": {
			key:     "",
			cmp:     query.Equal,
			value:   "2",
			wantErr: query.ErrFilterKeyInvalid,
		},
		"Empty comparator": {
			key:     "key_a",
			cmp:     query.Comparator(""),
			value:   "2",
			wantErr: query.ErrFilterCmpInvalid,
		},
		"Empty value": {
			key:     "key_a",
			cmp:     query.Equal,
			value:   "",
			wantErr: query.ErrFilterValueInvalid,
		},
		"Multiple comparator": {
			key:     "key_a",
			cmp:     query.Comparator(">=!="),
			value:   "2",
			wantErr: query.ErrFilterCmpInvalid,
		},
		"Multiple keys": {
			key:     "key_a key_b",
			cmp:     query.Equal,
			value:   "2",
			wantErr: query.ErrFilterKeyInvalid,
		},
		"Multiple values (space separated)": {
			key:     "key_a",
			cmp:     query.Equal,
			value:   "2 4",
			wantErr: query.ErrFilterValueInvalid,
		},
		"Multiple values (comma separated)": {
			key:     "key_a",
			cmp:     query.Equal,
			value:   "2,4",
			wantErr: query.ErrFilterValueInvalid,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := query.NewFilter(tt.key, tt.cmp, tt.value)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got %v, want %v", err, tt.wantErr)
			}
		})
	}
}
