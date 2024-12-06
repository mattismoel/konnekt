package sqlite

import (
	"context"
	"database/sql"
	"time"

	"github.com/mattismoel/konnekt/internal/storage"
)

type sessionRepo struct {
	store *Store
}

func NewSessionRepository(store *Store) *sessionRepo {
	return &sessionRepo{store: store}
}

func (r sessionRepo) InsertSession(ctx context.Context, sessionID string, userID int64, expiresAt time.Time) (storage.Session, error) {
	tx, err := r.store.BeginTx(ctx, nil)
	if err != nil {
		return storage.Session{}, err
	}

	defer tx.Rollback()

	session, err := insertSession(ctx, tx, sessionID, userID, expiresAt)
	if err != nil {
		return storage.Session{}, err
	}

	if err = tx.Commit(); err != nil {
		return storage.Session{}, err
	}

	return session, nil
}

func (r sessionRepo) FindSession(ctx context.Context, sessionID string) (storage.Session, storage.User, error) {
	tx, err := r.store.BeginTx(ctx, nil)
	if err != nil {
		return storage.Session{}, storage.User{}, err
	}

	defer tx.Rollback()

	session, user, err := findSession(ctx, tx, sessionID)
	if err != nil {
		return storage.Session{}, storage.User{}, err
	}

	if err = tx.Commit(); err != nil {
		return storage.Session{}, storage.User{}, err
	}

	return session, user, nil
}

func (r sessionRepo) UpdateSessionExpiry(ctx context.Context, sessionID string, expiry time.Time) (storage.Session, error) {
	tx, err := r.store.BeginTx(ctx, nil)
	if err != nil {
		return storage.Session{}, err
	}

	defer tx.Rollback()

	session, err := updateSessionExpiry(ctx, tx, sessionID, expiry)
	if err != nil {
		return storage.Session{}, err
	}

	if err = tx.Commit(); err != nil {
		return storage.Session{}, err
	}

	return session, nil
}

func (r sessionRepo) DeleteSession(ctx context.Context, sessionID string) error {
	tx, err := r.store.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	err = deleteSession(ctx, tx, sessionID)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func insertSession(ctx context.Context, tx *sql.Tx, sessionID string, userID int64, expiresAt time.Time) (storage.Session, error) {
	query := `
	INSERT INTO session (id, user_id, expires_at)
	VALUES (@id, @user_id, @expires_at)
	RETURNING id, user_id, expires_at`

	var session storage.Session
	var expiresAtUnix int64

	err := tx.QueryRowContext(ctx, query,
		sql.Named("id", sessionID),
		sql.Named("user_id", userID),
		sql.Named("expires_at", expiresAt.UnixMilli()),
	).Scan(
		&session.ID,
		&session.UserID,
		&expiresAtUnix,
	)

	session.ExpiresAt = time.UnixMilli(expiresAtUnix)

	if err != nil {
		return storage.Session{}, err
	}

	return session, nil
}

func findSession(ctx context.Context, tx *sql.Tx, sessionID string) (storage.Session, storage.User, error) {
	query := `
	SELECT 
		session.id,
		session.user_id,
		session.expires_at,
	FROM session
	WHERE id = @session_id`

	var session storage.Session
	var expiresAtUnix int64

	err := tx.QueryRowContext(ctx, query,
		sql.Named("session_id", sessionID),
	).Scan(
		&session.ID,
		&session.UserID,
		&expiresAtUnix,
	)

	if err != nil {
		return storage.Session{}, storage.User{}, err
	}

	user, err := findUserByID(ctx, tx, session.UserID)
	if err != nil {
		return storage.Session{}, storage.User{}, err
	}

	return session, user, nil
}

func deleteSession(ctx context.Context, tx *sql.Tx, sessionID string) error {
	query := "DELETE FROM session WHERE id = @session_id"

	_, err := tx.ExecContext(ctx, query, sql.Named("session_id", sessionID))
	if err != nil {
		return err
	}

	return nil
}

func updateSessionExpiry(ctx context.Context, tx *sql.Tx, sessionID string, expiry time.Time) (storage.Session, error) {
	query := `
	UPDATE session
	SET
		expires_at = @expires_at
	WHERE id = @session_id
	RETURNING id, user_id, expires_at`

	var session storage.Session

	var expiresAtUnix int64

	err := tx.QueryRowContext(ctx, query,
		sql.Named("expires_at", expiry),
		sql.Named("session_id", sessionID),
	).Scan(
		&session.ID,
		&session.UserID,
		&expiresAtUnix,
	)

	if err != nil {
		return storage.Session{}, err
	}

	session.ExpiresAt = time.UnixMilli(expiresAtUnix)

	return session, nil
}
