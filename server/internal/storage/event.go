package storage

import "time"

type BaseEvent struct {
	ID          int64
	Title       string
	Description string
	FromDate    time.Time
	ToDate      time.Time
	AddressID   int64
}

type Event struct {
	BaseEvent
	Address Address
	Genres  []Genre
}
