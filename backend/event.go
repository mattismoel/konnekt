package konnekt

import (
	"time"

	"github.com/mattismoel/konnekt/backend/mask"
)

type Event struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	TicketURL   string    `json:"ticketUrl"`
	ImageURL    string    `json:"imageUrl"`
	Concerts    []Concert `json:"concerts"`
	Venue       Venue     `json:"venue"`
}

type Concert struct {
	ID     int64     `json:"id"`
	Artist Artist    `json:"artist"`
	From   time.Time `json:"from"`
	To     time.Time `json:"to"`
}

type CreateEvent struct {
	Title       string          `json:"title"`
	Description string          `json:"description"`
	TicketURL   string          `json:"ticketUrl"`
	ImageURL    string          `json:"imageUrl"`
	IsPublic    bool            `json:"isPublic"`
	VenueID     int64           `json:"venueId"`
	Concerts    []CreateConcert `json:"concerts"`
}

type CreateConcert struct {
	ArtistID int64     `json:"artistId"`
	From     time.Time `json:"from"`
	To       time.Time `json:"to"`
}

type UpdateEvent struct {
	Title       string          `json:"title"`
	Description string          `json:"desccription"`
	IsPublic    bool            `json:"isPublic"`
	VenueID     int64           `json:"venueId"`
	Concerts    []CreateConcert `json:"concerts"`
}

func (ue UpdateEvent) Fields() mask.FieldMap {
	return mask.FieldMap{
		"title":       ue.Title,
		"description": ue.Description,
		"is_public":   ue.IsPublic,
		"venue_id":    ue.VenueID,
		"concerts":    ue.Concerts,
	}
}
