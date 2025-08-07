package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/mattismoel/konnekt/internal/domain/venue"
	"github.com/mattismoel/konnekt/internal/query"
)

type Venue struct {
	ID          int64
	Name        string
	CountryCode string
	City        string

	UnixTimestamps
	AuditFields
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

func (repo VenueRepository) List(ctx context.Context, q venue.Query) (query.ListResult[venue.Venue], error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return query.ListResult[venue.Venue]{}, err
	}

	defer tx.Rollback()

	dbVenues, err := listVenues(ctx, tx, QueryParams{
		Offset: q.Offset(),
		Limit:  q.Limit,
	})

	if err != nil {
		return query.ListResult[venue.Venue]{}, fmt.Errorf("Could not list venues: %v", err)
	}

	totalCount, err := count(ctx, tx, "venue")
	if err != nil {
		return query.ListResult[venue.Venue]{}, fmt.Errorf("Could not count venues: %v", err)
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
		return venue.Venue{}, fmt.Errorf("Could not get venue with ID %d: %v", venueID, err)
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
		AuditFields: AuditFields{
			CreatedBy: v.CreatedBy,
			UpdatedBy: v.UpdatedBy,
		},
	})

	if err != nil {
		return 0, fmt.Errorf("Could not insert venue: %v", err)
	}

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

	err = updateVenue(ctx, tx, venueID, Venue{
		Name:        v.Name,
		City:        v.City,
		CountryCode: v.CountryCode,
	})

	if err != nil {
		return fmt.Errorf("Could not update venue: %v", err)
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
		return fmt.Errorf("Could not delete venue with ID %d: %v", venueID, err)
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

var venueBuilder = sq.
	Select(
		"venue.id",
		"venue.name",
		"venue.country_code",
		"venue.city",
		"venue.created_at",
		"venue.updated_at",
		"venue.created_by",
		"venue.updated_by",
	).
	From("venue")

func scanVenue(s Scanner, dst *Venue) error {
	err := s.Scan(
		&dst.ID,
		&dst.Name,
		&dst.CountryCode,
		&dst.City,
		&dst.CreatedAt,
		&dst.UpdatedAt,
		&dst.CreatedBy,
		&dst.UpdatedBy,
	)

	if err != nil {
		return fmt.Errorf("Could not scan venue into venue struct: %v", err)
	}

	return nil
}

func listVenues(ctx context.Context, tx *sql.Tx, params QueryParams) ([]Venue, error) {
	builder := venueBuilder

	builder = withPagination(builder, params)

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
		Columns("name", "country_code", "city", "created_by", "updated_by").
		Values(v.Name, v.CountryCode, v.City, v.CreatedBy, v.UpdatedBy).
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
		CreatedAt:   v.CreatedAt.Time(),
		UpdatedAt:   v.UpdatedAt.Time(),
		CreatedBy:   v.CreatedBy,
		UpdatedBy:   v.UpdatedBy,
	}
}
