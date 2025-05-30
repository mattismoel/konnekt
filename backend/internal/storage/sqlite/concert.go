package sqlite

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/mattismoel/konnekt/internal/domain/artist"
	"github.com/mattismoel/konnekt/internal/domain/concert"
)

type Concert struct {
	ID       int64
	From     time.Time
	To       time.Time
	ArtistID int64
	EventID  int64
}

type Concerts []Concert

func (cs Concerts) Internalize(ctx context.Context, tx *sql.Tx) ([]concert.Concert, error) {
	concerts := make([]concert.Concert, 0)

	for _, dbConcert := range cs {
		dbArtist, err := artistByID(ctx, tx, dbConcert.ArtistID)
		if err != nil {
			return nil, err
		}

		dbGenres, err := artistGenres(ctx, tx, dbArtist.ID)
		if err != nil {
			return nil, err
		}

		genres := make([]artist.Genre, 0)
		for _, dbGenre := range dbGenres {
			genres = append(genres, artist.Genre{
				ID:   dbGenre.ID,
				Name: dbGenre.Name,
			})
		}

		dbSocials, err := artistSocials(ctx, tx, dbArtist.ID)
		if err != nil {
			return nil, err
		}

		socials := make([]artist.Social, 0)
		for _, dbSocial := range dbSocials {
			socials = append(socials, artist.Social(dbSocial.URL))
		}

		artist := dbArtist.ToInternal(genres, socials)

		concerts = append(concerts, dbConcert.ToInternal(artist))
	}

	return concerts, nil
}

func (c Concert) ToInternal(a artist.Artist) concert.Concert {
	return concert.Concert{
		ID:     c.ID,
		Artist: a,
		From:   c.From,
		To:     c.To,
	}
}

func ConcertFromInternal(c concert.Concert, eventID int64) Concert {
	return Concert{
		ID:       c.ID,
		From:     c.From,
		To:       c.To,
		ArtistID: c.Artist.ID,
		EventID:  eventID,
	}
}

func insertConcert(ctx context.Context, tx *sql.Tx, c Concert) (int64, error) {
	query, args, err := sq.
		Insert("concert").
		Columns("event_id", "artist_id", "from_date", "to_date").
		Values(
			c.EventID,
			c.ArtistID,
			c.From.UTC().Format(time.RFC3339),
			c.To.UTC().Format(time.RFC3339),
		).
		ToSql()

	if err != nil {
		return 0, err
	}

	res, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	concertID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return concertID, nil
}

var concertBuilder = sq.
	Select(
		"concert.id",
		"concert.artist_id",
		"concert.from_date",
		"concert.to_date",
	).
	From("concert")

func scanConcert(s Scanner, dst *Concert) error {
	err := s.Scan(&dst.ID, &dst.ArtistID, &dst.From, &dst.To)
	if err != nil {
		return err
	}

	return nil
}

func eventConcerts(ctx context.Context, tx *sql.Tx, eventID int64) (Concerts, error) {
	query, args, err := concertBuilder.
		Where(sq.Eq{"event_id": eventID}).
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	concerts := make([]Concert, 0)
	for rows.Next() {
		var c Concert
		if err := scanConcert(rows, &c); err != nil {
			return nil, err
		}

		concerts = append(concerts, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return concerts, nil
}

func setEventConcerts(ctx context.Context, tx *sql.Tx, eventID int64, concerts ...Concert) (Concerts, error) {
	err := deleteEventConcerts(ctx, tx, eventID)
	if err != nil {
		return nil, err
	}

	insertedConcerts := make([]Concert, 0)
	for _, c := range concerts {
		concertID, err := insertConcert(ctx, tx, c)
		if err != nil {
			return nil, err
		}

		insertedConcerts = append(insertedConcerts, Concert{
			ID:       concertID,
			ArtistID: c.ArtistID,
			EventID:  c.EventID,
			From:     c.From,
			To:       c.To,
		})
	}

	return insertedConcerts, nil
}

func deleteEventConcerts(ctx context.Context, tx *sql.Tx, eventID int64) error {
	concerts, err := eventConcerts(ctx, tx, eventID)
	if err != nil {
		return err
	}

	for _, c := range concerts {
		err = deleteConcert(ctx, tx, c.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func deleteConcert(ctx context.Context, tx *sql.Tx, concertID int64) error {
	query, args, err := sq.
		Delete("concert").
		Where(sq.Eq{"id": concertID}).
		ToSql()

	if err != nil {
		return err
	}

	res, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected <= 0 {
		return ErrNotFound
	}

	return nil
}
