package storage

import "time"

type Session struct {
	ID        string
	UserID    int64
	ExpiresAt time.Time
}
