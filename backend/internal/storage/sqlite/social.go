package sqlite

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)

type Social struct {
	ID  int64
	URL string
}

// Clears any previous artist socials and sets them to the input socials.
func setArtistSocials(ctx context.Context, tx *sql.Tx, artistID int64, socials ...Social) error {
	err := deleteArtistSocials(ctx, tx, artistID)
	if err != nil {
		return err
	}

	for _, social := range socials {
		socialID, err := insertSocial(ctx, tx, social.URL)
		if err != nil {
			return err
		}

		err = associateArtistWithSocial(ctx, tx, artistID, socialID)
		if err != nil {
			return err
		}
	}

	return nil
}

// Deletes all artist's socials.
func deleteArtistSocials(ctx context.Context, tx *sql.Tx, artistID int64) error {
	socials, err := artistSocials(ctx, tx, artistID)
	if err != nil {
		return err
	}

	query, args, err := sq.
		Delete("artists_socials").
		Where(sq.Eq{"artist_id": artistID}).
		ToSql()

	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	for _, social := range socials {
		err = deleteSocial(ctx, tx, social.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

// Associates a social entry with a given artist.
func associateArtistWithSocial(ctx context.Context, tx *sql.Tx, artistID int64, socialID int64) error {
	query, args, err := sq.
		Insert("artists_socials").
		Columns("artist_id", "social_id").
		Values(artistID, socialID).
		ToSql()

	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

// Inserts a social entry into the database.
func insertSocial(ctx context.Context, tx *sql.Tx, url string) (int64, error) {
	query, args, err := sq.
		Insert("social").
		Columns("url").
		Values(url).
		ToSql()

	if err != nil {
		return 0, err
	}

	res, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	socialID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return socialID, nil
}

var socialBuilder = sq.
	Select(
		"social.id",
		"social.url",
	).
	From("social")

func scanSocial(s Scanner, dst *Social) error {
	if err := s.Scan(&dst.ID, &dst.URL); err != nil {
		return err
	}

	return nil
}

// Lists all social entries associated with an artist.
func artistSocials(ctx context.Context, tx *sql.Tx, artistID int64) ([]Social, error) {
	query, args, err := socialBuilder.
		Join("artists_socials ON artists_socials.social_id = social.id").
		Where(sq.Eq{"artists_socials.artist_id": artistID}).
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	socials := make([]Social, 0)

	for rows.Next() {
		var s Social
		if err := scanSocial(rows, &s); err != nil {
			return nil, err
		}

		socials = append(socials, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return socials, nil
}

// Deletes a social given its ID.
func deleteSocial(ctx context.Context, tx *sql.Tx, socialID int64) error {
	query, args, err := sq.
		Delete("social").
		Where(sq.Eq{"id": socialID}).
		ToSql()

	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
