package artist_test

import (
	"errors"
	"testing"

	"github.com/mattismoel/konnekt/internal/domain/artist"
)

func TestNewGenre(t *testing.T) {
	type test struct {
		name     string
		wantName string
		wantErr  error
	}

	tests := map[string]test{
		"Empty name": {
			name:     "",
			wantName: "",
			wantErr:  artist.ErrInvalidGenreName,
		},
		"Valid name": {
			name:     "Rock",
			wantName: "Rock",
			wantErr:  nil,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			g, err := artist.NewGenre(tt.name)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got %v, want %v\n", err, tt.wantErr)
			}

			if g.Name != tt.wantName {
				t.Fatalf("got %q, want %q\n", g.Name, tt.wantName)
			}
		})
	}
}
