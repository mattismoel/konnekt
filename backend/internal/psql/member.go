package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	konnekt "github.com/mattismoel/konnekt/backend"
	"github.com/mattismoel/konnekt/backend/api"
	"github.com/mattismoel/konnekt/backend/internal/server"
	"golang.org/x/crypto/bcrypt"
)

var _ server.MemberRepo = MemberRepo{}

type MemberRepo struct {
	Pool *pgxpool.Pool
}

// ListMembers implements server.MemberRepo.
func (m MemberRepo) ListMembers(ctx context.Context, lr api.ListRequest) (api.ListResponse[konnekt.Member], error) {
	members := make([]konnekt.Member, 0)
	err := pgx.BeginFunc(ctx, m.Pool, func(tx pgx.Tx) error {
		pg := paginationFromListRequest(lr)
		dbMembers, err := listMembers(ctx, tx, pg)
		if err != nil {
			return fmt.Errorf("Could not list members: %v", err)
		}

		members, err = dbMembers.Assemble(ctx, tx)
		if err != nil {
			return fmt.Errorf("Could not assemble members: %v", err)
		}

		return nil
	})

	if err != nil {
		return api.ListResponse[konnekt.Member]{}, err
	}

	return api.ListResponse[konnekt.Member]{
		Records: members,
	}, nil
}

type Member struct {
	ID           int64          `db:"id"`
	Email        string         `db:"email"`
	FirstName    string         `db:"first_name"`
	LastName     string         `db:"last_name"`
	PasswordHash []byte         `db:"password_hash"`
	AvatarURL    string         `db:"avatar_url"`
	SpecialRole  sql.NullString `db:"special_role"`
	Approved     bool           `db:"approved"`
	Timestamps
}

// MemberPasswordHash implements konnekt.MemberRepo.
func (m MemberRepo) MemberPasswordHash(ctx context.Context, memberID int64) ([]byte, error) {
	var hash []byte

	err := pgx.BeginFunc(ctx, m.Pool, func(tx pgx.Tx) error {
		member, err := memberByID(ctx, tx, memberID)
		if err != nil {
			return nil
		}

		hash = member.PasswordHash
		return nil
	})

	if err != nil {
		return nil, err
	}

	return hash, nil
}

// MemberByEmail implements konnekt.MemberRepo.
func (m MemberRepo) MemberByEmail(ctx context.Context, email string) (konnekt.Member, error) {
	var member konnekt.Member
	err := pgx.BeginFunc(ctx, m.Pool, func(tx pgx.Tx) error {
		dbMember, err := memberByEmail(ctx, tx, email)
		if err != nil {
			return fmt.Errorf("Could not find member with email %q: %v", email, err)
		}

		member, err = dbMember.Assemble(ctx, tx)
		if err != nil {
			return fmt.Errorf("Could not assemble DB member: %v", err)
		}

		return nil
	})

	if err != nil {
		return konnekt.Member{}, err
	}

	return member, nil
}

// InsertMember implements konnekt.MemberRepo.
func (m MemberRepo) InsertMember(ctx context.Context, cm konnekt.CreateMember) (int64, error) {
	var memberID int64

	hash, err := bcrypt.GenerateFromPassword([]byte(cm.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, fmt.Errorf("Could not generate password hash: %v", err)
	}

	err = pgx.BeginFunc(ctx, m.Pool, func(tx pgx.Tx) error {
		insertedID, err := insertMember(ctx, tx, Member{
			Email:        cm.Email,
			FirstName:    cm.FirstName,
			LastName:     cm.LastName,
			PasswordHash: hash,
			AvatarURL:    cm.AvatarURL,
		})

		if err != nil {
			return fmt.Errorf("Could not insert member: %v", err)
		}

		memberID = insertedID
		return nil
	})

	if err != nil {
		return 0, err
	}

	return memberID, nil
}

// MemberByID implements konnekt.MemberRepo.
func (m MemberRepo) MemberByID(ctx context.Context, memberID int64) (konnekt.Member, error) {
	var member konnekt.Member
	err := pgx.BeginFunc(ctx, m.Pool, func(tx pgx.Tx) error {
		dbMember, err := memberByID(ctx, tx, memberID)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return err
			}
			return fmt.Errorf("Could not get member with ID: %d: %v", memberID, err)
		}

		member, err = dbMember.Assemble(ctx, tx)
		if err != nil {
			return fmt.Errorf("Could not assemble DB member: %v", err)
		}

		return nil
	})

	if err != nil {
		return konnekt.Member{}, err
	}

	return member, nil
}

func insertMember(ctx context.Context, tx pgx.Tx, m Member) (int64, error) {
	query, args, err := psql.
		Insert("member").
		Columns("email", "first_name", "last_name", "password_hash", "avatar_url").
		Values(m.Email, m.FirstName, m.LastName, m.PasswordHash, m.AvatarURL).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return 0, NewQueryBuildError("insert member", err)
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return 0, fmt.Errorf("Could not insert member: %v", err)
	}

	id, err := pgx.CollectExactlyOneRow(rows, pgx.RowTo[int64])
	if err != nil {
		return 0, fmt.Errorf("Could not get inserted ID: %v", err)
	}

	return id, nil
}

func memberByID(ctx context.Context, tx pgx.Tx, memberID int64) (Member, error) {
	query, args, err := psql.
		Select("*").
		From("member").
		Where(sq.Eq{"id": memberID}).
		ToSql()

	if err != nil {
		return Member{}, NewQueryBuildError("member by ID", err)
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return Member{}, fmt.Errorf("Could not query for member with ID %d: %v", memberID, err)
	}

	member, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[Member])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Member{}, err
		}
		return Member{}, fmt.Errorf("Could not collect member row: %v", err)
	}

	return member, nil
}

func memberByEmail(ctx context.Context, tx pgx.Tx, email string) (Member, error) {
	query, args, err := psql.
		Select("*").
		From("member").
		Where(sq.Eq{"email": email}).
		ToSql()

	if err != nil {
		return Member{}, NewQueryBuildError("member by email", err)
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return Member{}, fmt.Errorf("Could not query for member with email %q: %v", email, err)
	}

	member, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[Member])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Member{}, err
		}
		return Member{}, fmt.Errorf("Could not collect member row: %v", err)
	}

	return member, nil
}

func listMembers(ctx context.Context, tx pgx.Tx, pg Pagination) (Collection[Member, konnekt.Member], error) {
	builder := psql.Select("*").From("member")
	builder = applyPagination(builder, pg)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, NewQueryBuildError("list members", err)
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("Could not query for members: %v", err)
	}

	members, err := pgx.CollectRows(rows, pgx.RowToStructByName[Member])
	if err != nil {
		return nil, fmt.Errorf("Could not collect members: %v", err)
	}

	return members, nil
}

func (m Member) Assemble(ctx context.Context, tx pgx.Tx) (konnekt.Member, error) {
	return konnekt.Member{
		ID:          m.ID,
		Email:       m.Email,
		FirstName:   m.FirstName,
		LastName:    m.LastName,
		AvatarURL:   m.AvatarURL,
		SpecialRole: m.SpecialRole.String,
		Approved:    m.Approved,
	}, nil
}
