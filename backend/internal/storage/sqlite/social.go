package sqlite

import (
	"context"
	"database/sql"
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

	query, err := NewQuery("DELETE FROM artists_socials")
	if err != nil {
		return err
	}

	err = query.AddFilter("artist_id = ?", artistID)
	if err != nil {
		return err
	}

	queryStr, args := query.Build()

	_, err = tx.ExecContext(ctx, queryStr, args...)
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

func associateArtistWithSocial(ctx context.Context, tx *sql.Tx, artistID int64, socialID int64) error {
	query := `
	INSERT INTO artists_socials (artist_id, social_id)
	VALUES (@artist_id, @social_id)`

	_, err := tx.ExecContext(ctx, query,
		sql.Named("artist_id", artistID),
		sql.Named("social_id", socialID),
	)

	if err != nil {
		return err
	}

	return nil
}

func insertSocial(ctx context.Context, tx *sql.Tx, url string) (int64, error) {
	query := `INSERT INTO social (url) VALUES (@url)`

	res, err := tx.ExecContext(ctx, query, sql.Named("url", url))
	if err != nil {
		return 0, err
	}

	socialID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return socialID, nil
}

func artistSocials(ctx context.Context, tx *sql.Tx, artistID int64) ([]Social, error) {
	query := `
	SELECT id, url FROM social
	JOIN artists_socials ON artists_socials.social_id = social.id
	WHERE artists_socials.artist_id = @artist_id`

	rows, err := tx.QueryContext(ctx, query, sql.Named("artist_id", artistID))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	socials := make([]Social, 0)

	for rows.Next() {
		var id int64
		var url string

		if err := rows.Scan(&id, &url); err != nil {
			return nil, err
		}

		socials = append(socials, Social{ID: id, URL: url})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return socials, nil
}

func deleteSocial(ctx context.Context, tx *sql.Tx, socialID int64) error {
	query := `DELETE FROM social WHERE id = @id`

	_, err := tx.ExecContext(ctx, query, sql.Named("id", socialID))
	if err != nil {
		return err
	}

	return nil
}
