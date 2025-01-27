package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/mattismoel/konnekt/internal/domain/artist"
	"github.com/mattismoel/konnekt/internal/domain/concert"
	"github.com/mattismoel/konnekt/internal/domain/event"
	"github.com/mattismoel/konnekt/internal/domain/venue"
)

type Event struct {
	ID            int64
	Title         string
	Description   string
	CoverImageURL string
	VenueID       int64
}

type Concert struct {
	ID       int64
	From     time.Time
	To       time.Time
	ArtistID int64
}

var _ event.Repository = (*EventRepository)(nil)

type EventRepository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) (*EventRepository, error) {
	return &EventRepository{
		db: db,
	}, nil
}

func (repo EventRepository) Insert(ctx context.Context, e event.Event) (int64, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	defer tx.Rollback()

	eventID, err := insertEvent(ctx, tx, Event{
		Title:         e.Title,
		Description:   e.Description,
		CoverImageURL: e.CoverImageURL,
		VenueID:       e.Venue.ID,
	})

	if err != nil {
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return eventID, nil
}

func (repo EventRepository) List(ctx context.Context) ([]event.Event, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	dbEvents, err := listEvents(ctx, tx)
	if err != nil {
		return nil, err
	}

	events := make([]event.Event, 0)

	for _, dbEvent := range dbEvents {
		dbVenue, err := venueByID(ctx, tx, dbEvent.VenueID)
		if err != nil {
			return nil, err
		}

		venue := dbVenue.ToInternal()

		dbConcerts, err := eventConcerts(ctx, tx, dbEvent.ID)
		if err != nil {
			return nil, err
		}

		concerts := make([]concert.Concert, 0)
		for _, dbConcert := range dbConcerts {
			dbArtist, err := artistByID(ctx, tx, dbConcert.ArtistID)
			if err != nil {
				return nil, err
			}

			dbGenres, err := artistGenres(ctx, tx, dbArtist.ID)
			if err != nil {
				return nil, err
			}

			genres := make([]artist.Genre, 0)
			for _, g := range dbGenres {
				genres = append(genres, artist.Genre(g.Name))
			}

			dbSocials, err := artistSocials(ctx, tx, dbArtist.ID)
			if err != nil {
				return nil, err
			}

			socials := make([]artist.Social, 0)
			for _, s := range dbSocials {
				socials = append(socials, artist.Social(s.URL))
			}

			artist := dbArtist.ToInternal(genres, socials)

			concert := dbConcert.ToInternal(artist)
			concerts = append(concerts, concert)
		}

		event := dbEvent.ToInternal(venue, concerts)
		events = append(events, event)
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return events, nil
}

func listEvents(ctx context.Context, tx *sql.Tx) ([]Event, error) {
	query := `
	SELECT id, title, description, cover_image_url, venue_id FROM event`

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	events := make([]Event, 0)
	for rows.Next() {
		var id, venueID int64
		var title, description, coverImageURL string

		err := rows.Scan(&id, &title, &description, &coverImageURL, &venueID)
		if err != nil {
			return nil, err
		}

		events = append(events, Event{
			ID:            id,
			Title:         title,
			Description:   description,
			CoverImageURL: coverImageURL,
			VenueID:       venueID,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

func eventConcerts(ctx context.Context, tx *sql.Tx, eventID int64) ([]Concert, error) {
	query := `
	SELECT id, from_date, to_date, artist_id FROM concert
	WHERE event_id = @event_id`

	rows, err := tx.QueryContext(ctx, query, sql.Named("event_id", eventID))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	concerts := make([]Concert, 0)
	for rows.Next() {
		var id int64
		var fromDate, toDate time.Time
		var artistID int64

		if err := rows.Scan(&id, &fromDate, &toDate, &artistID); err != nil {
			return nil, err
		}

		concerts = append(concerts, Concert{
			ID:       id,
			From:     fromDate,
			To:       toDate,
			ArtistID: artistID,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return concerts, nil
}

func insertEvent(ctx context.Context, tx *sql.Tx, e Event) (int64, error) {
	query := `
	INSERT INTO event (title, description, cover_image_url, venue_id) 
	VALUES (@title, @description, @cover_image_url, @venue_id)`

	res, err := tx.ExecContext(ctx, query,
		sql.Named("title", e.Title),
		sql.Named("description", e.Description),
		sql.Named("cover_image_url", e.CoverImageURL),
		sql.Named("venue_id", e.VenueID),
	)

	if err != nil {
		return 0, err
	}

	eventID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return eventID, nil
}

func (e Event) ToInternal(venue venue.Venue, concerts []concert.Concert) event.Event {
	return event.Event{
		ID:            e.ID,
		Title:         e.Title,
		Description:   e.Description,
		CoverImageURL: e.CoverImageURL,
		Venue:         venue,
		Concerts:      concerts,
	}
}

func (c Concert) ToInternal(a artist.Artist) concert.Concert {
	return concert.Concert{
		ID:     c.ID,
		Artist: a,
		From:   c.From,
		To:     c.To,
	}
}
