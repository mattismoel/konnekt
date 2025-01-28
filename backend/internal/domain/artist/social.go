package artist

import (
	"errors"
	"net/http"
)

type Social string

var (
	ErrInvalidSocialURL = errors.New("Social URL must be valid and accessible")
)

func NewSocial(url string) (Social, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", ErrInvalidSocialURL
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return "", ErrInvalidSocialURL
	}

	return Social(url), nil
}
