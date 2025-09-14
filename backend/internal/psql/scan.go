package psql

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Scanner interface {
	Scan(...any) error
}

func scan(s Scanner, dst ...any) error {
	if err := s.Scan(dst...); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return err
		}

		return fmt.Errorf("Could not scan into provided destinations: %v", err)
	}

	return nil
}
