package main

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	konnekt "github.com/mattismoel/konnekt/backend"
)

var teams = []konnekt.Team{
	{Name: "admin", DisplayName: "Administrator", Description: "Administrator af hjemmesiden."},
	{Name: "member", DisplayName: "Medlem", Description: "Medlem af foreningen."},
	{Name: "project-leader", DisplayName: "Projektleder", Description: "Projektleder af foreningen."},
	{Name: "booking", DisplayName: "Booking", Description: "Booking af kunstnere."},
	{Name: "public-relations", DisplayName: "PR", Description: "Håndtering af foreningens offentlige og medie-mæssige tilstedeværelse."},
	{Name: "visual-identity", DisplayName: "Visuel Identitet", Description: "Håndtering af foreningens visuelle identitet og design."},
	{Name: "event-management", DisplayName: "Event-management", Description: "Håndtering, planlægning og afvikling foreningens events"},
	{Name: "economy", DisplayName: "Økonomi", Description: "Håndtering af foreningens økonomi."},
}

var perms = []konnekt.Permission{
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

var teamPermissions = map[string][]string{
	"admin": {
		"view:event", "edit:event", "delete:event",
		"view:artist", "edit:artist", "delete:artist",
		"view:venue", "edit:venue", "delete:venue",
		"view:genre", "edit:genre", "delete:genre",
		"view:content", "edit:content", "delete:content",
		"view:member", "edit:member", "delete:member",
		"view:team",
		"view:permission",
	},
	"member": {
		"view:event",
		"view:artist",
		"view:venue",
		"view:genre",
		"view:content",
		"view:member",
		"view:team",
		"view:permission",
	},
	"project-leader": {
		"view:event", "edit:event", "delete:event",
		"view:artist", "edit:artist", "delete:artist",
		"view:member", "delete:member",
	},
	"booking": {
		"view:artist", "edit:artist", "delete:artist",
		"view:venue", "edit:venue", "delete:venue",
		"view:genre", "edit:genre", "delete:genre",
	},
	"public-relations": {
		"view:event", "edit:event", "delete:event",
		"view:venue", "edit:venue", "delete:venue",
		"view:content", "edit:content", "delete:content",
	},
	"visual-identity": {
		"view:event", "edit:event",
		"view:artist", "edit:artist",
		"view:content", "edit:content", "delete:content",
	},
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

func seedTeamPermissions(ctx context.Context, tx pgx.Tx) error {
	for teamName, perms := range teamPermissions {
		query, args, err := psql.
			Select("id").
			From("team").
			Where(sq.Eq{"name": teamName}).
			ToSql()

		if err != nil {
			return fmt.Errorf("Could not create 'get team id' query: %v", err)
		}

		rows, err := tx.Query(ctx, query, args...)
		if err != nil {
			return fmt.Errorf("Could not get team with name %q: %v", teamName, err)
		}

		teamID, err := pgx.CollectExactlyOneRow(rows, pgx.RowTo[int64])
		if err != nil {
			return fmt.Errorf("Could not collect team ID: %v", err)
		}

		for _, permName := range perms {
			query, args, err := psql.
				Select("id").
				From("permission").
				Where(sq.Eq{"name": permName}).
				ToSql()

			if err != nil {
				return fmt.Errorf("Could not create 'get permission ID' query: %v", err)
			}

			rows, err := tx.Query(ctx, query, args...)
			if err != nil {
				return fmt.Errorf("Could not query for permission with name %q: %v", permName, err)
			}

			permID, err := pgx.CollectExactlyOneRow(rows, pgx.RowTo[int64])
			if err != nil {
				return fmt.Errorf("Could not collect permission ID: %v", err)
			}

			query, args, err = psql.
				Insert("teams_permissions").
				Columns("team_id", "permission_id").
				Values(teamID, permID).
				Suffix("ON CONFLICT DO NOTHING").
				ToSql()

			if err != nil {
				return fmt.Errorf("Could not create 'insert team permission' query: %v", err)
			}

			if _, err := tx.Exec(ctx, query, args...); err != nil {
				return fmt.Errorf("Could not insert team permission: %v", err)
			}
		}
	}

	return nil
}

func seedTeams(ctx context.Context, tx pgx.Tx) error {
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
