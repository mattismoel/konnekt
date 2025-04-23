package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

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

	fmt.Printf("Q FILTERS: %+v\n", q.Filters)

	dbTeams, err := listTeams(ctx, tx, QueryParams{
		Offset:  q.Offset(),
		Limit:   q.Limit,
		OrderBy: q.OrderBy,
		Filters: q.Filters,
	})

	if err != nil {
		return query.ListResult[team.Team]{}, err
	}

	totalCount, err := teamCount(ctx, tx)
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

func insertTeam(ctx context.Context, tx *sql.Tx, r Team) (int64, error) {
	query := `
	INSERT INTO team (name, display_name, description) 
	VALUES (@name, @display_name, @description)`

	res, err := tx.ExecContext(ctx, query,
		sql.Named("name", r.Name),
		sql.Named("display_name", r.DisplayName),
		sql.Named("description", r.Description),
	)

	if err != nil {
		return 0, err
	}

	teamID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return teamID, nil
}

func listTeams(ctx context.Context, tx *sql.Tx, params QueryParams) (TeamCollection, error) {
	q, err := NewQuery(`
	SELECT DISTINCT id, name, description, display_name
	FROM team`)

	if err != nil {
		return nil, err
	}


	if filters, ok := params.Filters["id"]; ok {
		for _, filter := range filters {
			if err := q.AddFilter("id", filter.Cmp, filter.Value); err != nil {
				return nil, err
			}
		}
	}


	queryStr, args := q.Build()


	rows, err := tx.QueryContext(ctx, queryStr, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	teams := make(TeamCollection, 0)

	for rows.Next() {
		var id int64
		var name, description, displayName string

		err := rows.Scan(&id, &name, &description, &displayName)
		if err != nil {
			return nil, err
		}

		teams = append(teams, Team{
			ID:          id,
			Name:        name,
			Description: description,
			DisplayName: displayName,
		})
	}

	return teams, nil
}

func teamByID(ctx context.Context, tx *sql.Tx, id int64) (Team, error) {
	q, err := NewQuery("SELECT name, display_name, description FROM team")
	if err != nil {
		return Team{}, err
	}

	if err := q.AddFilter("id", query.Equal, strconv.Itoa(int(id))); err != nil {
		return Team{}, err
	}

	queryStr, args := q.Build()

	var name, displayName, description string

	err = tx.QueryRowContext(ctx, queryStr, args...).Scan(
		&name, &displayName, &description,
	)

	if err != nil {
		return Team{}, err
	}

	return Team{
		ID:          id,
		Name:        name,
		DisplayName: displayName,
		Description: description,
	}, nil
}

func teamByName(ctx context.Context, tx *sql.Tx, name string) (Team, error) {
	q, err := NewQuery("SELECT id, display_name, description FROM team")
	if err != nil {
		return Team{}, err
	}

	if err := q.AddFilter("name", query.Equal, name); err != nil {
		return Team{}, err
	}

	queryStr, args := q.Build()

	var id int64
	var displayName, description string

	err = tx.QueryRowContext(ctx, queryStr, args...).Scan(
		&id, &displayName, &description,
	)

	if err != nil {
		return Team{}, err
	}

	return Team{
		ID:          id,
		Name:        name,
		DisplayName: displayName,
		Description: description,
	}, nil
}

func deleteTeam(ctx context.Context, tx *sql.Tx, teamID int64) error {
	query := "DELETE FROM team WHERE id = @team_id"

	res, err := tx.ExecContext(ctx, query, sql.Named("team_id", teamID))
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
	query := `
	SELECT t.id, t.name, t.display_name, t.description
	FROM team t
	JOIN members_teams mt ON mt.team_id = t.id
	WHERE mt.member_id = @member_id`

	rows, err := tx.QueryContext(ctx, query, sql.Named("member_id", memberID))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	teams := make(TeamCollection, 0)

	for rows.Next() {
		var id int64
		var name, displayName, description string

		err := rows.Scan(&id, &name, &displayName, &description)
		if err != nil {
			return nil, err
		}

		teams = append(teams, Team{
			ID:          id,
			Name:        name,
			DisplayName: displayName,
			Description: description,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return teams, nil
}

func associateMemberWithTeam(ctx context.Context, tx *sql.Tx, memberID int64, teamID int64) error {
	query := `INSERT OR IGNORE INTO members_teams (member_id, team_id) VALUES (@member_id, @team_id)`

	_, err := tx.ExecContext(ctx, query, sql.Named("member_id", memberID), sql.Named("team_id", teamID))
	if err != nil {
		return err
	}

	return nil
}

func teamCount(ctx context.Context, tx *sql.Tx) (int, error) {
	q, err := NewQuery("SELECT COUNT(*) FROM team")
	if err != nil {
		return 0, err
	}

	queryStr, args := q.Build()

	var count int

	err = tx.QueryRowContext(ctx, queryStr, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
