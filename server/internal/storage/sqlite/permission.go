package sqlite

import (
	"context"
	"database/sql"

	"github.com/mattismoel/konnekt/internal/storage"
)

type permissionRepository struct {
	store *Store
}

func NewPermissionRepository(store *Store) *permissionRepository {
	return &permissionRepository{store: store}
}

func (r permissionRepository) HasPermission(ctx context.Context, userID int64, permName string) error {
	tx, err := r.store.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	roles, err := findUserRoles(ctx, tx, userID)
	if err != nil {
		return err
	}

	permissions := []storage.Permission{}

	for _, role := range roles {
		rolePerms, err := findRolePermissions(ctx, tx, role.ID)
		if err != nil {
			return err
		}

		permissions = append(permissions, rolePerms...)
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	for _, perm := range permissions {
		if perm.Name == permName {
			return nil
		}
	}

	return storage.ErrNotFound
}

func findUserRoles(ctx context.Context, tx *sql.Tx, userID int64) ([]storage.Role, error) {
	query := `
	SELECT
		id,
		name,
		description
	FROM role
	INNER join user_role ON user_role.role_id = role.id
	WHERE user_role.user_id = @user_id`

	rows, err := tx.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	roles := []storage.Role{}

	for rows.Next() {
		var role storage.Role

		err := rows.Scan(&role.ID, &role.Name, &role.Description)
		if err != nil {
			return nil, err
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}

func findRolePermissions(ctx context.Context, tx *sql.Tx, roleID int64) ([]storage.Permission, error) {
	query := `
	SELECT
		id,
		name,
		description
	FROM permission
	INNER JOIN role_permission ON role_permission.permission_id = permission.id
	WHERE role_id = @role_id`

	rows, err := tx.QueryContext(ctx, query, sql.Named("role_id", roleID))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	permissions := []storage.Permission{}

	for rows.Next() {
		var permission storage.Permission

		err := rows.Scan(&permission.ID, &permission.Name, &permission.Description)
		if err != nil {
			return nil, err
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}
