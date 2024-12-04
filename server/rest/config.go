package rest

import (
	"fmt"
	"strings"
)

type Cfg struct {
	EventService EventService
	UserService  UserService
	GenreService GenreService
	Port         int
	Host         string
}

func (cfg Cfg) Validate() error {
	if cfg.EventService == nil {
		return fmt.Errorf("Event service must be set")
	}

	if cfg.UserService == nil {
		return fmt.Errorf("User service must be set")
	}

	if cfg.GenreService == nil {
		return fmt.Errorf("Genre service must be set")
	}

	if strings.TrimSpace(cfg.Host) == "" {
		return fmt.Errorf("Host must be set")
	}

	if cfg.Port < 0 {
		return fmt.Errorf("Port must be a positive integer within port range")
	}

	return nil
}
