package psql

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mattismoel/konnekt/backend/auth"
)

var _ auth.SessionRepo = SessionRepo{}

type Session struct {
	ID         int64
	SessionID  string
	SecretHash []byte
	CreatedAt  time.Time
}

type SessionRepo struct {
	Pool *pgxpool.Pool
}

// DeleteSession implements auth.SessionRepo.
func (s SessionRepo) DeleteSession(ctx context.Context, sessionID auth.SessionID) error {
	err := pgx.BeginFunc(ctx, s.Pool, func(tx pgx.Tx) error {
		if err := deleteSession(ctx, tx, string(sessionID)); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

// InsertSession implements auth.SessionRepo.
func (s SessionRepo) InsertSession(ctx context.Context, session auth.Session) (int64, error) {
	var id int64
	err := pgx.BeginFunc(ctx, s.Pool, func(tx pgx.Tx) error {
		insertedID, err := insertSession(ctx, tx, Session{
			SessionID:  string(session.ID),
			SecretHash: session.SecretHash,
			CreatedAt:  session.CreatedAt,
		})

		if err != nil {
			return fmt.Errorf("Could not insert session: %v", err)
		}

		id = insertedID
		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s SessionRepo) GetSession(ctx context.Context, sessionID auth.SessionID) (auth.Session, error) {
	return auth.Session{}, nil
}

func insertSession(ctx context.Context, tx pgx.Tx, s Session) (int64, error) {
	query, args, err := psql.
		Insert("session").
		Columns("session_id", "secret_hash", "created_at").
		Values(s.SessionID, s.SecretHash, s.CreatedAt).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return 0, NewQueryBuildError("insert session", err)
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return 0, fmt.Errorf("Could not insert session: %v", err)
	}

	id, err := pgx.CollectExactlyOneRow(rows, pgx.RowTo[int64])
	if err != nil {
		return 0, fmt.Errorf("Could not get inserted ID: %v", err)
	}

	return id, nil
}

func deleteSession(ctx context.Context, tx pgx.Tx, sessionID string) error {
	query, args, err := psql.
		Delete("session").
		Where(sq.Eq{"session_id": sessionID}).
		ToSql()

	if err != nil {
		return NewQueryBuildError("delete session", err)
	}

	if _, err := tx.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("Could not delete session: %v", err)
	}

	return nil
}
