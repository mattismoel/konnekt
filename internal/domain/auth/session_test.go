package auth_test

import (
	"testing"
	"time"

	"github.com/mattismoel/konnekt/internal/domain/auth"
)

func TestIsRefreshable(t *testing.T) {
	const SESSION_LIFETIME = 30 * 24 * time.Hour        // 30 days.
	const REFRESH_BUFFER_DURATION = 15 * 24 * time.Hour // 15 days.

	now := time.Now()

	type test struct {
		expiry            time.Time
		expectRefreshable bool
	}

	tests := map[string]test{
		"Non-refreshable": {
			expiry:            now.Add(16 * 24 * time.Hour), // Expiry in 16 days, out of refresh bounds.
			expectRefreshable: false,
		},
		"Refreshable (Within buffer)": {
			expiry:            now.Add(14 * 24 * time.Hour), // Expiry in 14 days.
			expectRefreshable: true,
		},
		"Refreshable (Expires now)": {
			expiry:            now, // Expiry now.
			expectRefreshable: true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			userID := 1

			token, err := auth.NewSessionToken()
			if err != nil {
				t.Fatal(err)
			}

			s := auth.NewSession(token, int64(userID), SESSION_LIFETIME)

			refreshable := s.IsRefreshable(REFRESH_BUFFER_DURATION)
			if tt.expectRefreshable != refreshable {
				t.Fatalf("expected refreshable %v, got %v", tt.expectRefreshable, refreshable)
			}
		})
	}
}
