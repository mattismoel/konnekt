package psql

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	konnekt "github.com/mattismoel/konnekt/backend"
)

type Social struct {
	ID       int64  `db:"id"`
	URL      string `db:"url"`
	ArtistID int64  `db:"artist_id"`
}

func artistSocials(ctx context.Context, tx pgx.Tx, artistID int64) (Collection[Social, konnekt.Social], error) {
	query, args, err := psql.
		Select("*").
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

	socials, err := pgx.CollectRows(rows, pgx.RowToStructByName[Social])
	if err != nil {
		return nil, fmt.Errorf("Could not get artist socials: %v", err)
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

func setArtistSocials(ctx context.Context, tx pgx.Tx, artistID int64, socials ...string) error {
	if err := deleteArtistSocials(ctx, tx, artistID); err != nil {
		return fmt.Errorf("Could not insert artist socials: %v", err)
	}

	builder := psql.Insert("social").Columns("url", "artist_id")

	for _, socialURL := range socials {
		builder = builder.Values(socialURL, artistID)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return NewQueryBuildError("insert artist social", err)
	}

	if _, err := tx.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("Could not insert artist social: %v", err)
	}

	return nil
}

func (s Social) Assemble(context.Context, pgx.Tx) (konnekt.Social, error) {
	return konnekt.Social(s.URL), nil
}
