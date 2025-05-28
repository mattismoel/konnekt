package sqlite

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/mattismoel/konnekt/internal/domain/team"
	"github.com/mattismoel/konnekt/internal/query"
)

var _ team.Repository = (*TeamRepository)(nil)

type TeamRepository struct {
	db *sql.DB
}

func NewTeamRepository(db *sql.DB) (*TeamRepository, error) {
	return &TeamRepository{db: db}, nil
}

type Team struct {
	ID          int64
	Name        string
	DisplayName string
	Description string
}

type TeamCollection []Team

func (r Team) ToInternal() team.Team {
	return team.Team{
		ID:          r.ID,
		Name:        r.Name,
		DisplayName: r.DisplayName,
		Description: r.Description,
	}
}

func (rc TeamCollection) ToInternal() team.TeamCollection {
	teams := make(team.TeamCollection, 0)

	for _, dbTeam := range rc {
		teams = append(teams, dbTeam.ToInternal())
	}

	return teams
}

func (repo TeamRepository) Insert(ctx context.Context, t team.Team) (int64, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	defer tx.Rollback()

	teamID, err := insertTeam(ctx, tx, Team{
		Name:        t.Name,
		DisplayName: t.DisplayName,
		Description: t.Description,
	})

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return teamID, nil
}

func (repo TeamRepository) List(ctx context.Context, q query.ListQuery) (query.ListResult[team.Team], error) {
	teams := []team.Team{}

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return query.ListResult[team.Team]{}, err
	}

	defer tx.Rollback()

	dbTeams, err := listTeams(ctx, tx, QueryParams{
		Offset:  q.Offset(),
		Limit:   q.Limit,
		OrderBy: q.OrderBy,
		Filters: q.Filters,
	})

	if err != nil {
		return query.ListResult[team.Team]{}, err
	}

	totalCount, err := count(ctx, tx, "team")
	if err != nil {
		return query.ListResult[team.Team]{}, err
	}

	if err := tx.Commit(); err != nil {
		return query.ListResult[team.Team]{}, err
	}

	for _, dbTeam := range dbTeams {
		teams = append(teams, dbTeam.ToInternal())
	}

	return query.ListResult[team.Team]{
		Records:    teams,
		Page:       q.Page,
		PerPage:    q.PerPage,
		TotalCount: totalCount,
		PageCount:  q.PageCount(totalCount),
	}, nil
}

func (repo TeamRepository) ByID(ctx context.Context, id int64) (team.Team, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return team.Team{}, err
	}

	defer tx.Rollback()

	dbTeam, err := teamByID(ctx, tx, id)
	if err != nil {
		return team.Team{}, err
	}

	if err := tx.Commit(); err != nil {
		return team.Team{}, err
	}

	return dbTeam.ToInternal(), nil
}

func (repo TeamRepository) ByName(ctx context.Context, name string) (team.Team, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return team.Team{}, err
	}

	defer tx.Rollback()

	dbTeam, err := teamByName(ctx, tx, name)
	if err != nil {
		return team.Team{}, err
	}

	if err := tx.Commit(); err != nil {
		return team.Team{}, err
	}

	return dbTeam.ToInternal(), nil
}

func (repo TeamRepository) Delete(ctx context.Context, teamID int64) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	if err := deleteTeam(ctx, tx, teamID); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (repo TeamRepository) AddMemberTeams(ctx context.Context, memberID int64, teamIDs ...int64) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	for _, teamID := range teamIDs {
		err := associateMemberWithTeam(ctx, tx, memberID, teamID)
		if err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (repo TeamRepository) MemberTeams(ctx context.Context, memberID int64) (team.TeamCollection, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	dbTeams, err := memberTeams(ctx, tx, memberID)
	if err != nil {
		return nil, err
	}

	teams := make([]team.Team, 0)

	for _, dbTeam := range dbTeams {
		teams = append(teams, dbTeam.ToInternal())
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return teams, nil
}

func insertTeam(ctx context.Context, tx *sql.Tx, t Team) (int64, error) {
	query, args, err := sq.
		Insert("team").
		Columns("name", "display_name", "description").
		Values(t.Name, t.DisplayName, t.Description).
		ToSql()

	if err != nil {
		return 0, err
	}

	res, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	teamID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return teamID, nil
}

var teamBuilder = sq.
	Select("id", "name", "description", "display_name").
	From("team")

func scanTeam(s Scanner, dst *Team) error {
	err := s.Scan(&dst.ID, &dst.Name, &dst.Description, &dst.DisplayName)
	if err != nil {
		return err
	}

	return nil
}

func listTeams(ctx context.Context, tx *sql.Tx, params QueryParams) (TeamCollection, error) {
	builder := teamBuilder.Distinct()

	if filters, ok := params.Filters["id"]; ok {
		for _, filter := range filters {
			builder.Where(sq.Eq{"id": filter.Value})
		}
	}

	builder = withPagination(builder, params)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	teams := make(TeamCollection, 0)

	for rows.Next() {
		var t Team
		if err := scanTeam(rows, &t); err != nil {
			return nil, err
		}

		teams = append(teams, t)
	}

	return teams, nil
}

func teamByID(ctx context.Context, tx *sql.Tx, id int64) (Team, error) {
	query, args, err := teamBuilder.
		Where(sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		return Team{}, err
	}

	var t Team
	row := tx.QueryRowContext(ctx, query, args...)
	if err := scanTeam(row, &t); err != nil {
		return Team{}, err
	}

	return t, nil
}

func teamByName(ctx context.Context, tx *sql.Tx, name string) (Team, error) {
	query, args, err := teamBuilder.
		Where(sq.Eq{"name": name}).
		ToSql()

	if err != nil {
		return Team{}, err
	}

	var t Team
	row := tx.QueryRowContext(ctx, query, args...)
	if err := scanTeam(row, &t); err != nil {
		return Team{}, err
	}

	return t, nil
}

func deleteTeam(ctx context.Context, tx *sql.Tx, teamID int64) error {
	query, args, err := sq.
		Delete("team").
		Where(sq.Eq{"id": teamID}).
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

func memberTeams(ctx context.Context, tx *sql.Tx, memberID int64) (TeamCollection, error) {
	query, args, err := teamBuilder.
		Join("members_teams mt ON mt.team_id = team.id").
		Where(sq.Eq{"mt.member_id": memberID}).
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	teams := make(TeamCollection, 0)

	for rows.Next() {
		var t Team
		if err := scanTeam(rows, &t); err != nil {
			return nil, err
		}

		teams = append(teams, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return teams, nil
}

func associateMemberWithTeam(ctx context.Context, tx *sql.Tx, memberID int64, teamID int64) error {
	query, args, err := sq.
		Insert("members_teams").
		Options("OR IGNORE").
		Columns("member_id", "team_id").
		Values(memberID, teamID).
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
