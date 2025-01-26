package sqlite

import (
	"context"
	"database/sql"

	"github.com/mattismoel/konnekt/internal/domain/venue"
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

func insertVenue(ctx context.Context, tx *sql.Tx, v Venue) (int64, error) {
	query := `
	INSERT INTO venue (name, country_code, city) 
	VALUES (@name, @country_code, @city)`

	res, err := tx.ExecContext(ctx, query,
		sql.Named("name", v.Name),
		sql.Named("country_code", v.CountryCode),
		sql.Named("city", v.City),
	)

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
	query := `SELECT name, country_code, city FROM venue WHERE id = @venue_id`

	var name, countryCode, city string
	err := tx.QueryRowContext(ctx, query, sql.Named("venue_id", venueID)).Scan(
		&name, &countryCode, &city,
	)

	if err != nil {
		return Venue{}, err
	}

	return Venue{
		ID:          venueID,
		Name:        name,
		CountryCode: countryCode,
		City:        city,
	}, nil
}

func (v Venue) ToInternal() venue.Venue {
	return venue.Venue{
		ID:          v.ID,
		Name:        v.Name,
		CountryCode: v.CountryCode,
		City:        v.City,
	}
}
