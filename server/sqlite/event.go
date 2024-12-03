package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/mattismoel/konnekt"
)

// FindEventByID(context.Context, int64) (Event, error)
// FindEvents(context.Context, EventFilter) ([]Event, error)
// CreateEvent(context.Context, Event) (int64, error)
// UpdateEvent(context.Context, int64, EventUpdate) (Event, error)
// DeleteEvent(context.Context, int64) error

type eventService struct {
	repo *Repository
}

func NewEventService(repo *Repository) *eventService {
	return &eventService{repo: repo}
}

func (s eventService) CreateEvent(ctx context.Context, event konnekt.Event) (konnekt.Event, error) {
	tx, err := s.repo.db.BeginTx(ctx, nil)
	if err != nil {
		return konnekt.Event{}, err
	}

	defer tx.Rollback()

	addressId, err := createAddress(ctx, tx, event.Address)
	if err != nil {
		return konnekt.Event{}, fmt.Errorf("Could not create address: %v", err)
	}

	event.Genres, err = insertGenres(ctx, tx, event.Genres)
	if err != nil {
		return konnekt.Event{}, fmt.Errorf("Could not insert genres: %v", err)
	}

	event.ID, err = createEvent(ctx, tx, addressId, event)
	if err != nil {
		return konnekt.Event{}, err
	}

	err = setEventGenres(ctx, tx, event.ID, genreNames(event.Genres))
	if err != nil {
		return konnekt.Event{}, fmt.Errorf("Could not set event genres: %v", err)
	}

	if err = tx.Commit(); err != nil {
		return konnekt.Event{}, err
	}

	return event, nil
}

func (s eventService) FindEventByID(ctx context.Context, id int64) (konnekt.Event, error) {
	tx, err := s.repo.db.BeginTx(ctx, nil)
	if err != nil {
		return konnekt.Event{}, err
	}

	defer tx.Rollback()

	event, err := findEventByID(ctx, tx, id)
	if err != nil {
		return konnekt.Event{}, err
	}

	if err = tx.Commit(); err != nil {
		return konnekt.Event{}, err
	}

	return event, nil
}

func (s eventService) FindEvents(ctx context.Context, filter konnekt.EventFilter) ([]konnekt.Event, error) {
	tx, err := s.repo.db.BeginTx(ctx, nil)
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

func (s eventService) UpdateEvent(ctx context.Context, id int64, update konnekt.EventUpdate) (konnekt.Event, error) {
	tx, err := s.repo.db.BeginTx(ctx, nil)
	if err != nil {
		return konnekt.Event{}, err
	}

	defer tx.Rollback()

	event, err := updateEvent(ctx, tx, id, update)
	if err != nil {
		return konnekt.Event{}, err
	}

	if err := tx.Commit(); err != nil {
		return konnekt.Event{}, err
	}

	return event, nil

}

func (s eventService) DeleteEvent(ctx context.Context, id int64) error {
	tx, err := s.repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	if err = deleteEvent(ctx, tx, id); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func createEvent(ctx context.Context, tx *sql.Tx, addressId int64, event konnekt.Event) (int64, error) {
	err := event.Validate()
	if err != nil {
		return 0, err
	}

	query := `
	INSERT INTO event (
		title,
		description,
		from_date,
		to_date,
		address_id
	) 
	VALUES (?, ?, ?, ?, ?)`

	res, err := tx.ExecContext(ctx, query,
		event.Title,
		event.Description,
		event.FromDate.UnixMilli(),
		event.ToDate.UnixMilli(),
		addressId,
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

func findEventByID(ctx context.Context, tx *sql.Tx, id int64) (konnekt.Event, error) {
	events, _, err := findEvents(ctx, tx, konnekt.EventFilter{ID: &id})
	if err != nil {
		return konnekt.Event{}, err
	}

	if len(events) == 0 {
		return konnekt.Event{}, konnekt.Error{Code: konnekt.ERRNOTFOUND, Message: "Event not found"}
	}

	return events[0], nil

}

func findEvents(ctx context.Context, tx *sql.Tx, filter konnekt.EventFilter) ([]konnekt.Event, int, error) {
	events := []konnekt.Event{}
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
		var event konnekt.Event

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

		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return events, n, nil
}

func updateEvent(ctx context.Context, tx *sql.Tx, id int64, update konnekt.EventUpdate) (konnekt.Event, error) {
	event, err := findEventByID(ctx, tx, id)
	if err != nil {
		return konnekt.Event{}, err
	}

	if v := update.Title; v != nil {
		event.Title = *v
	}

	if v := update.Description; v != nil {
		event.Description = *v
	}

	if v := update.FromDate; !v.IsZero() {
		event.FromDate = v
	}

	if v := update.ToDate; !v.IsZero() {
		event.ToDate = v
	}

	if err := event.Validate(); err != nil {
		return konnekt.Event{}, err
	}

	query := `
	UPDATE event
	SET
		title = ?,
		description = ?,
		from_date = ?,
		to_date = ?
	WHERE id = ?`

	_, err = tx.ExecContext(
		ctx, query,
		event.Title,
		event.Description,
		event.FromDate.UnixMilli(),
		event.ToDate.UnixMilli(),
		id,
	)

	if err != nil {
		return konnekt.Event{}, err
	}

	// Return early, if address is not to be updated.
	if update.Address != nil {
		event.Address, err = updateEventAddress(ctx, tx, event.ID, *update.Address)
		if err != nil {
			return konnekt.Event{}, err
		}
	}

	if update.GenreNames == nil || len(update.GenreNames) <= 0 {
		return event, nil
	}

	for i, genreName := range update.GenreNames {
		exists, err := doesGenreExist(ctx, tx, genreName)
		if err != nil {
			return konnekt.Event{}, err
		}

		if !exists {
			event.Genres[i], err = insertGenre(ctx, tx, konnekt.Genre{Name: genreName})
			if err != nil {
				return konnekt.Event{}, err
			}
		}
	}

	err = setEventGenres(ctx, tx, event.ID, update.GenreNames)
	if err != nil {
		return konnekt.Event{}, err
	}

	return event, err
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
