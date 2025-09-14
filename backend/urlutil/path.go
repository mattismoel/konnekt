package urlutil

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

var (
	ErrNoSuchPathValue = errors.New("No such key in URL path")
)

func PathInt(r *http.Request, key string) (int, error) {
	s := r.PathValue(key)
	if s == "" {
		return 0, ErrNoSuchPathValue
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("Could not parse path value %s as integer: %v", key, err)
	}

	return i, nil
}
