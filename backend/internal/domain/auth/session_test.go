package auth_test

import (
	"errors"
	"testing"
	"time"

	"github.com/mattismoel/konnekt/internal/domain/auth"
)

func TestWithMemberID(t *testing.T) {
	type test struct {
		id      int64
		wantID  int64
		wantErr error
	}

	tests := map[string]test{
		"Negative ID": {
			id:      -1,
			wantID:  0,
			wantErr: auth.ErrInvalidMemberID,
		},
		"Zero ID": {
			id:      0,
			wantID:  0,
			wantErr: auth.ErrInvalidMemberID,
		},
		"Valid ID": {
			id:      10,
			wantID:  10,
			wantErr: nil,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			s, err := auth.NewSession(
				auth.WithMemberID(tt.id),
			)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got %v, want %v\n", err, tt.wantErr)
			}

			if s.MemberID != tt.wantID {
				t.Fatalf("got member id %v, want %v\n", s.MemberID, tt.wantID)
			}
		})
	}
}

func TestWithLifetime(t *testing.T) {
	maxDelta := 1 * time.Millisecond

	type test struct {
		d          time.Duration
		wantExpiry time.Time
		wantErr    error
	}

	tests := map[string]test{
		"Too long duration": {
			d:          12 * 365 * 24 * time.Hour, // 12 years.
			wantExpiry: time.Now().Add(auth.MAX_LIFETIME),
			wantErr:    nil,
		},
		"Valid duration": {
			d:          1 * 24 * time.Hour,
			wantExpiry: time.Now().Add(1 * 24 * time.Hour),
			wantErr:    nil,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			s, err := auth.NewSession(
				auth.WithLifetime(tt.d),
			)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got %v, want %v\n", err, tt.wantErr)
			}

			if diff := s.ExpiresAt.Sub(tt.wantExpiry); diff < -maxDelta || diff > maxDelta {
				t.Fatalf("got expiry %v, want %v (diff %v)\n", s.ExpiresAt, tt.wantExpiry, diff)
			}
		})
	}
}
