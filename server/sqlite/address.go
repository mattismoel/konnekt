package sqlite

import (
	"context"
	"database/sql"

	"github.com/mattismoel/konnekt"
)

func createAddress(ctx context.Context, tx *sql.Tx, address konnekt.Address) (int64, error) {
	query := `
	INSERT INTO address (
		country,
		city,
		street,
		house_number
	) 
	VALUES (?, ?, ?, ?)`

	res, err := tx.ExecContext(ctx, query,
		&address.Country,
		&address.City,
		&address.Street,
		&address.HouseNumber,
	)

	if err != nil {
		return 0, err
	}

	insertedID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return insertedID, nil
}

func findEventAddress(ctx context.Context, tx *sql.Tx, eventId int64) (konnekt.Address, error) {
	query := `
	SELECT
		address_id
	FROM event 
	WHERE id = ?`

	var addressID int64

	err := tx.QueryRowContext(ctx, query, eventId).Scan(&addressID)
	if err != nil {
		return konnekt.Address{}, err
	}

	query = `
	SELECT
		id,
		country,
		city,
		street,
		house_number
	FROM address
	WHERE id = ?`

	var address konnekt.Address

	err = tx.QueryRowContext(ctx, query, addressID).Scan(
		&address.ID,
		&address.Country,
		&address.City,
		&address.Street,
		&address.HouseNumber,
	)

	if err != nil {
		return konnekt.Address{}, err
	}

	return address, nil
}

func updateEventAddress(ctx context.Context, tx *sql.Tx, eventId int64, update konnekt.AddressUpdate) (konnekt.Address, error) {
	address, err := findEventAddress(ctx, tx, eventId)
	if err != nil {
		return konnekt.Address{}, err
	}

	if v := update.Country; v != nil {
		address.Country = *v
	}

	if v := update.City; v != nil {
		address.City = *v
	}

	if v := update.Street; v != nil {
		address.Street = *v
	}

	if v := update.HouseNumber; v != nil {
		address.HouseNumber = *v
	}

	if err = address.Validate(); err != nil {
		return konnekt.Address{}, err
	}

	query := `
	UPDATE address
	SET
		country = ?,
		city = ?,
		street = ?,
		house_number = ?`

	_, err = tx.ExecContext(ctx, query,
		address.Country,
		address.City,
		address.Street,
		address.HouseNumber,
	)

	if err != nil {
		return konnekt.Address{}, err
	}

	return address, nil
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
