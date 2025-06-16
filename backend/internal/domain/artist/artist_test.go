package artist_test

import (
	"errors"
	"slices"
	"testing"

	"github.com/mattismoel/konnekt/internal/domain/artist"
)

func TestWithID(t *testing.T) {
	type test struct {
		id     int64
		wantID int64
		err    error
	}

	tests := map[string]test{
		"Negative ID": {
			id:     -1,
			err:    artist.ErrInvalidID,
			wantID: 0,
		},
		"Zero-ID": {
			id:     0,
			err:    artist.ErrInvalidID,
			wantID: 0,
		},
		"Valid ID": {
			id:     10,
			err:    nil,
			wantID: 10,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			a, err := artist.NewArtist(
				artist.WithID(tt.id),
			)

			if !errors.Is(err, tt.err) {
				t.Fatalf("got %v, want %v\n", err, tt.err)
			}

			if a.ID != tt.wantID {
				t.Fatalf("got %d, want %d\n", a.ID, tt.wantID)
			}
		})
	}
}

func TestWithName(t *testing.T) {
	type test struct {
		name     string
		wantName string
		err      error
	}

	tests := map[string]test{
		"Emtpy Name": {
			name:     "",
			err:      artist.ErrInvalidName,
			wantName: "",
		},
		"Space-only name": {
			name:     " ",
			err:      artist.ErrInvalidName,
			wantName: "",
		},
		"Valid Name": {
			name:     "Christopher",
			err:      nil,
			wantName: "Christopher",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			a, err := artist.NewArtist(
				artist.WithName(tt.name),
			)

			if !errors.Is(err, tt.err) {
				t.Fatalf("got %v, want %v\n", err, tt.err)
			}

			if a.Name != tt.wantName {
				t.Fatalf("got %q, want %q\n", a.Name, tt.wantName)
			}
		})
	}
}

func TestWithDescription(t *testing.T) {
	type test struct {
		description     string
		wantDescription string
		err             error
	}

	tests := map[string]test{
		"Emtpy Description": {
			description:     "",
			err:             artist.ErrInvalidDescription,
			wantDescription: "",
		},
		"Space-only description": {
			description:     " ",
			err:             artist.ErrInvalidDescription,
			wantDescription: "",
		},
		"Valid Description": {
			description:     "Test description...",
			err:             nil,
			wantDescription: "Test description...",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			a, err := artist.NewArtist(
				artist.WithDescription(tt.description),
			)

			if !errors.Is(err, tt.err) {
				t.Fatalf("got %v, want %v\n", err, tt.err)
			}

			if a.Description != tt.wantDescription {
				t.Fatalf("got %q, want %q\n", a.Description, tt.wantDescription)
			}
		})
	}
}

func TestWithImageURL(t *testing.T) {
	type test struct {
		url     string
		wantURL string
		err     error
	}

	tests := map[string]test{
		"Emtpy URL": {
			url:     "",
			err:     artist.ErrInvalidImageURL,
			wantURL: "",
		},
		"Space-only URL": {
			url:     " ",
			err:     artist.ErrInvalidImageURL,
			wantURL: "",
		},
		"Valid URL": {
			url:     "https://konnekt-bucket.s3.eu-north-1.amazonaws.com/artists/128.png",
			err:     nil,
			wantURL: "https://konnekt-bucket.s3.eu-north-1.amazonaws.com/artists/128.png",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			a, err := artist.NewArtist(
				artist.WithImageURL(tt.url),
			)

			if !errors.Is(err, tt.err) {
				t.Fatalf("got %v, want %v\n", err, tt.err)
			}

			if a.ImageURL != tt.wantURL {
				t.Fatalf("got %q, want %q\n", a.ImageURL, tt.wantURL)
			}
		})
	}
}

func TestWithPreviewURL(t *testing.T) {
	type test struct {
		url     string
		wantURL string
		err     error
	}

	tests := map[string]test{
		"Emtpy URL": {
			url:     "",
			err:     artist.ErrInvalidPreviewURL,
			wantURL: "",
		},
		"Space-only URL": {
			url:     " ",
			err:     artist.ErrInvalidPreviewURL,
			wantURL: "",
		},
		"Valid URL": {
			url:     "https://example.com",
			err:     nil,
			wantURL: "https://example.com",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			a, err := artist.NewArtist(
				artist.WithPreviewURL(tt.url),
			)

			if !errors.Is(err, tt.err) {
				t.Fatalf("got %v, want %v\n", err, tt.err)
			}

			if a.PreviewURL != tt.wantURL {
				t.Fatalf("got %q, want %q\n", a.PreviewURL, tt.wantURL)
			}
		})
	}
}

func TestWithGenres(t *testing.T) {
	type test struct {
		genres     []artist.Genre
		wantGenres []artist.Genre
		err        error
	}

	tests := map[string]test{
		"No genres": {
			genres:     make([]artist.Genre, 0),
			wantGenres: make([]artist.Genre, 0),
			err:        artist.ErrNoGenres,
		},
		"Valid genres": {
			genres: []artist.Genre{
				{ID: 1, Name: "Rock"},
				{ID: 2, Name: "Indie"},
				{ID: 3, Name: "Pop"},
			},
			wantGenres: []artist.Genre{
				{ID: 1, Name: "Rock"},
				{ID: 2, Name: "Indie"},
				{ID: 3, Name: "Pop"},
			},
			err: nil,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			a, err := artist.NewArtist(
				artist.WithGenres(tt.genres...),
			)

			if !errors.Is(err, tt.err) {
				t.Fatalf("got %v, want %v\n", err, tt.err)
			}

			if !slices.Equal(a.Genres, tt.genres) {
				t.Fatalf("got %+v, want %+v\n", a.Genres, tt.genres)
			}
		})
	}
}

func TestWithSocials(t *testing.T) {
	type test struct {
		socials     []artist.Social
		wantSocials []artist.Social
		err         error
	}

	tests := map[string]test{
		"No socials": {
			socials:     make([]artist.Social, 0),
			wantSocials: make([]artist.Social, 0),
			err:         nil,
		},
		"Valid socials": {
			socials: []artist.Social{
				"https://example.com/1",
				"https://example.com/2",
			},
			wantSocials: []artist.Social{
				"https://example.com/1",
				"https://example.com/2",
			},
			err: nil,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			a, err := artist.NewArtist(
				artist.WithSocials(tt.socials...),
			)

			if !errors.Is(err, tt.err) {
				t.Fatalf("got %v, want %v\n", err, tt.err)
			}

			if !slices.Equal(a.Socials, tt.socials) {
				t.Fatalf("got %+v, want %+v\n", a.Socials, tt.socials)
			}
		})
	}
}
