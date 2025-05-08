package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/mattismoel/konnekt/internal/domain/auth"
	"github.com/mattismoel/konnekt/internal/query"
)

type Session struct {
	ID        string
	MemberID  int64
	ExpiresAt time.Time
}

type Permission struct {
	ID          int64
	Name        string
	DisplayName string
	Description string
}

type PermissionCollection []Permission

var _ auth.Repository = (*AuthRepository)(nil)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) (*AuthRepository, error) {
	return &AuthRepository{
		db: db,
	}, nil
}

func (repo AuthRepository) InsertSession(ctx context.Context, session auth.Session) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	dbSession := Session{
		ID:        string(session.ID),
		MemberID:  session.MemberID,
		ExpiresAt: session.ExpiresAt,
	}

	if err := insertSession(ctx, tx, dbSession); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (repo AuthRepository) Session(ctx context.Context, sessionID auth.SessionID) (auth.Session, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	defer tx.Rollback()

	dbSession, err := sessionByID(ctx, tx, string(sessionID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return auth.Session{}, auth.ErrNoSession
		}
	}

	if err := tx.Commit(); err != nil {
		return auth.Session{}, err
	}

	return dbSession.ToInternal(), nil
}

func (repo AuthRepository) SetSessionExpiry(ctx context.Context, sessionID auth.SessionID, newExpiry time.Time) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	err = setSessionExpiry(ctx, tx, string(sessionID), newExpiry)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (repo AuthRepository) DeleteMemberSession(ctx context.Context, memberID int64) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	if err := deleteMemberSession(ctx, tx, memberID); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil

}

func (repo AuthRepository) ListPermissions(ctx context.Context, q query.ListQuery) (query.ListResult[auth.Permission], error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return query.ListResult[auth.Permission]{}, err
	}

	defer tx.Rollback()

	dbPermissions, err := listPermissions(ctx, tx, QueryParams{
		Offset:  q.Offset(),
		Limit:   q.Limit,
		OrderBy: q.OrderBy,
		Filters: q.Filters,
	})

	if err != nil {
		return query.ListResult[auth.Permission]{}, err
	}

	permissions := make([]auth.Permission, 0)
	for _, dbPerm := range dbPermissions {
		permissions = append(permissions, dbPerm.ToInternal())
	}

	totalCount, err := permissionCount(ctx, tx)
	if err != nil {
		return query.ListResult[auth.Permission]{}, err
	}

	if err := tx.Commit(); err != nil {
		return query.ListResult[auth.Permission]{}, err
	}

	return query.ListResult[auth.Permission]{
		Page:       q.Page,
		PerPage:    q.PerPage,
		TotalCount: totalCount,
		PageCount:  q.PageCount(totalCount),
		Records:    permissions,
	}, nil
}

func (repo AuthRepository) TeamPermissions(ctx context.Context, teamID int64) (auth.PermissionCollection, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	dbPermissions, err := teamPermissions(ctx, tx, teamID)
	if err != nil {
		return nil, err
	}

	collection := make(auth.PermissionCollection, 0)

	for _, dbPerm := range dbPermissions {
		perm := dbPerm.ToInternal()
		collection = append(collection, perm)
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return collection, nil
}

func deleteMemberSession(ctx context.Context, tx *sql.Tx, memberID int64) error {
	query := `DELETE FROM session WHERE member_id = @member_id`
	_, err := tx.ExecContext(ctx, query, sql.Named("member_id", memberID))
	if err != nil {
		return err
	}

	return nil
}

func insertSession(ctx context.Context, tx *sql.Tx, session Session) error {
	query := `
	INSERT INTO session (id, member_id, expires_at)
	VALUES (@id, @member_id, @expires_at)`

	_, err := tx.ExecContext(ctx, query,
		sql.Named("id", session.ID),
		sql.Named("member_id", session.MemberID),
		sql.Named("expires_at", session.ExpiresAt),
	)
	if err != nil {
		return err
	}

	return nil
}

func sessionByID(ctx context.Context, tx *sql.Tx, sessionID string) (Session, error) {
	query := `
	SELECT member_id, expires_at 
	FROM session 
	WHERE id = @session_id`

	var memberID int64
	var expiresAt time.Time

	err := tx.QueryRowContext(ctx, query,
		sql.Named("session_id", sessionID),
	).Scan(&memberID, &expiresAt)

	if err != nil {
		return Session{}, err
	}

	return Session{
		ID:        sessionID,
		MemberID:  memberID,
		ExpiresAt: expiresAt,
	}, nil
}

func setSessionExpiry(ctx context.Context, tx *sql.Tx, sessionID string, newExpiry time.Time) error {
	query := `
	UPDATE session
	SET expires_at = @expires_at`

	_, err := tx.ExecContext(ctx, query, sql.Named("expires_at", newExpiry))
	if err != nil {
		return err
	}

	_, err = sessionByID(ctx, tx, sessionID)
	if err != nil {
		return nil
	}

	return nil
}

func (s Session) ToInternal() auth.Session {
	return auth.Session{
		ID:        auth.SessionID(s.ID),
		MemberID:  s.MemberID,
		ExpiresAt: s.ExpiresAt,
	}
}

func (p Permission) ToInternal() auth.Permission {
	return auth.Permission{
		ID:          p.ID,
		Name:        p.Name,
		DisplayName: p.DisplayName,
		Description: p.Description,
	}
}

func listPermissions(ctx context.Context, tx *sql.Tx, params QueryParams) (PermissionCollection, error) {
	q, err := NewQuery(`SELECT id, name, display_name, description FROM permission`)
	if err != nil {
		return nil, err
	}

	queryStr, args := q.Build()

	rows, err := tx.QueryContext(ctx, queryStr, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	permissions := make(PermissionCollection, 0)

	for rows.Next() {
		var id int64
		var name, displayName, description string

		err := rows.Scan(&id, &name, &displayName, &description)
		if err != nil {
			return nil, err
		}

		permissions = append(permissions, Permission{
			ID:          id,
			DisplayName: displayName,
			Name:        name,
			Description: description,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}

func teamPermissions(ctx context.Context, tx *sql.Tx, teamID int64) (PermissionCollection, error) {
	query := `
	SELECT p.id, p.name, p.display_name, p.description
	FROM permission p
	JOIN teams_permissions tp on tp.permission_id = p.id
	WHERE tp.team_id = @team_id`

	rows, err := tx.QueryContext(ctx, query, sql.Named("team_id", teamID))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	permissions := make(PermissionCollection, 0)

	for rows.Next() {
		var id int64
		var name, displayName, description string

		if err := rows.Scan(&id, &name, &displayName, &description); err != nil {
			return nil, err
		}

		permissions = append(permissions, Permission{
			ID:          id,
			Name:        name,
			DisplayName: displayName,
			Description: description,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}

func permissionCount(ctx context.Context, tx *sql.Tx) (int, error) {
	query := "SELECT COUNT(*) FROM permission"

	var count int
	err := tx.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (pc PermissionCollection) ToInternal() auth.PermissionCollection {
	perms := make(auth.PermissionCollection, 0)

	for _, dbPerm := range pc {
		perms = append(perms, dbPerm.ToInternal())
	}

	return perms
}
