package artist_test

import (
	"errors"
	"testing"

	"github.com/mattismoel/konnekt/internal/domain/artist"
)

func TestNewSocial(t *testing.T) {
	type test struct {
		url     string
		wantErr error
	}

	tests := map[string]test{
		"Valid URL": {
			url:     "https://google.com",
			wantErr: nil,
		},
		"Invalid URL": {
			url:     "http/google.com",
			wantErr: artist.ErrInvalidSocialURL,
		},
		"Inaccessible URL": {
			url:     "http://google.com/hello",
			wantErr: artist.ErrSocialURLInaccessible,
		},
		"Empty URL": {
			url:     "",
			wantErr: artist.ErrInvalidSocialURL,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := artist.NewSocial(tt.url)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got %v, want %v\n", err, tt.wantErr)
			}
		})
	}
}
