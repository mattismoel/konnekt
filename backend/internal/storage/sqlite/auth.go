package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"time"

	"github.com/mattismoel/konnekt/internal/domain/auth"
	"github.com/mattismoel/konnekt/internal/query"
)

type Session struct {
	ID        string
	UserID    int64
	ExpiresAt time.Time
}

type Role struct {
	ID          int64
	Name        string
	DisplayName string
	Description string
}

type Permission struct {
	ID          int64
	Name        string
	Description string
}

var _ auth.Repository = (*AuthRepository)(nil)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) (*AuthRepository, error) {
	return &AuthRepository{
		db: db,
	}, nil
}

func (repo AuthRepository) InsertRole(ctx context.Context, r auth.Role) (int64, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	defer tx.Rollback()

	roleID, err := insertRole(ctx, tx, Role{
		Name:        r.Name,
		DisplayName: r.DisplayName,
		Description: r.Description,
	})

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return roleID, nil
}

func (repo AuthRepository) RoleByID(ctx context.Context, id int64) (auth.Role, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return auth.Role{}, err
	}

	defer tx.Rollback()

	dbRole, err := roleByID(ctx, tx, id)
	if err != nil {
		return auth.Role{}, err
	}

	if err := tx.Commit(); err != nil {
		return auth.Role{}, err
	}

	return dbRole.ToInternal(), nil
}

func (repo AuthRepository) RoleByName(ctx context.Context, name string) (auth.Role, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return auth.Role{}, err
	}

	defer tx.Rollback()

	dbRole, err := roleByName(ctx, tx, name)
	if err != nil {
		return auth.Role{}, err
	}

	if err := tx.Commit(); err != nil {
		return auth.Role{}, err
	}

	return dbRole.ToInternal(), nil
}

func (repo AuthRepository) AddUserRoles(ctx context.Context, userID int64, roleIDs ...int64) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	for _, roleID := range roleIDs {
		err := associateUserWithRole(ctx, tx, userID, roleID)
		if err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (repo AuthRepository) InsertSession(ctx context.Context, session auth.Session) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	dbSession := Session{
		ID:        string(session.ID),
		UserID:    session.UserID,
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

func (repo AuthRepository) DeleteUserSession(ctx context.Context, userID int64) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	if err := deleteUserSession(ctx, tx, userID); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil

}

func (repo AuthRepository) UserRoles(ctx context.Context, userID int64) ([]auth.Role, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	dbRoles, err := userRoles(ctx, tx, userID)
	if err != nil {
		return nil, err
	}

	roles := make([]auth.Role, 0)

	for _, dbRole := range dbRoles {
		roles = append(roles, dbRole.ToInternal())
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return roles, nil
}

func (repo AuthRepository) ListRoles(ctx context.Context, q query.ListQuery) (query.ListResult[auth.Role], error) {
	roles := []auth.Role{}

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return query.ListResult[auth.Role]{}, err
	}

	defer tx.Rollback()

	dbRoles, err := listRoles(ctx, tx, QueryParams{
		Offset:  q.Offset(),
		Limit:   q.Limit,
		OrderBy: q.OrderBy,
		Filters: q.Filters,
	})

	if err != nil {
		return query.ListResult[auth.Role]{}, err
	}

	totalCount, err := roleCount(ctx, tx)
	if err != nil {
		return query.ListResult[auth.Role]{}, err
	}

	if err := tx.Commit(); err != nil {
		return query.ListResult[auth.Role]{}, err
	}

	for _, dbRole := range dbRoles {
		roles = append(roles, dbRole.ToInternal())
	}

	return query.ListResult[auth.Role]{
		Records:    roles,
		Page:       q.Page,
		PerPage:    q.PerPage,
		TotalCount: totalCount,
		PageCount:  q.PageCount(totalCount),
	}, nil
}

func (repo AuthRepository) RolePermissions(ctx context.Context, roleID int64) (auth.PermissionCollection, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	dbPermissions, err := rolePermissions(ctx, tx, roleID)
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

func userRoles(ctx context.Context, tx *sql.Tx, userID int64) ([]Role, error) {
	query := `
	SELECT r.id, r.name, r.display_name, r.description
	FROM role r
	JOIN users_roles ur ON ur.role_id = r.id
	WHERE ur.user_id = @user_id`

	rows, err := tx.QueryContext(ctx, query, sql.Named("user_id", userID))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	roles := make([]Role, 0)

	for rows.Next() {
		var id int64
		var name, displayName, description string

		err := rows.Scan(&id, &name, &displayName, &description)
		if err != nil {
			return nil, err
		}

		roles = append(roles, Role{
			ID:          id,
			Name:        name,
			DisplayName: displayName,
			Description: description,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}

func rolePermissions(ctx context.Context, tx *sql.Tx, roleID int64) ([]Permission, error) {
	query := `
	SELECT p.id, p.name, p.description
	FROM permission p
	JOIN roles_permissions rp on rp.permission_id = p.id
	WHERE rp.role_id = @role_id`

	rows, err := tx.QueryContext(ctx, query, sql.Named("role_id", roleID))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	permissions := make([]Permission, 0)

	for rows.Next() {
		var id int64
		var name, description string

		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}

		permissions = append(permissions, Permission{
			ID:          id,
			Name:        name,
			Description: description,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}

func deleteUserSession(ctx context.Context, tx *sql.Tx, userID int64) error {
	query := `DELETE FROM session WHERE user_id = @user_id`
	_, err := tx.ExecContext(ctx, query, sql.Named("user_id", userID))
	if err != nil {
		return err
	}

	return nil
}

func insertSession(ctx context.Context, tx *sql.Tx, session Session) error {
	query := `
	INSERT INTO session (id, user_id, expires_at)
	VALUES (@id, @user_id, @expires_at)`

	_, err := tx.ExecContext(ctx, query,
		sql.Named("id", session.ID),
		sql.Named("user_id", session.UserID),
		sql.Named("expires_at", session.ExpiresAt),
	)
	if err != nil {
		return err
	}

	return nil
}

func sessionByID(ctx context.Context, tx *sql.Tx, sessionID string) (Session, error) {
	query := `
	SELECT user_id, expires_at 
	FROM session 
	WHERE id = @session_id`

	var userID int64
	var expiresAt time.Time

	err := tx.QueryRowContext(ctx, query,
		sql.Named("session_id", sessionID),
	).Scan(&userID, &expiresAt)

	if err != nil {
		return Session{}, err
	}

	return Session{
		ID:        sessionID,
		UserID:    userID,
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

func listRoles(ctx context.Context, tx *sql.Tx, params QueryParams) ([]Role, error) {
	q, err := NewQuery(`
	SELECT DISTINCT id, name, description, display_name
	FROM role`)

	if err != nil {
		return nil, err
	}

	queryStr, args := q.Build()

	rows, err := tx.QueryContext(ctx, queryStr, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	roles := make([]Role, 0)

	for rows.Next() {
		var id int64
		var name, description, displayName string

		err := rows.Scan(&id, &name, &description, &displayName)
		if err != nil {
			return nil, err
		}

		roles = append(roles, Role{
			ID:          id,
			Name:        name,
			Description: description,
			DisplayName: displayName,
		})
	}

	return roles, nil
}

func (s Session) ToInternal() auth.Session {
	return auth.Session{
		ID:        auth.SessionID(s.ID),
		UserID:    s.UserID,
		ExpiresAt: s.ExpiresAt,
	}
}

func (p Permission) ToInternal() auth.Permission {
	return auth.Permission{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
	}
}

func (r Role) ToInternal() auth.Role {
	return auth.Role{
		ID:          r.ID,
		Name:        r.Name,
		DisplayName: r.DisplayName,
		Description: r.Description,
	}
}

func roleCount(ctx context.Context, tx *sql.Tx) (int, error) {
	q, err := NewQuery("SELECT COUNT(*) FROM role")
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

func insertRole(ctx context.Context, tx *sql.Tx, r Role) (int64, error) {
	query := `
	INSERT INTO role (name, display_name, description) 
	VALUES (@name, @display_name, @description)`

	res, err := tx.ExecContext(ctx, query,
		sql.Named("name", r.Name),
		sql.Named("display_name", r.DisplayName),
		sql.Named("description", r.Description),
	)

	if err != nil {
		return 0, err
	}

	roleID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return roleID, nil
}

func associateUserWithRole(ctx context.Context, tx *sql.Tx, userID int64, roleID int64) error {
	query := `INSERT INTO users_roles (user_id, role_id) VALUES (@user_id, @role_id)`

	_, err := tx.ExecContext(ctx, query, sql.Named("user_id", userID), sql.Named("role_id", roleID))
	if err != nil {
		return err
	}

	return nil
}

func roleByID(ctx context.Context, tx *sql.Tx, id int64) (Role, error) {
	q, err := NewQuery("SELECT name, display_name, description FROM role")
	if err != nil {
		return Role{}, err
	}

	if err := q.AddFilter("id", query.Equal, strconv.Itoa(int(id))); err != nil {
		return Role{}, err
	}

	queryStr, args := q.Build()

	var name, displayName, description string

	err = tx.QueryRowContext(ctx, queryStr, args...).Scan(
		&name, &displayName, &description,
	)

	if err != nil {
		return Role{}, err
	}

	return Role{
		ID:          id,
		Name:        name,
		DisplayName: displayName,
		Description: description,
	}, nil
}

func roleByName(ctx context.Context, tx *sql.Tx, name string) (Role, error) {
	q, err := NewQuery("SELECT id, display_name, description FROM role")
	if err != nil {
		return Role{}, err
	}

	if err := q.AddFilter("name", query.Equal, name); err != nil {
		return Role{}, err
	}

	queryStr, args := q.Build()

	var id int64
	var displayName, description string

	err = tx.QueryRowContext(ctx, queryStr, args...).Scan(
		&id, &displayName, &description,
	)

	if err != nil {
		return Role{}, err
	}

	return Role{
		ID:          id,
		Name:        name,
		DisplayName: displayName,
		Description: description,
	}, nil
}
