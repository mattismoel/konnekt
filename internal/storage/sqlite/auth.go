package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/mattismoel/konnekt/internal/domain/auth"
)

type Session struct {
	ID        string
	UserID    int64
	ExpiresAt time.Time
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

func (repo AuthRepository) SetSessionExpiry(ctx context.Context, sessionID auth.SessionID, newExpiry time.Time) (auth.Session, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return auth.Session{}, err
	}

	defer tx.Rollback()

	dbSession, err := setSessionExpiry(ctx, tx, string(sessionID), newExpiry)
	if err != nil {
		return auth.Session{}, err
	}

	if err := tx.Commit(); err != nil {
		return auth.Session{}, err
	}

	return dbSession.ToInternal(), nil
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

func setSessionExpiry(ctx context.Context, tx *sql.Tx, sessionID string, newExpiry time.Time) (Session, error) {
	query := `
	UPDATE session
	SET expires_at = @expires_at`

	_, err := tx.ExecContext(ctx, query, sql.Named("expires_at", newExpiry))
	if err != nil {
		return Session{}, err
	}

	session, err := sessionByID(ctx, tx, sessionID)
	if err != nil {
		return Session{}, nil
	}

	return session, nil
}

func (s Session) ToInternal() auth.Session {
	return auth.Session{
		ID:        auth.SessionID(s.ID),
		UserID:    s.UserID,
		ExpiresAt: s.ExpiresAt,
	}
}
