package rest

import (
	"fmt"
	"strings"
)

type Cfg struct {
	EventService EventService
	Port         int
	Host         string
}

func (cfg Cfg) Validate() error {
	if cfg.EventService == nil {
		return fmt.Errorf("Event service must be set")
	}

	if strings.TrimSpace(cfg.Host) == "" {
		return fmt.Errorf("Host must be set")
	}

	if cfg.Port < 0 {
		return fmt.Errorf("Port must be a positive integer within port range")
	}

	return nil
}
