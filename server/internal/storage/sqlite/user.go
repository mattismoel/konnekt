package sqlite

import (
	"context"
	"database/sql"
	"strings"

	"github.com/mattismoel/konnekt/internal/service"
	"github.com/mattismoel/konnekt/internal/storage"
)

type userRepository struct {
	store *Store
}

func NewUserRepository(store *Store) *userRepository {
	return &userRepository{store: store}
}

func (s userRepository) InsertUser(ctx context.Context, user storage.User) (storage.User, error) {
	tx, err := s.store.BeginTx(ctx, nil)
	if err != nil {
		return storage.User{}, err
	}

	defer tx.Rollback()

	user, err = insertUser(ctx, tx, user)
	if err != nil {
		return storage.User{}, err
	}

	if err = tx.Commit(); err != nil {
		return storage.User{}, err
	}

	return user, nil
}

func (s userRepository) DeleteUser(ctx context.Context, id int64) error {
	tx, err := s.store.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	if err = deleteUserByID(ctx, tx, id); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (s userRepository) FindUserByID(ctx context.Context, id int64) (storage.User, error) {
	tx, err := s.store.BeginTx(ctx, nil)
	if err != nil {
		return storage.User{}, err
	}

	defer tx.Rollback()

	user, err := findUserByID(ctx, tx, id)
	if err != nil {
		return storage.User{}, err
	}

	if err = tx.Commit(); err != nil {
		return storage.User{}, err
	}

	return user, nil
}

func (s userRepository) FindUserByEmail(ctx context.Context, email string) (storage.User, error) {
	tx, err := s.store.BeginTx(ctx, nil)
	if err != nil {
		return storage.User{}, err
	}

	defer tx.Rollback()

	user, err := findUserByEmail(ctx, tx, email)
	if err != nil {
		return storage.User{}, err
	}

	if err = tx.Commit(); err != nil {
		return storage.User{}, err
	}

	return user, nil

}

func (s userRepository) FindUsers(ctx context.Context, filter service.UserFilter) ([]storage.User, error) {
	tx, err := s.store.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	users, err := findUsers(ctx, tx, filter)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return users, nil
}

func (s userRepository) UpdateUser(ctx context.Context, id int64, update storage.User) (storage.User, error) {
	tx, err := s.store.BeginTx(ctx, nil)
	if err != nil {
		return storage.User{}, err
	}

	defer tx.Rollback()

	user, err := updateUser(ctx, tx, id, update)
	if err != nil {
		return storage.User{}, err
	}

	if err = tx.Commit(); err != nil {
		return storage.User{}, err
	}

	return user, nil
}

func (r userRepository) FindUserPasswordHash(ctx context.Context, userID int64) ([]byte, error) {
	tx, err := r.store.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	hash, err := findUserPasswordHash(ctx, tx, userID)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return hash, nil

}

func findUserByID(ctx context.Context, tx *sql.Tx, id int64) (storage.User, error) {
	users, err := findUsers(ctx, tx, service.UserFilter{ID: &id})
	if err != nil {
		return storage.User{}, err
	}

	return users[0], nil
}

func findUsers(ctx context.Context, tx *sql.Tx, filter service.UserFilter) ([]storage.User, error) {
	where, args := []string{"\n1 = 1\n"}, []any{}

	if v := filter.Email; v != nil {
		where, args = append(where, "email = ?\n"), append(args, *v)
	}

	if v := filter.ID; v != nil {
		where, args = append(where, "id = ?\n"), append(args, *v)
	}

	if v := filter.FirstName; v != nil {
		where, args = append(where, "first_name = ?\n"), append(args, *v)
	}

	if v := filter.LastName; v != nil {
		where, args = append(where, "last_name = ?\n"), append(args, *v)
	}

	query := `
	SELECT
		id,
		email,
		first_name,
		last_name
	FROM user
	WHERE`

	query += strings.Join(where, " AND ")
	query += formatLimitOffset(filter.Limit, filter.Offset)

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	users := []storage.User{}

	for rows.Next() {
		var user storage.User
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.FirstName,
			&user.LastName,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if len(users) == 0 || users == nil {
		return []storage.User{}, service.Errorf(service.ERRNOTFOUND, "No users found")
	}

	return users, nil
}

func insertUser(ctx context.Context, tx *sql.Tx, user storage.User) (storage.User, error) {
	query := `
	INSERT INTO user (
		first_name,
		last_name,
		email,
		password_hash
	) 
	VALUES (?, ?, ?, ?)`

	res, err := tx.ExecContext(ctx, query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.PasswordHash,
	)

	if err != nil {
		return storage.User{}, err
	}

	user.ID, err = res.LastInsertId()
	if err != nil {
		return storage.User{}, err
	}

	return user, nil
}

func deleteUserByID(ctx context.Context, tx *sql.Tx, id int64) error {
	query := "DELETE FROM user WHERE id = ?"

	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func updateUser(ctx context.Context, tx *sql.Tx, id int64, update storage.User) (storage.User, error) {
	query := `
	UPDATE user
	SET
		email = CASE
			WHEN @email = '' THEN user.email
			ELSE @email
		END,
		first_name = CASE
			WHEN @first_name = '' THEN user.first_name
			ELSE @first_name
		END,
		last_name = CASE
			WHEN @last_name = '' THEN user.last_name
			ELSE @last_name
		END
	WHERE id = @id
	RETURNING
		id,
		email,
		first_name,
		last_name
	`

	err := tx.QueryRowContext(ctx, query,
		sql.Named("email", update.Email),
		sql.Named("first_name", update.FirstName),
		sql.Named("last_name", update.LastName),
		sql.Named("id", id),
	).Scan(
		&update.ID,
		&update.Email,
		&update.FirstName,
		&update.LastName,
	)

	if err != nil {
		return storage.User{}, err
	}

	return update, nil
}

func findUserByEmail(ctx context.Context, tx *sql.Tx, email string) (storage.User, error) {
	users, err := findUsers(ctx, tx, service.UserFilter{Email: &email})
	if err != nil {
		return storage.User{}, err
	}

	return users[0], nil
}

func findUserPasswordHash(ctx context.Context, tx *sql.Tx, userID int64) ([]byte, error) {
	hash := []byte{}

	query := "SELECT password_hash FROM user WHERE id = @user_id"
	err := tx.QueryRowContext(ctx, query,
		sql.Named("user_id", userID),
	).Scan(&hash)

	if err != nil {
		return nil, err
	}

	return hash, nil
}
