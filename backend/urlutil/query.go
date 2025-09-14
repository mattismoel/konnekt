package urlutil

import (
	"errors"
	"net/http"
	"strconv"
)

var (
	ErrNoSuchQueryKey = errors.New("No such query key in URL")
)

// Returns the URL query parameter 'key' as an parsed int, if possible.
func QueryInt(r *http.Request, key string) (int, error) {
	s := r.URL.Query().Get(key)
	if s == "" {
		return 0, ErrNoSuchQueryKey
	}

	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return v, nil
}
