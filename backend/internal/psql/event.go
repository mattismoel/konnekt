package psql

import (
	"context"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	konnekt "github.com/mattismoel/konnekt/backend"
	"github.com/mattismoel/konnekt/backend/api"
	"github.com/mattismoel/konnekt/backend/internal/server"
	"github.com/mattismoel/konnekt/backend/mask"
	"github.com/mattismoel/konnekt/backend/order"
)

type Event struct {
	ID          int64  `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	TicketURL   string `db:"ticket_url"`
	ImageURL    string `db:"image_url"`
	VenueID     int64  `db:"venue_id"`
}

var _ server.EventRepo = EventRepo{}

type EventRepo struct {
	Pool *pgxpool.Pool
}

// Delete implements konnekt.EventRepo.
func (e EventRepo) Delete(context.Context, int64) error {
	panic("unimplemented")
}

// EventByID implements konnekt.EventRepo.
func (e EventRepo) EventByID(ctx context.Context, eventID int64) (konnekt.Event, error) {
	var event konnekt.Event

	err := pgx.BeginFunc(ctx, e.Pool, func(tx pgx.Tx) error {
		dbEvent, err := eventByID(ctx, tx, eventID)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return konnekt.ErrResourceNotFound
			}
		}

		event, err = dbEvent.Assemble(ctx, tx)
		if err != nil {
			return fmt.Errorf("Could not assemle DB event: %v", err)
		}

		return nil
	})

	if err != nil {
		return konnekt.Event{}, err
	}

	return event, nil
}

// InsertEvent implements konnekt.EventRepo.
func (e EventRepo) InsertEvent(ctx context.Context, ce konnekt.CreateEvent) (int64, error) {
	var eventID int64
	err := pgx.BeginFunc(ctx, e.Pool, func(tx pgx.Tx) error {
		insertedID, err := insertEvent(ctx, tx, Event{
			Title:       ce.Title,
			Description: ce.Description,
			TicketURL:   ce.TicketURL,
			ImageURL:    ce.ImageURL,
			VenueID:     int64(ce.VenueID),
		})

		if err != nil {
			return fmt.Errorf("Could not insert event: %v", err)
		}

		concerts := make([]Concert, 0)
		for _, c := range ce.Concerts {
			concerts = append(concerts, Concert{
				FromDate: c.From,
				ToDate:   c.To,
				ArtistID: c.ArtistID,
				EventID:  insertedID,
			})
		}

		if err := setEventConcerts(ctx, tx, insertedID, concerts...); err != nil {
			return fmt.Errorf("Could not insert event concerts: %v", err)
		}

		eventID = insertedID
		return nil
	})

	if err != nil {
		return 0, err
	}

	return eventID, nil
}

// ListEvents implements konnekt.EventRepo.
func (e EventRepo) ListEvents(ctx context.Context, lr api.ListRequest) (api.ListResponse[konnekt.Event], error) {
	events := make([]konnekt.Event, 0)
	err := pgx.BeginFunc(ctx, e.Pool, func(tx pgx.Tx) error {
		pg := paginationFromListRequest(lr)
		dbEvents, err := listEvents(ctx, tx, pg, lr.OrderMap)
		if err != nil {
			return fmt.Errorf("Could not list events: %v", err)
		}

		events, err = dbEvents.Assemble(ctx, tx)
		if err != nil {
			return fmt.Errorf("Could not assemble DB events: %v", err)
		}

		return nil
	})

	if err != nil {
		return api.ListResponse[konnekt.Event]{}, err
	}

	return api.ListResponse[konnekt.Event]{
		Records: events,
	}, nil
}

// Update implements konnekt.EventRepo.
func (e EventRepo) Update(ctx context.Context, eventID int64, ur api.UpdateRequest[konnekt.UpdateEvent]) error {
	err := pgx.BeginFunc(ctx, e.Pool, func(tx pgx.Tx) error {
		um := ur.UpdateMap()
		if err := updateEvent(ctx, tx, eventID, um); err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return konnekt.ErrResourceNotFound
			}

			return fmt.Errorf("Could not update event: %v", err)
		}

		if _, ok := um["concerts"]; ok {
			concerts := make([]Concert, 0)
			for _, c := range ur.Data.Concerts {
				concerts = append(concerts, Concert{
					FromDate: c.From,
					ToDate:   c.To,
					ArtistID: c.ArtistID,
					EventID:  eventID,
				})
			}

			if err := setEventConcerts(ctx, tx, eventID, concerts...); err != nil {
				return fmt.Errorf("Could not set event concerts: %v", err)
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func insertEvent(ctx context.Context, tx pgx.Tx, e Event) (int64, error) {
	query, args, err := psql.
		Insert("event").
		Columns("title", "description", "ticket_url", "image_url", "venue_id").
		Values(e.Title, e.Description, e.TicketURL, e.ImageURL, e.VenueID).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return 0, NewQueryBuildError("insert event", err)
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return 0, fmt.Errorf("Could not execute insert event query: %v", err)
	}

	id, err := pgx.CollectExactlyOneRow(rows, pgx.RowTo[int64])
	if err != nil {
		return 0, fmt.Errorf("Could not get inserted ID: %v", err)
	}

	return id, nil
}

func eventByID(ctx context.Context, tx pgx.Tx, eventID int64) (Event, error) {
	query, args, err := psql.
		Select("*").
		From("event").
		Where(sq.Eq{"id": eventID}).
		ToSql()

	if err != nil {
		return Event{}, NewQueryBuildError("event by id", err)
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Event{}, err
		}

		return Event{}, fmt.Errorf("Could not query for event with ID: %d: %v", eventID, err)
	}

	event, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[Event])
	if err != nil {
		return Event{}, fmt.Errorf("Could not collect event: %v", err)
	}

	return event, nil
}

func listEvents(ctx context.Context, tx pgx.Tx, pg Pagination, om order.Map) (Collection[Event, konnekt.Event], error) {
	builder := psql.Select("*").From("event")
	builder = applyPagination(builder, pg)
	builder = applyOrdering(builder, om)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, NewQueryBuildError("list events", err)
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("Could not query for events: %v", err)
	}

	events, err := pgx.CollectRows(rows, pgx.RowToStructByName[Event])
	if err != nil {
		return nil, fmt.Errorf("Could not collect events: %v", err)
	}

	return events, nil
}

func updateEvent(ctx context.Context, tx pgx.Tx, eventID int64, um mask.FieldMap) error {
	exists, err := checkIfExists(ctx, tx, "event", "id", eventID)
	if err != nil {
		return err
	}

	if !exists {
		return pgx.ErrNoRows
	}

	allowed := []string{"title", "description", "image_url", "ticket_url", "is_public", "venue_id"}

	updates := make(map[string]any)
	for _, key := range allowed {
		if v, ok := um[mask.FieldName(key)]; ok {
			updates[key] = v
		}
	}

	if len(updates) == 0 {
		return nil
	}

	query, args, err := psql.
		Update("event").
		Where(sq.Eq{"id": eventID}).
		SetMap(updates).
		ToSql()

	if err != nil {
		return NewQueryBuildError("update event", err)
	}

	if _, err := tx.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("Could not update event: %v", err)
	}

	return nil
}

func (e Event) Assemble(ctx context.Context, tx pgx.Tx) (konnekt.Event, error) {
	dbConcerts, err := eventConcerts(ctx, tx, e.ID)
	if err != nil {
		return konnekt.Event{}, fmt.Errorf("Could not get event concerts: %v", err)
	}

	dbVenue, err := venueByID(ctx, tx, e.VenueID)
	if err != nil {
		return konnekt.Event{}, fmt.Errorf("Could not get event venue: %v", err)
	}

	concerts, err := dbConcerts.Assemble(ctx, tx)
	if err != nil {
		return konnekt.Event{}, fmt.Errorf("Could not assemble event concerts: %v", err)
	}

	venue, err := dbVenue.Assemble(ctx, tx)
	if err != nil {
		return konnekt.Event{}, fmt.Errorf("Could not assemble event venue: %v", err)
	}

	return konnekt.Event{
		ID:          e.ID,
		Title:       e.Title,
		Description: e.Description,
		TicketURL:   e.TicketURL,
		ImageURL:    e.ImageURL,
		Concerts:    concerts,
		Venue:       venue,
	}, nil
}
