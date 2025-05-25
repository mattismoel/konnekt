package sqlite

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/mattismoel/konnekt/internal/domain/venue"
	"github.com/mattismoel/konnekt/internal/query"
)

type Venue struct {
	ID          int64
	Name        string
	CountryCode string
	City        string
}

var _ venue.Repository = (*VenueRepository)(nil)

type VenueRepository struct {
	db *sql.DB
}

func NewVenueRepository(db *sql.DB) (*VenueRepository, error) {
	return &VenueRepository{
		db: db,
	}, nil
}

type VenueQueryParams struct {
	QueryParams
}

func (repo VenueRepository) List(ctx context.Context, q venue.Query) (query.ListResult[venue.Venue], error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return query.ListResult[venue.Venue]{}, err
	}

	defer tx.Rollback()

	dbVenues, err := listVenues(ctx, tx, VenueQueryParams{
		QueryParams: QueryParams{
			Offset: q.Offset(),
			Limit:  q.Limit,
		},
	})

	if err != nil {
		return query.ListResult[venue.Venue]{}, err
	}

	totalCount, err := count(ctx, tx, "venue")
	if err != nil {
		return query.ListResult[venue.Venue]{}, err
	}

	if err := tx.Commit(); err != nil {
		return query.ListResult[venue.Venue]{}, err
	}

	venues := make([]venue.Venue, 0)
	for _, dbVenue := range dbVenues {
		venues = append(venues, dbVenue.ToInternal())
	}

	return query.ListResult[venue.Venue]{
		Page:       q.Page,
		PerPage:    q.PerPage,
		TotalCount: totalCount,
		PageCount:  q.PageCount(totalCount),
		Records:    venues,
	}, nil
}

func (repo VenueRepository) ByID(ctx context.Context, venueID int64) (venue.Venue, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return venue.Venue{}, err
	}

	defer tx.Rollback()

	dbVenue, err := venueByID(ctx, tx, venueID)
	if err != nil {
		return venue.Venue{}, err
	}

	if err := tx.Commit(); err != nil {
		return venue.Venue{}, err
	}

	return dbVenue.ToInternal(), nil
}

func (repo VenueRepository) Insert(ctx context.Context, v venue.Venue) (int64, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	defer tx.Rollback()

	venueID, err := insertVenue(ctx, tx, Venue{
		Name:        v.Name,
		City:        v.City,
		CountryCode: v.CountryCode,
	})

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return venueID, nil
}

func (repo VenueRepository) Update(ctx context.Context, venueID int64, v venue.Venue) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	if err := updateVenue(ctx, tx, venueID, Venue{
		Name:        v.Name,
		City:        v.City,
		CountryCode: v.CountryCode,
	}); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (repo VenueRepository) Delete(ctx context.Context, venueID int64) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	err = deleteVenue(ctx, tx, venueID)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

var venueBuilder = sq.Select("id", "name", "country_code", "city").From("venue")

func scanVenue(s Scanner, dst *Venue) error {
	err := s.Scan(&dst.ID, &dst.Name, &dst.CountryCode, &dst.City)
	if err != nil {
		return err
	}

	return nil
}

func listVenues(ctx context.Context, tx *sql.Tx, params VenueQueryParams) ([]Venue, error) {
	builder := venueBuilder

	if params.Limit > 0 {
		builder = builder.Limit(uint64(params.Limit))
	}

	if params.Offset > 0 {
		builder = builder.Offset(uint64(params.Offset))
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	venues := make([]Venue, 0)

	for rows.Next() {
		var v Venue
		if err := scanVenue(rows, &v); err != nil {
			return nil, err
		}

		venues = append(venues, v)
	}

	return venues, nil
}

func insertVenue(ctx context.Context, tx *sql.Tx, v Venue) (int64, error) {
	query, args, err := sq.
		Insert("venue").
		Columns("name", "country_code", "city").
		Values(v.Name, v.CountryCode, v.City).
		ToSql()

	if err != nil {
		return 0, err
	}

	res, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	venueID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return venueID, nil
}

func venueByID(ctx context.Context, tx *sql.Tx, venueID int64) (Venue, error) {
	query, args, err := venueBuilder.
		Where(sq.Eq{"id": venueID}).
		ToSql()

	if err != nil {
		return Venue{}, err
	}

	var v Venue
	row := tx.QueryRowContext(ctx, query, args...)
	if err := scanVenue(row, &v); err != nil {
		return Venue{}, err
	}

	return v, nil
}

func deleteVenue(ctx context.Context, tx *sql.Tx, venueID int64) error {
	query, args, err := sq.
		Delete("venue").
		Where(sq.Eq{"id": venueID}).
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

func updateVenue(ctx context.Context, tx *sql.Tx, venueID int64, v Venue) error {
	builder := sq.Update("venue").Where(sq.Eq{"id": venueID})

	if v.Name != "" {
		builder = builder.Set("name", v.Name)
	}
	if v.City != "" {
		builder = builder.Set("city", v.City)
	}
	if v.CountryCode != "" {
		builder = builder.Set("country_code", v.Name)
	}

	query, args, err := builder.ToSql()
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

func (v Venue) ToInternal() venue.Venue {
	return venue.Venue{
		ID:          v.ID,
		Name:        v.Name,
		CountryCode: v.CountryCode,
		City:        v.City,
	}
}
