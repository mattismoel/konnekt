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
	MemberID   int64
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
			MemberID:   session.MemberID,
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
	var session auth.Session
	err := pgx.BeginFunc(ctx, s.Pool, func(tx pgx.Tx) error {
		dbSession, err := sessionBySessionID(ctx, tx, string(sessionID))
		if err != nil {
			return fmt.Errorf("Could not get session: %v", err)
		}

		session, err = dbSession.Assemble(ctx, tx)
		if err != nil {
			return fmt.Errorf("Could not assemble DB session: %v", err)
		}

		return nil
	})

	if err != nil {
		return auth.Session{}, err
	}

	return session, nil
}

func sessionBySessionID(ctx context.Context, tx pgx.Tx, sessionID string) (Session, error) {
	query, args, err := psql.
		Select("*").
		From("session").
		Where(sq.Eq{"session_id": sessionID}).
		ToSql()

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return Session{}, NewQueryBuildError("session by session ID", err)
	}

	session, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[Session])
	if err != nil {
		return Session{}, fmt.Errorf("Could not collect session: %v", err)
	}

	return session, nil
}

func insertSession(ctx context.Context, tx pgx.Tx, s Session) (int64, error) {
	query, args, err := psql.
		Insert("session").
		Columns("session_id", "member_id", "secret_hash", "created_at").
		Values(s.SessionID, s.MemberID, s.SecretHash, s.CreatedAt).
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

func (s Session) Assemble(context.Context, pgx.Tx) (auth.Session, error) {
	return auth.Session{
		ID:         auth.SessionID(s.SessionID),
		CreatedAt:  s.CreatedAt,
		MemberID:   s.MemberID,
		SecretHash: s.SecretHash,
	}, nil
}
