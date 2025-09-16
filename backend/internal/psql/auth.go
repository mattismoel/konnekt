package psql

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	konnekt "github.com/mattismoel/konnekt/backend"
	"github.com/mattismoel/konnekt/backend/api"
	"github.com/mattismoel/konnekt/backend/internal/server"
)

var _ server.AuthRepo = AuthRepo{}

type AuthRepo struct {
	Pool *pgxpool.Pool
}

type Team struct {
	ID          int64  `db:"id"`
	Name        string `db:"name"`
	DisplayName string `db:"display_name"`
	Description string `db:"description"`
}

type Permission struct {
	ID          int64  `db:"id"`
	Name        string `db:"name"`
	DisplayName string `db:"display_name"`
	Description string `db:"description"`
}

// MemberTeams implements konnekt.AuthRepo.
func (a AuthRepo) MemberTeams(ctx context.Context, memberID int64) (api.ListResponse[konnekt.Team], error) {
	teams := make([]konnekt.Team, 0)
	err := pgx.BeginFunc(ctx, a.Pool, func(tx pgx.Tx) error {
		dbTeams, err := memberTeams(ctx, tx, memberID)
		if err != nil {
			return fmt.Errorf("Could not get member teams: %v", err)
		}

		teams = dbTeams.ToDomain()
		return nil
	})

	if err != nil {
		return api.ListResponse[konnekt.Team]{}, err
	}

	return api.ListResponse[konnekt.Team]{
		Records: teams,
	}, nil
}

// TeamPermissions implements konnekt.AuthRepo.
func (a AuthRepo) TeamPermissions(ctx context.Context, teamID int64) (api.ListResponse[konnekt.Permission], error) {
	perms := make([]konnekt.Permission, 0)
	err := pgx.BeginFunc(ctx, a.Pool, func(tx pgx.Tx) error {
		dbPerms, err := teamPermissions(ctx, tx, teamID)
		if err != nil {
			return fmt.Errorf("Could not get team permisssions: %v", err)
		}

		perms = dbPerms.ToDomain()
		return nil
	})

	if err != nil {
		return api.ListResponse[konnekt.Permission]{}, err
	}

	return api.ListResponse[konnekt.Permission]{
		Records: perms,
	}, nil
}

func memberTeams(ctx context.Context, tx pgx.Tx, memberID int64) (Collection[Team, konnekt.Team], error) {
	query, args, err := psql.
		Select("team.*").
		From("team").
		Join("members_teams mt ON mt.team_id = team.id").
		Where(sq.Eq{"mt.member_id": memberID}).
		ToSql()

	if err != nil {
		return nil, NewQueryBuildError("member teams", err)
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("Could not query for member teams: %v", err)
	}

	teams, err := pgx.CollectRows(rows, pgx.RowToStructByName[Team])
	if err != nil {
		return nil, fmt.Errorf("Could not collect member teams: %v", err)
	}

	return teams, nil
}

func teamPermissions(ctx context.Context, tx pgx.Tx, teamID int64) (Collection[Permission, konnekt.Permission], error) {
	query, args, err := psql.
		Select("permission.*").
		From("permission").
		Join("teams_permissions tp ON tp.permission_id = permission.id").
		Where(sq.Eq{"tp.team_id": teamID}).
		ToSql()

	if err != nil {
		return nil, NewQueryBuildError("team permissions", err)
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("Could not query for team permissions: %v", err)
	}

	permissions, err := pgx.CollectRows(rows, pgx.RowToStructByName[Permission])
	if err != nil {
		return nil, fmt.Errorf("Could not collect member teams: %v", err)
	}

	return permissions, nil
}

func (t Team) ToDomain() konnekt.Team {
	return konnekt.Team{
		ID:          t.ID,
		Name:        t.Name,
		DisplayName: t.DisplayName,
		Description: t.Description,
	}
}

func (p Permission) ToDomain() konnekt.Permission {
	return konnekt.Permission{
		ID:          p.ID,
		Name:        p.Name,
		DisplayName: p.DisplayName,
		Description: p.Description,
	}
}
