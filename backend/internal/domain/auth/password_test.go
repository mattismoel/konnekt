package auth_test

import (
	"errors"
	"testing"

	"github.com/mattismoel/konnekt/internal/domain/auth"
)

func TestDoPasswordsMatch(t *testing.T) {
	type test struct {
		p1  []byte
		p2  []byte
		err error
	}

	tests := map[string]test{
		"Matching passwords": {
			p1:  []byte("password123!"),
			p2:  []byte("password123!"),
			err: nil,
		},

		"Non-matching passwords": {
			p1:  []byte("wordpass123!"),
			p2:  []byte("password123!"),
			err: auth.ErrPasswordsNoMatch,
		},

		"Case sensitivity": {
			p1:  []byte("Password123!"),
			p2:  []byte("password123!"),
			err: auth.ErrPasswordsNoMatch,
		},

		"Empty p1": {
			p1:  []byte(""),
			p2:  []byte("password123!"),
			err: auth.ErrPasswordsNoMatch,
		},
		"Empty p2": {
			p1:  []byte("password123!"),
			p2:  []byte(""),
			err: auth.ErrPasswordsNoMatch,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			p := auth.Password(tt.p1)
			err := p.Matches(tt.p2)
			if !errors.Is(err, tt.err) {
				t.Fatal("want %w, got %w", tt.err, err)
			}
		})
	}
}
