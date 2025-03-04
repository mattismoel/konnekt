package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/mattismoel/konnekt/internal/domain/concert"
	"github.com/mattismoel/konnekt/internal/domain/event"
	"github.com/mattismoel/konnekt/internal/domain/venue"
)

type EventRepository struct {
	db *sql.DB
}

var _ event.Repository = (*EventRepository)(nil)

type Event struct {
	ID            int64
	Title         string
	Description   string
	CoverImageURL string
	VenueID       int64
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

func NewEventRepository(db *sql.DB) (*EventRepository, error) {
	return &EventRepository{
		db: db,
	}, nil
}

func (repo EventRepository) ByID(ctx context.Context, eventID int64) (event.Event, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return event.Event{}, err
	}

	defer tx.Rollback()

	dbEvent, err := eventByID(ctx, tx, eventID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return event.Event{}, event.ErrNoExist
		}

		return event.Event{}, err
	}

	dbConcerts, err := eventConcerts(ctx, tx, eventID)
	if err != nil {
		return event.Event{}, err
	}

	concerts, err := dbConcerts.Internalize(ctx, tx)
	if err != nil {
		return event.Event{}, err
	}

	dbVenue, err := venueByID(ctx, tx, dbEvent.VenueID)
	if err != nil {
		return event.Event{}, err
	}

	venue := dbVenue.ToInternal()

	if err := tx.Commit(); err != nil {
		return event.Event{}, err
	}

	e := dbEvent.ToInternal(venue, concerts)

	return e, nil
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

	for _, c := range e.Concerts {
		dbArtist, err := artistByID(ctx, tx, c.Artist.ID)
		if err != nil {
			return 0, err
		}

		_, err = insertConcert(ctx, tx, Concert{
			ArtistID: dbArtist.ID,
			EventID:  eventID,
			From:     c.From,
			To:       c.To,
		})

		if err != nil {
			return 0, err
		}
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return eventID, nil
}
func (repo EventRepository) Update(ctx context.Context, eventID int64, e event.Event) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	err = updateEvent(ctx, tx, eventID, Event{
		Title:         e.Title,
		Description:   e.Description,
		CoverImageURL: e.CoverImageURL,
		VenueID:       e.Venue.ID,
	})

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (repo EventRepository) SetCoverImageURL(ctx context.Context, eventID int64, coverImageURL string) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	err = setEventCoverImageURL(ctx, tx, eventID, coverImageURL)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

type EventQueryParams struct {
	QueryParams
	From time.Time
	To   time.Time
}

func (repo EventRepository) List(ctx context.Context, from, to time.Time, offset, limit int) ([]event.Event, int, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, 0, err
	}

	defer tx.Rollback()

	dbEvents, err := listEvents(ctx, tx, EventQueryParams{
		QueryParams: QueryParams{
			Offset: offset,
			Limit:  limit,
		},
		From: from,
		To:   to,
	},
	)

	if err != nil {
		return nil, 0, err
	}

	totalCount, err := eventCount(ctx, tx)
	if err != nil {
		return nil, 0, err
	}

	events := make([]event.Event, 0)

	for _, dbEvent := range dbEvents {
		dbVenue, err := venueByID(ctx, tx, dbEvent.VenueID)
		if err != nil {
			return nil, 0, err
		}

		venue := dbVenue.ToInternal()

		dbConcerts, err := eventConcerts(ctx, tx, dbEvent.ID)
		if err != nil {
			return nil, 0, err
		}

		concerts, err := dbConcerts.Internalize(ctx, tx)
		if err != nil {
			return nil, 0, err
		}

		event := dbEvent.ToInternal(venue, concerts)
		events = append(events, event)
	}

	if err := tx.Commit(); err != nil {
		return nil, 0, err
	}

	return events, totalCount, nil
}

func eventByID(ctx context.Context, tx *sql.Tx, eventID int64) (Event, error) {
	query := `
	SELECT title, description, cover_image_url, venue_id FROM event
	WHERE id = @event_id`

	var title, description, coverImageUrl string
	var venueID int64

	err := tx.QueryRowContext(ctx, query, sql.Named("event_id", eventID)).Scan(
		&title, &description, &coverImageUrl, &venueID,
	)

	if err != nil {
		return Event{}, err
	}

	return Event{
		ID:            eventID,
		Title:         title,
		Description:   description,
		CoverImageURL: coverImageUrl,
		VenueID:       venueID,
	}, nil
}

func eventCount(ctx context.Context, tx *sql.Tx) (int, error) {
	var count int

	err := tx.QueryRowContext(ctx, "select count(*) from event").Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func listEvents(ctx context.Context, tx *sql.Tx, params EventQueryParams) ([]Event, error) {
	query, err := NewQuery(`
    SELECT DISTINCT
		e.id,
		e.title,
		e.description,
		e.cover_image_url,
		e.venue_id
	FROM event e
		JOIN concert c ON c.event_id = e.id`)

	if err != nil {
		return nil, err
	}

	err = query.WithOffset(params.Offset)
	if err != nil {
		return nil, err
	}

	err = query.WithLimit(params.Limit)
	if err != nil {
		return nil, err
	}

	if !params.From.IsZero() {
		err = query.AddFilter("c.from_date >= ?", params.From)
	}

	if !params.To.IsZero() {
		query.AddFilter("c.to_date <= ?", params.To)
	}

	queryString, args := query.Build()

	rows, err := tx.QueryContext(ctx, queryString, args...)
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

func setEventCoverImageURL(ctx context.Context, tx *sql.Tx, eventID int64, url string) error {
	query := `UPDATE event SET cover_image_url = @url WHERE id = @id`

	_, err := tx.ExecContext(ctx, query,
		sql.Named("url", url),
		sql.Named("id", eventID),
	)

	if err != nil {
		return err
	}

	return nil
}

func updateEvent(ctx context.Context, tx *sql.Tx, eventID int64, e Event) error {
	query, err := NewQuery(`
		UPDATE event SET
		title = CASE
			WHEN @title = '' THEN title
			ELSE @title
		END
		description = CASE
			WHEN @description = '' THEN description
			ELSE @description
		END
		cover_image_url = CASE
			WHEN @cover_image_url = '' THEN cover_image_url
			ELSE @cover_image_url
		END`)

	err = query.AddFilter("id = ?", eventID)
	if err != nil {
		return nil
	}

	queryStr, args := query.Build()

	_, err = tx.ExecContext(ctx, queryStr, args...)
	if err != nil {
		return err
	}

	return nil
}
