package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/mattismoel/konnekt"
	"github.com/mattismoel/konnekt/internal/password"
)

type userService struct {
	repo *Repository
}

type UserService interface {
	FindUserByID(context.Context, int64) (konnekt.User, error)
	FindUsers(context.Context, konnekt.UserFilter) ([]konnekt.User, error)

	CreateUser(context.Context, konnekt.User, password.Password, password.Password) (int64, error)

	UpdateUser(context.Context, konnekt.UpdateUser) (konnekt.User, error)
	DeleteUser(context.Context, int64) error
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
