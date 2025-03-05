package sqlite

import (
	"context"
	"database/sql"
	"time"

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

func insertConcert(ctx context.Context, tx *sql.Tx, c Concert) (int64, error) {
	query := `
	INSERT into concert (event_id, artist_id, from_date, to_date)
	VALUES (@event_id, @artist_id, @from_date, @to_date)`

	res, err := tx.ExecContext(ctx, query,
		sql.Named("event_id", c.EventID),
		sql.Named("artist_id", c.ArtistID),
		sql.Named("from_date", c.From),
		sql.Named("to_date", c.To),
	)

	if err != nil {
		return 0, err
	}

	concertID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return concertID, nil
}

func eventConcerts(ctx context.Context, tx *sql.Tx, eventID int64) (Concerts, error) {
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
	query := "DELETE FROM concert WHERE id = @id"

	res, err := tx.ExecContext(ctx, query, sql.Named("id", concertID))
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

func updateConcert(ctx context.Context, tx *sql.Tx, concertID int64, c Concert) error {
	query := `UPDATE concert SET
	from_date = CASE
		WHEN @from_date = '' THEN from_date ELSE @from_date
	END,
	to_date = CASE
		WHEN @to_date = '' THEN to_date ELSE @to_date
	END,
	artist_id = CASE
		WHEN @artist_id <= 0 THEN artist_id ELSE @artist_id
	END
	WHERE id = @id`

	res, err := tx.ExecContext(ctx, query,
		sql.Named("from_date", c.From),
		sql.Named("to_date", c.To),
		sql.Named("artist_id", c.ArtistID),
		sql.Named("id", concertID),
	)

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
