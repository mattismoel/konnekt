package sqlite

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mattismoel/konnekt/internal/domain/auth"
	"github.com/mattismoel/konnekt/internal/domain/user"
)

var _ user.Repository = (*UserRepository)(nil)

type User struct {
	ID           int64
	Email        string
	FirstName    string
	LastName     string
	PasswordHash []byte
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) (*UserRepository, error) {
	return &UserRepository{
		db: db,
	}, nil
}

func (repo UserRepository) Insert(ctx context.Context, u user.User) (int64, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	defer tx.Rollback()

	userID, err := insertUser(ctx, tx, User{
		ID:           u.ID,
		Email:        u.Email,
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		PasswordHash: u.PasswordHash,
	})

	if err != nil {
		switch {
		case errors.Is(err, ErrUserAlreadyExists):
			return 0, user.ErrAlreadyExists
		}
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return userID, nil
}

func (repo UserRepository) ByID(ctx context.Context, userID int64) (user.User, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return user.User{}, err
	}

	defer tx.Rollback()

	usr, err := userByID(ctx, tx, userID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return user.User{}, ErrNotFound
		default:
			return user.User{}, err
		}
	}

	userRoles, err := userRoles(ctx, tx, userID)
	if err != nil {
		return user.User{}, err
	}

	userPerms, err := userPermsissions(ctx, tx, userID)
	if err != nil {
		return user.User{}, err
	}

	if err := tx.Commit(); err != nil {
		return user.User{}, err
	}

	return usr.ToInternal(userRoles.ToInternal(), userPerms.ToInternal()), nil
}

}

func (repo UserRepository) ByEmail(ctx context.Context, email string) (user.User, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return user.User{}, err
	}

	defer tx.Rollback()

	usr, err := userByEmail(ctx, tx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user.User{}, user.ErrNotFound
		}

		return user.User{}, err
	}

	userRoles, err := userRoles(ctx, tx, usr.ID)
	if err != nil {
		return user.User{}, err
	}

	userPerms, err := userPermsissions(ctx, tx, usr.ID)
	if err != nil {
		return user.User{}, err
	}

	if err := tx.Commit(); err != nil {
		return user.User{}, err
	}

	return usr.ToInternal(userRoles.ToInternal(), userPerms.ToInternal()), nil
}

func (repo UserRepository) PasswordHash(ctx context.Context, userID int64) (user.PasswordHash, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	ph, err := userPasswordHash(ctx, tx, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, user.ErrNotFound
		}

		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return ph, nil
}

func insertUser(ctx context.Context, tx *sql.Tx, u User) (int64, error) {
	query := `
	INSERT OR IGNORE INTO user (email, first_name, last_name, password_hash) 
	VALUES (@email, @first_name, @last_name, @password_hash)`

	res, err := tx.ExecContext(ctx, query,
		sql.Named("email", u.Email),
		sql.Named("first_name", u.FirstName),
		sql.Named("last_name", u.LastName),
		sql.Named("password_hash", u.PasswordHash),
	)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	if rowsAffected <= 0 {
		return 0, ErrUserAlreadyExists
	}

	userID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func userByEmail(ctx context.Context, tx *sql.Tx, email string) (User, error) {
	query := `
  SELECT id, first_name, last_name, password_hash FROM user
  WHERE email = @email`

	var id int64
	var firstName, lastName string
	var passwordHash []byte

	err := tx.QueryRowContext(ctx, query, sql.Named("email", email)).Scan(
		&id, &firstName, &lastName, &passwordHash,
	)

	if err != nil {
		return User{}, err
	}

	return User{
		ID:           id,
		Email:        email,
		FirstName:    firstName,
		LastName:     lastName,
		PasswordHash: passwordHash,
	}, nil
}

func userByID(ctx context.Context, tx *sql.Tx, userID int64) (User, error) {
	query := `
  SELECT email, first_name, last_name, password_hash FROM user
  WHERE id = @user_id`

	var email, firstName, lastName string
	var passwordHash []byte

	err := tx.QueryRowContext(ctx, query, sql.Named("user_id", userID)).Scan(
		&email, &firstName, &lastName, &passwordHash,
	)

	if err != nil {
		return User{}, err
	}

	return User{
		ID:           userID,
		Email:        email,
		FirstName:    firstName,
		LastName:     lastName,
		PasswordHash: passwordHash,
	}, nil
}

func userPasswordHash(ctx context.Context, tx *sql.Tx, userID int64) ([]byte, error) {
	query := `SELECT password_hash FROM user WHERE id = @id`

	var passwordHash []byte

	err := tx.QueryRowContext(ctx, query, sql.Named("id", userID)).Scan(&passwordHash)
	if err != nil {
		return nil, err
	}

	return passwordHash, nil
}

func (u User) ToInternal(roles []auth.Role, perms auth.PermissionCollection) user.User {
	return user.User{
		ID:           u.ID,
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,

		Roles:       roles,
		Permissions: perms,
	}
}
