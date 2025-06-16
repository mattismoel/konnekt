package artist

import (
	"errors"
	"net/http"
	"net/url"
)

type Social string

var (
	ErrInvalidSocialURL      = errors.New("Social URL must be valid")
	ErrInaccessibleSocialURL = errors.New("Social URL must be accessible")
)

func NewSocial(urlString string) (Social, error) {
	url, err := url.ParseRequestURI(urlString)
	if err != nil {
		return "", ErrInvalidSocialURL
	}

	resp, err := http.Get(url.String())
	if err != nil {
		return "", ErrInaccessibleSocialURL
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return "", ErrInaccessibleSocialURL
	}

	return Social(url.String()), nil
}
