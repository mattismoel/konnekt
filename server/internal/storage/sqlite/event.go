package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/mattismoel/konnekt"
	"github.com/mattismoel/konnekt/internal/prnt"
	"github.com/mattismoel/konnekt/internal/service"
	"github.com/mattismoel/konnekt/internal/storage"
)

type eventRepository struct {
	store *Store
}

func NewEventRepository(store *Store) *eventRepository {
	return &eventRepository{store: store}
}

type InsertEventRequest struct {
	Title        string
	Description  string
	FromDateUnix int64
	ToDateUnix   int64
}

func (s eventRepository) InsertEvent(ctx context.Context, baseEvent storage.BaseEvent, address storage.Address, genres []string) (storage.Event, error) {
	var event storage.Event

	tx, err := s.store.BeginTx(ctx, nil)
	if err != nil {
		return storage.Event{}, err
	}

	defer tx.Rollback()

	addr, err := createAddress(ctx, tx, address)
	if err != nil {
		return storage.Event{}, fmt.Errorf("Could not create address: %v", err)
	}

	baseEvent.AddressID = addr.ID

	event.Genres, err = insertGenres(ctx, tx, genres)
	if err != nil {
		return storage.Event{}, fmt.Errorf("Could not insert genres: %v", err)
	}

	event.ID, err = insertBaseEvent(ctx, tx, baseEvent)
	if err != nil {
		return storage.Event{}, err
	}

	err = updateEventGenres(ctx, tx, event.ID, genres)
	if err != nil {
		return storage.Event{}, fmt.Errorf("Could not set event genres: %v", err)
	}

	if err = tx.Commit(); err != nil {
		return storage.Event{}, err
	}

	return event, nil
}

func (s eventRepository) FindEventByID(ctx context.Context, id int64) (storage.Event, error) {
	tx, err := s.store.BeginTx(ctx, nil)
	if err != nil {
		return storage.Event{}, err
	}

	defer tx.Rollback()

	event, err := findEventByID(ctx, tx, id)
	if err != nil {
		return storage.Event{}, err
	}

	if err = tx.Commit(); err != nil {
		return storage.Event{}, err
	}

	return event, nil
}

func (s eventRepository) FindEvents(ctx context.Context, filter service.EventFilter) ([]storage.Event, error) {
	tx, err := s.store.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	events, _, err := findEvents(ctx, tx, filter)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return events, nil
}

func (s eventRepository) UpdateEvent(ctx context.Context, id int64, updatedEvent storage.Event) (storage.Event, error) {
	tx, err := s.store.BeginTx(ctx, nil)
	if err != nil {
		return storage.Event{}, err
	}

	defer tx.Rollback()

	event, err := updateEvent(ctx, tx, id, updatedEvent)
	if err != nil {
		return storage.Event{}, err
	}

	event.Address, err = updateEventAddress(ctx, tx, id, updatedEvent.Address)
	if err != nil {
		return storage.Event{}, err
	}

	genreNames := []string{}
	genres := []storage.Genre{}

	for _, genre := range updatedEvent.Genres {
		genres = append(genres, genre)

		genreNames = append(genreNames, genre.Name)
	}

	if len(genreNames) > 0 {
		_, err := insertGenres(ctx, tx, genreNames)
		if err != nil {
			return storage.Event{}, err
		}

		err = updateEventGenres(ctx, tx, id, genreNames)
		if err != nil {
			return storage.Event{}, err
		}

		// event.Genres = genres
	}

	event.Genres, err = findEventGenres(ctx, tx, id)
	if err != nil {
		return storage.Event{}, err
	}

	if err := tx.Commit(); err != nil {
		return storage.Event{}, err
	}

	return event, nil
}

func (s eventRepository) DeleteEvent(ctx context.Context, id int64) error {
	tx, err := s.store.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	if err = deleteEventAddress(ctx, tx, id); err != nil {
		return err
	}

	if err = deleteEvent(ctx, tx, id); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func insertBaseEvent(ctx context.Context, tx *sql.Tx, event storage.BaseEvent) (int64, error) {
	query := `
	INSERT INTO event (
		title,
		description,
		from_date,
		to_date,
		address_id
	) VALUES (
		@title, 
		@description, 
		@from_date, 
		@to_date, 
		@address_id
	)`

	res, err := tx.ExecContext(ctx, query,
		sql.Named("title", event.Title),
		sql.Named("description", event.Description),
		sql.Named("from_date", event.FromDate.UnixMilli()),
		sql.Named("to_date", event.ToDate.UnixMilli()),
		sql.Named("address_id", event.AddressID),
	)

	if err != nil {
		return 0, err
	}

	insertedId, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Could not get inserted ID: %v", err)
	}

	return insertedId, nil
}

func findEventByID(ctx context.Context, tx *sql.Tx, id int64) (storage.Event, error) {
	events, _, err := findEvents(ctx, tx, service.EventFilter{ID: &id})
	if err != nil {
		return storage.Event{}, err
	}

	if len(events) == 0 {
		return storage.Event{}, konnekt.Error{Code: konnekt.ERRNOTFOUND, Message: "Event not found"}
	}

	return events[0], nil

}

func findEvents(ctx context.Context, tx *sql.Tx, filter service.EventFilter) ([]storage.Event, int, error) {
	events := []storage.Event{}
	var n int

	where, args := []string{"1 = 1\n"}, []any{}

	if v := filter.ID; v != nil {
		where, args = append(where, "id = ?\n"), append(args, v)
	}

	if v := filter.MinDate; !v.IsZero() {
		where, args = append(where, "from_date >= ?\n"), append(args, v.UnixMilli())
	}

	if v := filter.MaxDate; !v.IsZero() {
		where, args = append(where, "from_date <= ?\n"), append(args, v.UnixMilli())
	}

	query := fmt.Sprintf(`
	SELECT
		id,
		title,
		description,
		from_date,
		to_date,
		COUNT(*) OVER()
	FROM event
	WHERE `)

	query += strings.Join(where, " AND ")
	query += "ORDER BY from_date ASC\n"
	query += formatLimitOffset(filter.Limit, filter.Offset)

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	for rows.Next() {
		var event storage.Event

		var fromDateUnix int64
		var toDateUnix int64

		err := rows.Scan(
			&event.ID,
			&event.Title,
			&event.Description,
			&fromDateUnix,
			&toDateUnix,
			&n,
		)
		if err != nil {
			return nil, 0, err
		}

		event.FromDate = time.UnixMilli(fromDateUnix)
		event.ToDate = time.UnixMilli(toDateUnix)

		event.Address, err = findEventAddress(ctx, tx, event.ID)
		if err != nil {
			return nil, 0, err
		}

		event.Genres, err = findEventGenres(ctx, tx, event.ID)
		if err != nil {
			return nil, 0, err
		}

		fmt.Printf("REPO %s\n", prnt.Pretty(event))

		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return events, n, nil
}

// Updates the base event.
//
// This means that no address or genres are updated,
// as these should be updated with 'updateEventAddress' and 'updateEventGenres'.
func updateEvent(ctx context.Context, tx *sql.Tx, id int64, updatedEvent storage.Event) (storage.Event, error) {
	// _, err := findEventByID(ctx, tx, id)
	// if err != nil {
	// 	return storage.Event{}, err
	// }

	query := `
	UPDATE event
	SET
		title = CASE
			WHEN @title = '' THEN event.title
			ELSE @title
		END,
		description = CASE
			WHEN @description = '' THEN event.description
			ELSE @description
		END,
		from_date = CASE
			WHEN @from_date = 0 THEN event.from_date
			ELSE @from_date
		END,
		to_date = CASE
			WHEN @to_date = 0 THEN event.to_date
			ELSE @to_date
		END
	WHERE id = @id
	RETURNING
		id,
		title,
		description,
		from_date,
		to_date
	`

	var fromDateUnix, toDateUnix int64

	err := tx.QueryRowContext(
		ctx, query,
		sql.Named("title", updatedEvent.Title),
		sql.Named("description", updatedEvent.Description),
		sql.Named("from_date", updatedEvent.FromDate.UnixMilli()),
		sql.Named(
			"to_date",
			updatedEvent.ToDate.UnixMilli(),
		),
		sql.Named("id", id),
	).Scan(
		&updatedEvent.ID,
		&updatedEvent.Title,
		&updatedEvent.Description,
		&fromDateUnix,
		&toDateUnix,
	)

	if err != nil {
		return storage.Event{}, err
	}

	updatedEvent.FromDate = time.UnixMilli(fromDateUnix)
	updatedEvent.ToDate = time.UnixMilli(toDateUnix)

	// for i, genreName := range updatedEvent.Genres {
	// 	event.Genres[i], err = insertGenre(ctx, tx, genreName)
	// 	if err != nil {
	// 		return storage.Event{}, err
	// 	}
	// }

	// err = setEventGenres(ctx, tx, event.ID, event.Genres)
	// if err != nil {
	// 	return storage.Event{}, err
	// }
	//
	return updatedEvent, err
}

func deleteEvent(ctx context.Context, tx *sql.Tx, id int64) error {
	query := "DELETE FROM event WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	if err = deleteEventGenres(ctx, tx, id); err != nil {
		return err
	}

	return nil
}
