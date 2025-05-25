package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/mattismoel/konnekt/internal/domain/concert"
	"github.com/mattismoel/konnekt/internal/domain/event"
	"github.com/mattismoel/konnekt/internal/domain/venue"
	"github.com/mattismoel/konnekt/internal/query"
)

type EventRepository struct {
	db *sql.DB
}

var _ event.Repository = (*EventRepository)(nil)

type Event struct {
	ID          int64
	Title       string
	Description string
	TicketURL   string
	ImageURL    string
	VenueID     int64
}

func (e Event) ToInternal(venue venue.Venue, concerts []concert.Concert) event.Event {
	return event.Event{
		ID:          e.ID,
		Title:       e.Title,
		Description: e.Description,
		TicketURL:   e.TicketURL,
		ImageURL:    e.ImageURL,
		Venue:       venue,
		Concerts:    concerts,
	}
}

func NewEventRepository(db *sql.DB) (*EventRepository, error) {
	return &EventRepository{
		db: db,
	}, nil
}
func (repo EventRepository) ByID(ctx context.Context, eventID int64) (event.Event, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return event.Event{}, err
	}

	defer tx.Rollback()

	dbEvent, err := eventByID(ctx, tx, eventID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return event.Event{}, event.ErrNoExist
		}

		return event.Event{}, err
	}

	dbConcerts, err := eventConcerts(ctx, tx, eventID)
	if err != nil {
		return event.Event{}, err
	}

	concerts, err := dbConcerts.Internalize(ctx, tx)
	if err != nil {
		return event.Event{}, err
	}

	dbVenue, err := venueByID(ctx, tx, dbEvent.VenueID)
	if err != nil {
		return event.Event{}, err
	}

	venue := dbVenue.ToInternal()

	if err := tx.Commit(); err != nil {
		return event.Event{}, err
	}

	e := dbEvent.ToInternal(venue, concerts)

	return e, nil
}

func (repo EventRepository) Insert(ctx context.Context, e event.Event) (int64, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	defer tx.Rollback()

	eventID, err := insertEvent(ctx, tx, Event{
		Title:       e.Title,
		Description: e.Description,
		TicketURL:   e.TicketURL,
		ImageURL:    e.ImageURL,
		VenueID:     e.Venue.ID,
	})

	if err != nil {
		return 0, err
	}

	for _, c := range e.Concerts {
		dbArtist, err := artistByID(ctx, tx, c.Artist.ID)
		if err != nil {
			return 0, err
		}

		_, err = insertConcert(ctx, tx, Concert{
			ArtistID: dbArtist.ID,
			EventID:  eventID,
			From:     c.From,
			To:       c.To,
		})

		if err != nil {
			return 0, err
		}
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return eventID, nil
}

func (repo EventRepository) Update(ctx context.Context, eventID int64, e event.Event) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	err = updateEvent(ctx, tx, eventID, Event{
		Title:       e.Title,
		Description: e.Description,
		TicketURL:   e.TicketURL,
		ImageURL:    e.ImageURL,
		VenueID:     e.Venue.ID,
	})

	if err != nil {
		return err
	}

	concerts := make([]Concert, 0)
	for _, c := range e.Concerts {
		concerts = append(concerts, Concert{
			ID:       c.ID,
			EventID:  eventID,
			ArtistID: c.Artist.ID,
			From:     c.From,
			To:       c.To,
		})
	}

	_, err = setEventConcerts(ctx, tx, eventID, concerts...)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (repo EventRepository) Delete(ctx context.Context, eventID int64) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	if err := deleteEventConcerts(ctx, tx, eventID); err != nil {
		return err
	}

	if err := deleteEvent(ctx, tx, eventID); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (repo EventRepository) SetImageURL(ctx context.Context, eventID int64, coverImageURL string) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	err = setEventImageURL(ctx, tx, eventID, coverImageURL)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

type EventQueryParams struct {
	QueryParams
	From      time.Time
	To        time.Time
	ID        int64
	ArtistIDs []int64
}

func (repo EventRepository) List(ctx context.Context, q query.ListQuery) (query.ListResult[event.Event], error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return query.ListResult[event.Event]{}, err
	}

	defer tx.Rollback()

	dbEvents, err := listEvents(ctx, tx, EventQueryParams{
		QueryParams: QueryParams{
			Offset:  q.Offset(),
			Limit:   q.Limit,
			Filters: q.Filters,
			OrderBy: q.OrderBy,
		}},
	)

	if err != nil {
		return query.ListResult[event.Event]{}, err
	}

	totalCount, err := count(ctx, tx, "event")
	if err != nil {
		return query.ListResult[event.Event]{}, err
	}

	events := make([]event.Event, 0)

	for _, dbEvent := range dbEvents {
		dbVenue, err := venueByID(ctx, tx, dbEvent.VenueID)
		if err != nil {
			return query.ListResult[event.Event]{}, err
		}

		venue := dbVenue.ToInternal()

		dbConcerts, err := eventConcerts(ctx, tx, dbEvent.ID)
		if err != nil {
			return query.ListResult[event.Event]{}, err
		}

		concerts, err := dbConcerts.Internalize(ctx, tx)
		if err != nil {
			return query.ListResult[event.Event]{}, err
		}

		event := dbEvent.ToInternal(venue, concerts)
		events = append(events, event)
	}

	if err := tx.Commit(); err != nil {
		return query.ListResult[event.Event]{}, err
	}

	return query.ListResult[event.Event]{
		Page:       q.Page,
		PerPage:    q.PerPage,
		TotalCount: totalCount,
		PageCount:  q.PageCount(totalCount),
		Records:    events,
	}, nil
}

func eventByID(ctx context.Context, tx *sql.Tx, eventID int64) (Event, error) {
	event := sq.Select(
		"title",
		"description",
		"ticket_url",
		"image_url",
		"venue_id",
	).
		From("event").
		Where(sq.Eq{"id": eventID})

	query, args, err := event.ToSql()
	if err != nil {
		return Event{}, err
	}

	var title, description, ticketURL, coverImageURL string
	var venueID int64

	err = tx.QueryRowContext(ctx, query, args...).Scan(
		&title, &description, &ticketURL, &coverImageURL, &venueID,
	)

	if err != nil {
		return Event{}, err
	}

	return Event{
		ID:          eventID,
		Title:       title,
		Description: description,
		TicketURL:   ticketURL,
		ImageURL:    coverImageURL,
		VenueID:     venueID,
	}, nil
}


	if err != nil {
		return err
	}

	return nil
}

func listEvents(ctx context.Context, tx *sql.Tx, params EventQueryParams) ([]Event, error) {
	builder := sq.
		Select(
			"e.id",
			"e.title",
			"e.description",
			"e.ticket_url",
			"e.image_url",
			"e.venue_id",
		).
		Distinct().
		From("event e").
		Join("concert c ON c.event_id = e.id")

	if filters, ok := params.Filters["title"]; ok {
		for _, f := range filters {
			builder = builder.Where(sq.Like{"e.title": f.Value})
		}
	}

	if filters, ok := params.Filters["from_date"]; ok {
		for _, f := range filters {
			fmt.Printf("c.from_date %s %s", f.Cmp, f.Value)
			builder = builder.Where(fmt.Sprintf("c.from_date %s '%s'", f.Cmp, f.Value))
		}
	}

	if filters, ok := params.Filters["to_date"]; ok {
		for _, f := range filters {
			builder = builder.Where(fmt.Sprintf("c.to_date %s '%s'", f.Cmp, f.Value))
		}
	}

	if filters, ok := params.Filters["id"]; ok {
		for _, f := range filters {
			builder = builder.Where(sq.Eq{"e.id": f.Value})
		}
	}

	if filters, ok := params.Filters["artist_id"]; ok {
		for _, f := range filters {
			builder = builder.Where(sq.Eq{"c.artist_id": f.Value})
		}
	}

	if order, ok := params.OrderBy["from_date"]; ok {
		builder = builder.OrderBy("c.from_date " + string(order))
	} else {
		builder = builder.OrderBy("c.from_date ASC")
	}

	if params.Limit > 0 {
		builder = builder.Limit(uint64(params.Offset))
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

	events := make([]Event, 0)
	for rows.Next() {
		var id, venueID int64
		var title, description, ticketURL, coverImageURL string

		err := rows.Scan(&id, &title, &description, &ticketURL, &coverImageURL, &venueID)
		if err != nil {
			return nil, err
		}

		events = append(events, Event{
			ID:          id,
			Title:       title,
			Description: description,
			TicketURL:   ticketURL,
			ImageURL:    coverImageURL,
			VenueID:     venueID,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

func insertEvent(ctx context.Context, tx *sql.Tx, e Event) (int64, error) {
	query, args, err := sq.Insert("event").
		Columns("title", "description", "ticket_url", "image_url", "venue_id").
		Values(e.Title, e.Description, e.TicketURL, e.ImageURL, e.VenueID).ToSql()

	if err != nil {
		return 0, err
	}

	res, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	eventID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return eventID, nil
}

func setEventImageURL(ctx context.Context, tx *sql.Tx, eventID int64, url string) error {
	query, args, err := sq.
		Update("event").
		Set("image_url", url).
		Where(sq.Eq{"id": eventID}).
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

func updateEvent(ctx context.Context, tx *sql.Tx, eventID int64, e Event) error {
	builder := sq.Update("event").Where(sq.Eq{"id": eventID})

	if e.Title != "" {
		builder = builder.Set("title", e.Title)
	}

	if e.Description != "" {
		builder = builder.Set("description", e.Description)
	}

	if e.TicketURL != "" {
		builder = builder.Set("ticket_url", e.TicketURL)
	}

	if e.ImageURL != "" {
		builder = builder.Set("image_url", e.ImageURL)
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
	if rowsAffected <= 0 {
		return ErrNotFound
	}

	return nil
}

func deleteEvent(ctx context.Context, tx *sql.Tx, eventID int64) error {
	query, args, err := sq.
		Delete("event").
		Where(sq.Eq{"id": eventID}).
		ToSql()

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
