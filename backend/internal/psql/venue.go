package psql

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	konnekt "github.com/mattismoel/konnekt/backend"
	"github.com/mattismoel/konnekt/backend/api"
	"github.com/mattismoel/konnekt/backend/order"
)

var _ konnekt.VenueRepo = VenueRepo{}

type Venue struct {
	ID      int64  `db:"id"`
	Name    string `db:"name"`
	City    string `db:"city"`
	Country string `db:"country"`

	Timestamps
	AuditFields
}

type VenueRepo struct {
	Pool *pgxpool.Pool
}

// ListVenues implements konnekt.VenueRepo.
func (v VenueRepo) ListVenues(ctx context.Context, lr api.ListRequest) (api.ListResponse[konnekt.Venue], error) {
	venues := make([]konnekt.Venue, 0)
	err := pgx.BeginFunc(ctx, v.Pool, func(tx pgx.Tx) error {
		pg := paginationFromListRequest(lr)

		dbVenues, err := listVenues(ctx, tx, pg, lr.OrderMap)
		if err != nil {
			return fmt.Errorf("Could not list venues: %v", err)
		}

		venues = dbVenues.ToDomain()
		return nil
	})

	if err != nil {
		return api.ListResponse[konnekt.Venue]{}, err
	}

	return api.ListResponse[konnekt.Venue]{
		Records: venues,
	}, nil
}

// VenueByID implements konnekt.VenueRepo.
func (v VenueRepo) VenueByID(ctx context.Context, venueID int64) (konnekt.Venue, error) {
	var venue konnekt.Venue
	err := pgx.BeginFunc(ctx, v.Pool, func(tx pgx.Tx) error {
		dbVenue, err := venueByID(ctx, tx, venueID)
		if err != nil {
			return fmt.Errorf("Could not get venue with ID %d: %v", venueID, err)
		}

		venue = dbVenue.ToDomain()
		return nil
	})

	if err != nil {
		return konnekt.Venue{}, err
	}

	return venue, nil
}

// InsertVenue implements konnekt.VenueRepo.
func (v VenueRepo) InsertVenue(ctx context.Context, cv konnekt.CreateVenue) (int64, error) {
	var venueID int64
	err := pgx.BeginFunc(ctx, v.Pool, func(tx pgx.Tx) error {
		insertedID, err := insertVenue(ctx, tx, Venue{
			Name:    cv.Name,
			City:    cv.City,
			Country: cv.Country,
			AuditFields: AuditFields{
				CreatedBy: cv.CreatedBy,
				UpdatedBy: cv.CreatedBy,
			},
		})

		if err != nil {
			return fmt.Errorf("Could not insert venue: %v", err)
		}

		venueID = insertedID
		return nil
	})

	if err != nil {
		return 0, err
	}

	return venueID, nil
}

func venueByID(ctx context.Context, tx pgx.Tx, venueID int64) (Venue, error) {
	query, args, err := psql.
		Select("*").
		From("venue").
		Where(sq.Eq{"id": venueID}).
		ToSql()

	if err != nil {
		return Venue{}, NewQueryBuildError("venue by ID", err)
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return Venue{}, fmt.Errorf("Could not query for venue with ID %d: %v", venueID, err)
	}

	venue, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[Venue])
	if err != nil {
		return Venue{}, fmt.Errorf("Could not collect venue: %v", err)
	}

	return venue, nil
}

func insertVenue(ctx context.Context, tx pgx.Tx, v Venue) (int64, error) {
	query, args, err := psql.
		Insert("venue").
		Columns("name", "city", "country", "created_by", "updated_by").
		Values(v.Name, v.City, v.Country, v.CreatedBy, v.UpdatedBy).
		Suffix("ON CONFLICT (name) DO UPDATE SET name = EXCLUDED.name").
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return 0, NewQueryBuildError("insert venue", err)
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {

	}

	id, err := pgx.CollectExactlyOneRow(rows, pgx.RowTo[int64])
	if err != nil {
		return 0, fmt.Errorf("Could not get inserted ID: %v", err)
	}

	return id, nil
}

func listVenues(ctx context.Context, tx pgx.Tx, pg Pagination, om order.Map) (Collection[Venue, konnekt.Venue], error) {
	builder := psql.Select("*").From("venue")
	builder = applyPagination(builder, pg)
	builder = applyOrdering(builder, om)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, NewQueryBuildError("list venues", err)
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("Could not query for venues: %v", err)
	}

	venues, err := pgx.CollectRows(rows, pgx.RowToStructByName[Venue])
	if err != nil {
		return nil, fmt.Errorf("Could not collect venues: %v", err)
	}

	return venues, nil
}

func (v Venue) ToDomain() konnekt.Venue {
	return konnekt.Venue{
		ID:      v.ID,
		Name:    v.Name,
		City:    v.City,
		Country: v.Country,
	}
}
