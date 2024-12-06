package sqlite

import (
	"context"
	"database/sql"

	"github.com/mattismoel/konnekt/internal/storage"
)

func createAddress(ctx context.Context, tx *sql.Tx, address storage.Address) (storage.Address, error) {
	query := `
	INSERT INTO address (
		country,
		city,
		street,
		house_number
	) 
	VALUES (
		@country,
		@city,
		@street,
		@house_number
	)`

	res, err := tx.ExecContext(ctx, query,
		sql.Named("country", address.Country),
		sql.Named("city", address.City),
		sql.Named("street", address.Street),
		sql.Named("house_number", address.HouseNumber),
	)

	if err != nil {
		return storage.Address{}, err
	}

	insertedID, err := res.LastInsertId()
	if err != nil {
		return storage.Address{}, err
	}

	address.ID = insertedID

	return address, nil
}

func findEventAddress(ctx context.Context, tx *sql.Tx, eventId int64) (storage.Address, error) {
	query := `
	SELECT
		address_id
	FROM event 
	WHERE id = @id`

	var addressID int64

	err := tx.QueryRowContext(ctx, query, sql.Named("id", eventId)).Scan(&addressID)
	if err != nil {
		return storage.Address{}, err
	}

	query = `
	SELECT
		id,
		country,
		city,
		street,
		house_number
	FROM address
	WHERE id = @id`

	var address storage.Address

	err = tx.QueryRowContext(ctx, query, sql.Named("id", addressID)).Scan(
		&address.ID,
		&address.Country,
		&address.City,
		&address.Street,
		&address.HouseNumber,
	)

	if err != nil {
		return storage.Address{}, err
	}

	return address, nil
}

// Updates an events address.
//
// Any non-nil-valued property of the updated address struct will be set.
func updateEventAddress(ctx context.Context, tx *sql.Tx, eventId int64, update storage.Address) (storage.Address, error) {
	address, err := findEventAddress(ctx, tx, eventId)
	if err != nil {
		return storage.Address{}, err
	}

	query := `
	UPDATE address
	SET
		country = CASE
			WHEN @country = '' THEN address.country
			ELSE @country
		END,
		city = CASE
			WHEN @city = '' THEN address.city
			ELSE @city
		END,
		street = CASE
			WHEN @street = '' THEN address.street
			ELSE @street
		END,
		house_number = CASE
			WHEN @house_number = '' THEN address.house_number
			ELSE @house_number
		END
	WHERE id = @address_id
	RETURNING 
		id, 
		country, 
		city, 
		street, 
		house_number
	`

	err = tx.QueryRowContext(ctx, query,
		sql.Named("country", update.Country),
		sql.Named("city", update.City),
		sql.Named("street", update.Street),
		sql.Named("house_number", update.HouseNumber),
		sql.Named("address_id", address.ID),
	).Scan(
		&update.ID,
		&update.Country,
		&update.City,
		&update.Street,
		&update.HouseNumber,
	)

	if err != nil {
		return storage.Address{}, err
	}

	return update, nil
}

func deleteEventAddress(ctx context.Context, tx *sql.Tx, eventID int64) error {
	address, err := findEventAddress(ctx, tx, eventID)
	if err != nil {
		return err
	}

	query := "DELETE FROM address WHERE id = ?"

	_, err = tx.ExecContext(ctx, query, address.ID)
	if err != nil {
		return err
	}

	return nil
}
