package sqlite

import (
	"context"
	"database/sql"

	"github.com/mattismoel/konnekt/internal/domain/event"
)

type Event struct {
	ID            int64
	Title         string
	Description   string
	CoverImageURL string
	VenueID       int64
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
