package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

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

func createEventTables(ctx context.Context, tx pgx.Tx) error {
	query := `
	CREATE TABLE IF NOT EXISTS event (
		id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		ticket_url TEXT NOT NULL,
		image_url TEXT NOT NULL,
		venue_id INTEGER NOT NULL REFERENCES venue(id)
	);

	CREATE TABLE IF NOT EXISTS concert (
		id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		from_date TIMESTAMP NOT NULL,
		to_date TIMESTAMP NOT NULL,
		event_id INTEGER NOT NULL REFERENCES event(id),
		artist_id INTEGER NOT NULL REFERENCES artist(id)
	);`

	if _, err := tx.Exec(ctx, query); err != nil {
		return fmt.Errorf("Could not execute query: %v", err)
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

func createArtistTables(ctx context.Context, tx pgx.Tx) error {
	query := `
	CREATE TABLE IF NOT EXISTS artist (
		id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		name TEXT UNIQUE NOT NULL,
		description TEXT NOT NULL,
		image_url TEXT NOT NULL,
		preview_url TEXT,

		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
		created_by INTEGER NOT NULL REFERENCES member(id),
		updated_by INTEGER NOT NULL REFERENCES member(id)
	);


	CREATE TABLE IF NOT EXISTS social (
		id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		url TEXT NOT NULL,
		artist_id INTEGER NOT NULL REFERENCES artist(id)
	);

	CREATE TABLE IF NOT EXISTS genre (
		id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		name TEXT UNIQUE NOT NULL,

		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW()
	);

	CREATE TABLE IF NOT EXISTS artists_genres (
		artist_id INTEGER NOT NULL REFERENCES ARTIST(ID),
		genre_id INTEGER NOT NULL REFERENCES GENRE(ID),

		PRIMARY KEY (artist_id, genre_id)
	);

	CREATE TRIGGER update_artist_timestamps BEFORE INSERT OR UPDATE ON artist
	FOR EACH ROW EXECUTE FUNCTION update_timestamps();

	CREATE TRIGGER update_genre_timestamps BEFORE INSERT OR UPDATE ON genre 
	FOR EACH ROW EXECUTE FUNCTION update_timestamps();`

	if _, err := tx.Exec(ctx, query); err != nil {
		return fmt.Errorf("Could not execute query: %v", err)
	}

	return nil
}

func createVenueTable(ctx context.Context, tx pgx.Tx) error {
	query := `
	CREATE TABLE IF NOT EXISTS venue (
		id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		name TEXT UNIQUE NOT NULL,
		city TEXT NOT NULL,
		country TEXT NOT NULL,

		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
		created_by INTEGER NOT NULL REFERENCES member(id),
		updated_by INTEGER NOT NULL REFERENCES member(id)
	);

	CREATE TRIGGER update_venue_timestamps BEFORE INSERT OR UPDATE ON venue
	FOR EACH ROW EXECUTE FUNCTION update_timestamps();`

	if _, err := tx.Exec(ctx, query); err != nil {
		return fmt.Errorf("Could not execute query: %v", err)
	}

	return nil
}

func createMemberTables(ctx context.Context, tx pgx.Tx) error {
	query := `
	CREATE TABLE IF NOT EXISTS member (
		id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		email TEXT UNIQUE NOT NULL,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		password_hash CHAR(60) NOT NULL,
		avatar_url TEXT NOT NULL,
		special_role TEXT,
		active BOOLEAN NOT NULL DEFAULT FALSE,
		approved_by_id INTEGER REFERENCES member(id),

		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW()
	);

	CREATE TRIGGER update_member_timestamps BEFORE INSERT OR UPDATE ON member
	FOR EACH ROW EXECUTE FUNCTION update_timestamps();`

	if _, err := tx.Exec(ctx, query); err != nil {
		return fmt.Errorf("Could not execute query: %v", err)
	}

	return nil
}

func createSessionTable(ctx context.Context, tx pgx.Tx) error {
	query := `
	CREATE TABLE IF NOT EXISTS session (
		id TEXT PRIMARY KEY,
		member_id INTEGER NOT NULL REFERENCES member(id),
		expires_at TIMESTAMP NOT NULL
	);`

	if _, err := tx.Exec(ctx, query); err != nil {
		return fmt.Errorf("Could not execute query: %v", err)
	}

	return nil
}

func createAuthTables(ctx context.Context, tx pgx.Tx) error {
	query := `
	CREATE TABLE IF NOT EXISTS team (
		id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		name TEXT UNIQUE NOT NULL,
		display_name TEXT NOT NULL,
		description TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS permission (
		id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		name TEXT UNIQUE NOT NULL,
		display_name TEXT NOT NULL,
		description TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS members_teams (
		member_id INTEGER NOT NULL REFERENCES member(id),
		team_id INTEGER NOT NULL REFERENCES team(id),
		PRIMARY KEY (member_id, team_id)
	);

	CREATE TABLE IF NOT EXISTS teams_permissions (
		team_id INTEGER NOT NULL REFERENCES team(id),
		permission_id INTEGER NOT NULL REFERENCES permission(id),
		PRIMARY KEY (team_id, permission_id)
	);`

	if _, err := tx.Exec(ctx, query); err != nil {
		return fmt.Errorf("Could not execute query: %v", err)
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

func seedTeams(ctx context.Context, tx pgx.Tx) error {
	query := `
	INSERT INTO team (name, display_name, description) VALUES 
	('admin', 'Administrator', 'Administrator af hjemmesiden.'),
	('member', 'Medlem', 'Medlem af foreningen.'),
	('project-leader', 'Projektleder', 'Projektleder af foreningen.'),
	('booking', 'Booking', 'Booking af kunstnere.'),
	('public-relations', 'PR', 'Håndtering af foreningens offentlige og medie-mæssige tilstedeværelse.'),
	('visual-identity', 'Visuel Identitet', 'Håndtering af foreningens visuelle identitet og design.'),
	('event-management', 'Event-management', 'Håndtering, planlægning og afvikling foreningens events'),
	('economy', 'Økonomi', 'Håndtering af foreningens økonomi.')

	ON CONFLICT DO NOTHING;`

	if _, err := tx.Exec(ctx, query); err != nil {
		return fmt.Errorf("Could not insert seed teams: %v", err)
	}

	return nil
}

func seedPermissions(ctx context.Context, tx pgx.Tx) error {
	query := `
	INSERT INTO permission (name, display_name, description) VALUES
	('view:event', 'Events (se)', 'Tilladeslse til at se events.'),
	('edit:event', 'Events (redigér)', 'Tilladelse til at redigére events.'),
	('delete:event', 'Events (slet)', 'Tilladelse til at slette events.'),

	('view:artist', 'Kunstnere (se)', 'Tilladeslse til at se kunstnere.'),
	('edit:artist', 'Kunstnere (redigér)', 'Tilladelse til at redigére kunstnere.'),
	('delete:artist', 'Kunstnere (slet)', 'Tilladelse til at slette kunstnere.'),

	('view:venue', 'Venues (se)', 'Tilladeslse til at se venues.'),
	('edit:venue', 'Venues (redigér)', 'Tilladelse til at redigére venues.'),
	('delete:venue', 'Venues (slet)', 'Tilladelse til at slette venues.'),

	('view:genre', 'Genrer (se)', 'Tilladeslse til at se genrer.'),
	('edit:genre', 'Genrer (redigér)', 'Tilladelse til at redigére genrer.'),
	('delete:genre', 'Genrer (slet)', 'Tilladelse til at slette genrer.'),

	('view:content', 'Indhold (se)', 'Tilladeslse til at se hjemmesideindhold.'),
	('edit:content', 'Indhold (redigér)', 'Tilladelse til at redigére hjemmesideindhold.'),
	('delete:content', 'Indhold (slet)', 'Tilladelse til at slette hjemmesideindhold.'),

	('view:member', 'Medlemmer (se)', 'Tilladeslse til at se medlemmer.'),
	('edit:member', 'Medlemmer (redigér)', 'Tilladelse til at redigére medlemmer.'),
	('delete:member', 'Medlemmer (slet)', 'Tilladelse til at slette medlemmer.'),

	('view:team', 'Hold (se)', 'Tilladeslse til at se hold.'),
	('edit:team', 'Hold (redigér)', 'Tilladelse til at redigére hold.'),
	('delete:team', 'Hold (slet)', 'Tilladelse til at slette hold.'),

	('view:permission', 'Tilladelser (se)', 'Tilladeslse til at se tilladelse.'),
	('edit:permission', 'Tilladelser (redigér)', 'Tilladelse til at redigére tilladelser.'),
	('delete:permission', 'Tilladelser (slet)', 'Tilladelse til at slette tilladelser.')

	ON CONFLICT DO NOTHING;`

	if _, err := tx.Exec(ctx, query); err != nil {
		return fmt.Errorf("Could not insert seed permissions: %v", err)
	}

	return nil
}

func seedGenres(ctx context.Context, tx pgx.Tx) error {
	query := `
	INSERT INTO genre (name, created_by, updated_by) VALUES 
	('Rock', 1, 1),
	('Punk', 1, 1),
	('Pop', 1, 1),
	('Hip-Hip', 1, 1),
	('Rap', 1, 1),
	('Folk', 1, 1),
	('Klassisk', 1, 1);`

	if _, err := tx.Exec(ctx, query); err != nil {
		return fmt.Errorf("Could not insert genres: %v", err)
	}

	return nil
}
