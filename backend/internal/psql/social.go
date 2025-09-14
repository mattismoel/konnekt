package psql

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	konnekt "github.com/mattismoel/konnekt/backend"
)

type Social struct {
	ID       int64
	URL      string
	ArtistID int64
}

func insertSocial(ctx context.Context, tx pgx.Tx, s Social) (int64, error) {
	query, args, err := psql.Insert("social").
		Columns("url", "artist_id").
		Values(s.URL, s.ArtistID).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return 0, NewQueryBuildError("insert social", err)
	}

	var id int64
	if err := tx.QueryRow(ctx, query, args...).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func artistSocials(ctx context.Context, tx pgx.Tx, artistID int64) (Collection[Social, konnekt.Social], error) {
	query, args, err := psql.
		Select("id", "url", "artist_id").
		From("social").
		Where(sq.Eq{"artist_id": artistID}).
		ToSql()

	if err != nil {
		return nil, NewQueryBuildError("artist socials", err)
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("Could not query for artist socials: %v", err)
	}

	defer rows.Close()

	socials := make([]Social, 0)
	for rows.Next() {
		s, err := scanSocial(rows)
		if err != nil {
			return nil, err
		}

		socials = append(socials, s)
	}

	return socials, nil
}

func deleteArtistSocials(ctx context.Context, tx pgx.Tx, artistID int64) error {
	query, args, err := psql.
		Delete("social").
		Where(sq.Eq{"artist_id": artistID}).
		ToSql()

	if err != nil {
		return NewQueryBuildError("delete artist socials", err)
	}

	if _, err := tx.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("Could not delete artist socials: %v", err)
	}

	return nil
}
func scanSocial(s Scanner) (Social, error) {
	var social Social

	if err := scan(s, &social.ID, &social.URL, &social.ArtistID); err != nil {
		return Social{}, err
	}

	return social, nil

}

func (s Social) ToDomain() konnekt.Social {
	return konnekt.Social(s.URL)
}
