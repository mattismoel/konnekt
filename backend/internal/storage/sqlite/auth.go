package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"time"

	sq "github.com/Masterminds/squirrel"
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

	totalCount, err := count(ctx, tx, "permission")
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

var sessionBuilder = sq.
	Select("id", "member_id", "expires_at").
	From("session")

func scanSession(s Scanner, dst *Session) error {
	err := s.Scan(&dst.ID, &dst.MemberID, &dst.ExpiresAt)
	if err != nil {
		return err
	}

	return nil
}

var permissionBuilder = sq.
	Select("id", "name", "display_name", "description").
	From("permission")

func scanPermission(s Scanner, dst *Permission) error {
	err := s.Scan(&dst.ID, &dst.Name, &dst.DisplayName, &dst.Description)
	if err != nil {
		return err
	}

	return nil
}

func deleteMemberSession(ctx context.Context, tx *sql.Tx, memberID int64) error {
	query, args, err := sq.
		Delete("session").
		Where(sq.Eq{"member_id": memberID}).
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

func insertSession(ctx context.Context, tx *sql.Tx, session Session) error {
	query, args, err := sq.
		Insert("session").
		Columns("id", "member_id", "expires_at").
		Values(session.ID, session.MemberID, session.ExpiresAt).
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

func sessionByID(ctx context.Context, tx *sql.Tx, sessionID string) (Session, error) {
	query, args, err := sessionBuilder.
		Where(sq.Eq{"id": sessionID}).
		ToSql()

	if err != nil {
		return Session{}, err
	}

	var s Session
	row := tx.QueryRowContext(ctx, query, args...)
	if err := scanSession(row, &s); err != nil {
		return Session{}, err
	}

	return s, nil
}

func setSessionExpiry(ctx context.Context, tx *sql.Tx, sessionID string, newExpiry time.Time) error {
	query, args, err := sq.
		Update("session").
		Where(sq.Eq{"id": sessionID}).
		Set("expires_at", newExpiry).
		ToSql()

	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, query, args...)
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
	query, args, err := permissionBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	permissions := make(PermissionCollection, 0)

	for rows.Next() {
		var p Permission
		if err := scanPermission(rows, &p); err != nil {
			return nil, err
		}

		permissions = append(permissions, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}

func teamPermissions(ctx context.Context, tx *sql.Tx, teamID int64) (PermissionCollection, error) {
	query, args, err := sq.
		Select("p.id", "p.name", "p.display_name", "p.description").
		From("permission p").
		Join("teams_permissions tp on tp.permission_id = p.id").
		Where(sq.Eq{"tp.team_id": teamID}).
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	permissions := make(PermissionCollection, 0)

	for rows.Next() {
		var p Permission

		if err := scanPermission(rows, &p); err != nil {
			return nil, err
		}

		permissions = append(permissions, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}

func (pc PermissionCollection) ToInternal() auth.PermissionCollection {
	perms := make(auth.PermissionCollection, 0)

	for _, dbPerm := range pc {
		perms = append(perms, dbPerm.ToInternal())
	}

	return perms
}
