package main

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

func clearDatabase(ctx context.Context, conn *pgx.Conn) error {
	err := pgx.BeginFunc(ctx, conn, func(tx pgx.Tx) error {
		if _, err := tx.Exec(ctx, "DROP SCHEMA public CASCADE"); err != nil {
			return fmt.Errorf("Could not drop schema %q: %v", "public", err)
		}

		if _, err := tx.Exec(ctx, "CREATE SCHEMA public"); err != nil {
			return fmt.Errorf("Could create schema %q: %v", "public", err)
		}

		if _, err := tx.Exec(ctx, "GRANT ALL ON SCHEMA public TO postgres"); err != nil {
			return fmt.Errorf("Could grant permissions on schema %q: %v", "public", err)
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func initialiseDatabase(ctx context.Context, conn *pgx.Conn) error {
	err := pgx.BeginFunc(ctx, conn, func(tx pgx.Tx) error {
		if err := createUpdateTrigger(ctx, tx); err != nil {
			return fmt.Errorf("Could not create update trigger: %v", err)
		}

		if err := createMemberTables(ctx, tx); err != nil {
			return fmt.Errorf("Could not create member tables: %v", err)
		}

		if err := createVenueTable(ctx, tx); err != nil {
			return fmt.Errorf("Could not create venue table: %v", err)
		}

		if err := createArtistTables(ctx, tx); err != nil {
			return fmt.Errorf("Could not create artist tables: %v", err)
		}

		if err := createEventTables(ctx, tx); err != nil {
			return fmt.Errorf("Could not create event tables: %v", err)
		}

		if err := createAuthTables(ctx, tx); err != nil {
			return fmt.Errorf("Could not create auth tables: %v", err)
		}

		if err := createSessionTable(ctx, tx); err != nil {
			return fmt.Errorf("Could not create session tables: %v", err)
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func createUpdateTrigger(ctx context.Context, tx pgx.Tx) error {
	query := `
	CREATE FUNCTION update_timestamps() RETURNS TRIGGER AS $$
	BEGIN
		NEW.created_at = 
		CASE 
			WHEN TG_OP = 'INSERT' THEN NOW() 
			ELSE OLD.created_at
		END;

		NEW.updated_at = 
		CASE 
			WHEN TG_OP = 'UPDATE' AND OLD.updated_at >= NOW() THEN OLD.updated_at + INTERVAL '1 millisecond' 
			ELSE NOW() 
		END;
		RETURN NEW;
	END;
	$$ LANGUAGE plpgsql;`

	if _, err := tx.Exec(ctx, query); err != nil {
		return fmt.Errorf("Could not create update timestamp trigger: %v", err)
	}

	return nil
}

func seedDb(ctx context.Context, conn *pgx.Conn) error {
	err := pgx.BeginFunc(ctx, conn, func(tx pgx.Tx) error {
		if err := seedTeams(ctx, tx); err != nil {
			return fmt.Errorf("Could not seed teams: %v", err)
		}

		if err := seedPermissions(ctx, tx); err != nil {
			return fmt.Errorf("Could not seed permissions: %v", err)
		}

		if err := seedTeamPermissions(ctx, tx); err != nil {
			return fmt.Errorf("Could not seed team permissions: %v", err)
		}

		if err := seedGenres(ctx, tx); err != nil {
			return fmt.Errorf("Could not seed genres: %v", err)
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
