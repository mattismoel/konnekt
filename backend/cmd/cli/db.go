package main

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	konnekt "github.com/mattismoel/konnekt/backend"
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
		approved BOOLEAN NOT NULL DEFAULT FALSE,

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
		id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		session_id TEXT UNIQUE NOT NULL,
		secret_hash BYTEA NOT NULL,
		created_at TIMESTAMP NOT NULL
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
	teams := []konnekt.Team{
		{Name: "admin", DisplayName: "Administrator", Description: "Administrator af hjemmesiden."},
		{Name: "member", DisplayName: "Medlem", Description: "Medlem af foreningen."},
		{Name: "project-leader", DisplayName: "Projektleder", Description: "Projektleder af foreningen."},
		{Name: "booking", DisplayName: "Booking", Description: "Booking af kunstnere."},
		{Name: "public-relations", DisplayName: "PR", Description: "Håndtering af foreningens offentlige og medie-mæssige tilstedeværelse."},
		{Name: "visual-identity", DisplayName: "Visuel Identitet", Description: "Håndtering af foreningens visuelle identitet og design."},
		{Name: "event-management", DisplayName: "Event-management", Description: "Håndtering, planlægning og afvikling foreningens events"},
		{Name: "economy", DisplayName: "Økonomi", Description: "Håndtering af foreningens økonomi."},
	}

	builder := psql.
		Insert("team").
		Columns("name", "display_name", "description").
		Suffix("ON CONFLICT DO NOTHING")

	for _, t := range teams {
		builder = builder.Values(t.Name, t.DisplayName, t.Description)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("Could not create team seed query: %v", err)
	}

	if _, err := tx.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("Could not insert seed teams: %v", err)
	}

	return nil
}

func seedPermissions(ctx context.Context, tx pgx.Tx) error {
	perms := []konnekt.Permission{
		{Name: "view:event", DisplayName: "Events (se)", Description: "Tilladeslse til at se events."},
		{Name: "edit:event", DisplayName: "Events (redigér)", Description: "Tilladelse til at redigére events."},
		{Name: "delete:event", DisplayName: "Events (slet)", Description: "Tilladelse til at slette events."},

		{Name: "view:artist", DisplayName: "Kunstnere (se)", Description: "Tilladeslse til at se kunstnere."},
		{Name: "edit:artist", DisplayName: "Kunstnere (redigér)", Description: "Tilladelse til at redigére kunstnere."},
		{Name: "delete:artist", DisplayName: "Kunstnere (slet)", Description: "Tilladelse til at slette kunstnere."},

		{Name: "view:venue", DisplayName: "Venues (se)", Description: "Tilladeslse til at se venues."},
		{Name: "edit:venue", DisplayName: "Venues (redigér)", Description: "Tilladelse til at redigére venues."},
		{Name: "delete:venue", DisplayName: "Venues (slet)", Description: "Tilladelse til at slette venues."},

		{Name: "view:genre", DisplayName: "Genrer (se)", Description: "Tilladeslse til at se genrer."},
		{Name: "edit:genre", DisplayName: "Genrer (redigér)", Description: "Tilladelse til at redigére genrer."},
		{Name: "delete:genre", DisplayName: "Genrer (slet)", Description: "Tilladelse til at slette genrer."},

		{Name: "view:content", DisplayName: "Indhold (se)", Description: "Tilladeslse til at se hjemmesideindhold."},
		{Name: "edit:content", DisplayName: "Indhold (redigér)", Description: "Tilladelse til at redigére hjemmesideindhold."},
		{Name: "delete:content", DisplayName: "Indhold (slet)", Description: "Tilladelse til at slette hjemmesideindhold."},

		{Name: "view:member", DisplayName: "Medlemmer (se)", Description: "Tilladeslse til at se medlemmer."},
		{Name: "edit:member", DisplayName: "Medlemmer (redigér)", Description: "Tilladelse til at redigére medlemmer."},
		{Name: "delete:member", DisplayName: "Medlemmer (slet)", Description: "Tilladelse til at slette medlemmer."},

		{Name: "view:team", DisplayName: "Hold (se)", Description: "Tilladeslse til at se hold."},
		{Name: "edit:team", DisplayName: "Hold (redigér)", Description: "Tilladelse til at redigére hold."},
		{Name: "delete:team", DisplayName: "Hold (slet)", Description: "Tilladelse til at slette hold."},

		{Name: "view:permission", DisplayName: "Tilladelser (se)", Description: "Tilladeslse til at se tilladelse."},
		{Name: "edit:permission", DisplayName: "Tilladelser (redigér)", Description: "Tilladelse til at redigére tilladelser."},
		{Name: "delete:permission", DisplayName: "Tilladelser (slet)", Description: "Tilladelse til at slette tilladelser.}"},
	}

	builder := psql.
		Insert("permission").
		Columns("name", "display_name", "description").
		Suffix("ON CONFLICT DO NOTHING")

	for _, p := range perms {
		builder = builder.Values(p.Name, p.DisplayName, p.Description)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("Could not create seed permission query: %v", err)
	}

	if _, err := tx.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("Could not insert seed permissions: %v", err)
	}

	return nil
}

func seedGenres(ctx context.Context, tx pgx.Tx) error {
	genres := []konnekt.Genre{"Rock", "Punk", "Pop", "Hip-Hop", "Rap", "Folk",
		"Indie", "R&B", "Elektronisk", "Country", "Jazz", "Blues", "Funk", "Latin"}

	builder := psql.Insert("genre").Columns("name")

	for _, genreName := range genres {
		builder = builder.Values(genreName)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("Could not create seed genres query: %v", err)
	}

	if _, err := tx.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("Could not insert genres: %v", err)
	}

	return nil
}
