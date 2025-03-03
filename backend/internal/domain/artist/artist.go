package artist

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
)

var (
	ErrInvalidID            = errors.New("ID must be a positive integer")
	ErrEmptyName            = errors.New("Name must not be empty")
	ErrEmptyDescription     = errors.New("Description must not be empty")
	ErrInvalidImageURL      = errors.New("Image URL must be valid")
	ErrImageURLInaccessible = errors.New("Image URL must be accessible")
	ErrNoGenres             = errors.New("Artist must have at least one genre")
)

type ArtistCfg func(a *Artist) error

type Artist struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	ImageURL    string   `json:"imageUrl"`
	Genres      []Genre  `json:"genres"`
	Socials     []Social `json:"socials"`
}

func NewArtist(cfgs ...ArtistCfg) (*Artist, error) {
	a := &Artist{}

	err := a.WithCfgs(cfgs...)
	if err != nil {
		return &Artist{}, err
	}

	return a, nil
}

func (a *Artist) WithCfgs(cfgs ...ArtistCfg) error {
	for _, cfg := range cfgs {
		if err := cfg(a); err != nil {
			return err
		}
	}

	return nil
}

func WithID(id int64) ArtistCfg {
	return func(a *Artist) error {
		if id <= 0 {
			return ErrInvalidID
		}

		a.ID = id
		return nil
	}
}

func WithName(name string) ArtistCfg {
	return func(a *Artist) error {
		if strings.TrimSpace(name) == "" {
			return ErrEmptyName
		}

		a.Name = name

		return nil
	}
}

func WithDescription(desc string) ArtistCfg {
	return func(a *Artist) error {
		if strings.TrimSpace(desc) == "" {
			return ErrEmptyDescription
		}

		a.Description = desc

		return nil
	}
}

func WithImageURL(u string) ArtistCfg {
	return func(a *Artist) error {
		url, err := url.ParseRequestURI(u)
		if err != nil {
			return ErrInvalidImageURL
		}

		resp, err := http.Get(url.String())
		if err != nil {
			return ErrImageURLInaccessible
		}

		if !(resp.StatusCode >= 200) || !(resp.StatusCode < 400) {
			return ErrImageURLInaccessible
		}

		a.ImageURL = url.String()

		return nil
	}
}

func WithGenres(genres ...Genre) ArtistCfg {
	return func(a *Artist) error {
		a.Genres = genres
		return nil
	}
}

func WithSocials(socials ...Social) ArtistCfg {
	return func(a *Artist) error {
		a.Socials = socials
		return nil
	}
}
