package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/mattismoel/konnekt"
	"github.com/mattismoel/konnekt/internal/password"
)

type userService struct {
	repo *Repository
}

func NewUserService(repo *Repository) *userService {
	return &userService{repo: repo}
}

func (s userService) CreateUser(ctx context.Context, user konnekt.User, password password.Password, passwordConfirm password.Password) (konnekt.User, error) {
	err := user.Validate()
	if err != nil {
		return konnekt.User{}, err
	}

	tx, err := s.repo.db.BeginTx(ctx, nil)
	if err != nil {
		return konnekt.User{}, err
	}

	defer tx.Rollback()

	passwordErrors := password.Validate()
	if passwordErrors != nil {
		return konnekt.User{}, konnekt.Errorf(konnekt.ERRINVALID, fmt.Sprint(passwordErrors))
	}

	if !password.Equals(passwordConfirm) {
		return konnekt.User{}, konnekt.Errorf(konnekt.ERRINVALID, "Passwords do not match")
	}

	passwordHash, err := password.Hash()
	if err != nil {
		return konnekt.User{}, err
	}

	user, err = insertUser(ctx, tx, user, passwordHash)
	if err != nil {
		return konnekt.User{}, err
	}

	if err = tx.Commit(); err != nil {
		return konnekt.User{}, err
	}

	return user, nil
}

func (s userService) DeleteUser(ctx context.Context, id int64) error {
	tx, err := s.repo.db.BeginTx(ctx, nil)
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

func (s userService) FindUserByID(ctx context.Context, id int64) (konnekt.User, error) {
	tx, err := s.repo.db.BeginTx(ctx, nil)
	if err != nil {
		return konnekt.User{}, err
	}

	defer tx.Rollback()

	users, err := findUsers(ctx, tx, konnekt.UserFilter{ID: &id})
	if err != nil {
		return konnekt.User{}, err
	}

	if err = tx.Commit(); err != nil {
		return konnekt.User{}, err
	}

	return users[0], nil
}

func (s userService) FindUsers(ctx context.Context, filter konnekt.UserFilter) ([]konnekt.User, error) {
	tx, err := s.repo.db.BeginTx(ctx, nil)
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

func (s userService) UpdateUser(ctx context.Context, id int64, update konnekt.UpdateUser) (konnekt.User, error) {
	return konnekt.User{}, nil
}

func findUsers(ctx context.Context, tx *sql.Tx, filter konnekt.UserFilter) ([]konnekt.User, error) {
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
	users := []konnekt.User{}

	for rows.Next() {
		var user konnekt.User
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
		return []konnekt.User{}, konnekt.Errorf(konnekt.ERRNOTFOUND, "No users found")
	}

	return users, nil
}

func insertUser(ctx context.Context, tx *sql.Tx, user konnekt.User, passwordHash []byte) (konnekt.User, error) {
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
		passwordHash,
	)

	if err != nil {
		return konnekt.User{}, err
	}

	user.ID, err = res.LastInsertId()
	if err != nil {
		return konnekt.User{}, err
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
