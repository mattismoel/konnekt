package artist_test

import (
	"errors"
	"testing"

	"github.com/mattismoel/konnekt/internal/domain/artist"
)

func TestNewSocial(t *testing.T) {
	type test struct {
		url        string
		wantSocial artist.Social
		wantErr    error
	}

	tests := map[string]test{
		"Empty URL": {
			url:        "",
			wantSocial: "",
			wantErr:    artist.ErrInvalidSocialURL,
		},
		"Valid URL": {
			url:        "https://example.com",
			wantSocial: "https://example.com",
			wantErr:    nil,
		},
		"Invalid URL": {
			url:        "http/google.com",
			wantSocial: "",
			wantErr:    artist.ErrInvalidSocialURL,
		},
		"Inaccessible URL": {
			url:        "http://google.com/hello",
			wantSocial: "",
			wantErr:    artist.ErrInaccessibleSocialURL,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			s, err := artist.NewSocial(tt.url)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got %v, want %v\n", err, tt.wantErr)
			}

			if s != tt.wantSocial {
				t.Fatalf("got %q, want %q\n", s, tt.wantSocial)
			}
		})
	}
}
