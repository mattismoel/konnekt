package psql

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	konnekt "github.com/mattismoel/konnekt/backend"
)

type Concert struct {
	ID       int64     `db:"id"`
	FromDate time.Time `db:"from_date"`
	ToDate   time.Time `db:"to_date"`
	EventID  int64     `db:"event_id"`
	ArtistID int64     `db:"artist_id"`
}

func eventConcerts(ctx context.Context, tx pgx.Tx, eventID int64) (Collection[Concert, konnekt.Concert], error) {
	query, args, err := psql.
		Select("*").
		From("concert").
		Where(sq.Eq{"event_id": eventID}).
		ToSql()

	if err != nil {
		return nil, NewQueryBuildError("event concerts", err)
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("Could not scan rows: %v", err)
	}

	concerts, err := pgx.CollectRows(rows, pgx.RowToStructByName[Concert])
	if err != nil {
		return nil, fmt.Errorf("Could not collect concerts: %v", err)
	}

	return concerts, nil
}

func deleteEventConcerts(ctx context.Context, tx pgx.Tx, eventID int64) error {
	query, args, err := psql.
		Delete("concert").
		Where(sq.Eq{"event_id": eventID}).
		ToSql()

	if err != nil {
		return NewQueryBuildError("delete event concerts", err)
	}

	if _, err := tx.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("Could not delete event concerts: %v", err)
	}

	return nil
}

func setEventConcerts(ctx context.Context, tx pgx.Tx, eventID int64, concerts ...Concert) error {
	if err := deleteEventConcerts(ctx, tx, eventID); err != nil {
		return err
	}

	builder := psql.
		Insert("concert").
		Columns("from_date", "to_date", "event_id", "artist_id")

	for _, c := range concerts {
		builder = builder.Values(c.FromDate, c.ToDate, c.EventID, c.ArtistID)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return NewQueryBuildError("insert event concerts", err)
	}

	if _, err := tx.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("Could not insert event concerts: %v", err)
	}

	return nil
}

func (c Concert) Assemble(ctx context.Context, tx pgx.Tx) (konnekt.Concert, error) {
	dbArtist, err := artistByID(ctx, tx, c.ArtistID)
	if err != nil {
		return konnekt.Concert{}, fmt.Errorf("Could not get concert artist: %v", err)
	}

	artist, err := dbArtist.Assemble(ctx, tx)
	if err != nil {
		return konnekt.Concert{}, fmt.Errorf("Could not assemble concert artist: %v", err)
	}

	return konnekt.Concert{
		ID:     c.ID,
		From:   c.FromDate,
		To:     c.ToDate,
		Artist: artist,
	}, nil
}
