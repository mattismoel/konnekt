package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/mattismoel/konnekt/internal/domain/member"
	"github.com/mattismoel/konnekt/internal/query"
)

var _ member.Repository = (*MemberRepository)(nil)

type Member struct {
	ID                int64
	Email             string
	FirstName         string
	LastName          string
	PasswordHash      []byte
	Active            bool
	ProfilePictureURL string
}

type MemberCollection []Member

type MemberRepository struct {
	db *sql.DB
}

func NewMemberRepository(db *sql.DB) (*MemberRepository, error) {
	return &MemberRepository{
		db: db,
	}, nil
}

func (repo MemberRepository) Insert(ctx context.Context, m member.Member) (int64, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	defer tx.Rollback()

	memberID, err := insertMember(ctx, tx, Member{
		ID:                m.ID,
		Email:             m.Email,
		FirstName:         m.FirstName,
		LastName:          m.LastName,
		PasswordHash:      m.PasswordHash,
		ProfilePictureURL: m.ProfilePictureURL,
	})

	if err != nil {
		switch {
		case errors.Is(err, ErrAlreadyExists):
			return 0, member.ErrAlreadyExists
		}
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return memberID, nil
}

func (repo MemberRepository) ByID(ctx context.Context, memberID int64) (member.Member, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return member.Member{}, err
	}

	defer tx.Rollback()

	m, err := memberByID(ctx, tx, memberID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return member.Member{}, ErrNotFound
		default:
			return member.Member{}, err
		}
	}

	if err := tx.Commit(); err != nil {
		return member.Member{}, err
	}

	return m.ToInternal(), nil
}

func (repo MemberRepository) Approve(ctx context.Context, memberID int64) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	if err := approveMember(ctx, tx, memberID); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// TODO: Implement...
func (repo MemberRepository) SetProfilePictureURL(ctx context.Context, memberID int64, url string) error {
	return nil
}

func (repo MemberRepository) SetMemberTeams(ctx context.Context, memberID int64, teamIDs ...int64) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	if err := deleteMemberTeams(ctx, tx, memberID); err != nil {
		return err
	}

	for _, teamID := range teamIDs {
		if err := associateMemberWithTeam(ctx, tx, memberID, teamID); err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (repo MemberRepository) List(ctx context.Context, q query.ListQuery) (query.ListResult[member.Member], error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return query.ListResult[member.Member]{}, err
	}

	defer tx.Rollback()

	dbMembers, err := listMembers(ctx, tx, QueryParams{
		Offset:  q.Offset(),
		Limit:   q.Limit,
		OrderBy: q.OrderBy,
		Filters: q.Filters,
	})

	if err != nil {
		return query.ListResult[member.Member]{}, err
	}

	members := make([]member.Member, 0)

	for _, dbMember := range dbMembers {
		members = append(members, dbMember.ToInternal())
	}

	totalCount, err := memberCount(ctx, tx)
	if err != nil {
		return query.ListResult[member.Member]{}, err
	}

	if err := tx.Commit(); err != nil {
		return query.ListResult[member.Member]{}, err
	}

	return query.ListResult[member.Member]{
		Page:       q.Page,
		PerPage:    q.PerPage,
		TotalCount: totalCount,
		PageCount:  q.PageCount(totalCount),
		Records:    members,
	}, nil
}

func (repo MemberRepository) Update(ctx context.Context, memberID int64, m member.Member) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	err = updateMember(ctx, tx, memberID, Member{
		FirstName:         m.FirstName,
		LastName:          m.LastName,
		Email:             m.Email,
		ProfilePictureURL: m.ProfilePictureURL,
	})

	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (repo MemberRepository) ByEmail(ctx context.Context, email string) (member.Member, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return member.Member{}, err
	}

	defer tx.Rollback()

	m, err := memberByEmail(ctx, tx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return member.Member{}, member.ErrNotFound
		}

		return member.Member{}, err
	}

	if err := tx.Commit(); err != nil {
		return member.Member{}, err
	}

	return m.ToInternal(), nil
}

func (repo MemberRepository) Delete(ctx context.Context, memberID int64) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	if err := deleteMember(ctx, tx, memberID); err != nil {
		return err
	}

	if err := deleteMemberTeams(ctx, tx, memberID); err != nil {
		return err
	}

	if err := deleteMemberSession(ctx, tx, memberID); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
func (repo MemberRepository) PasswordHash(ctx context.Context, memberID int64) (member.PasswordHash, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	ph, err := memberPasswordHash(ctx, tx, memberID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, member.ErrNotFound
		}

		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return ph, nil
}

func insertMember(ctx context.Context, tx *sql.Tx, m Member) (int64, error) {
	query, args, err := sq.
		Insert("member").
		Options("OR IGNORE").
		Columns("email", "first_name", "last_name", "password_hash", "profile_picture_url").
		Values(m.Email, m.FirstName, m.LastName, m.PasswordHash, m.ProfilePictureURL).
		ToSql()

	if err != nil {
		return 0, err
	}

	res, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	if rowsAffected <= 0 {
		return 0, ErrAlreadyExists
	}

	memberID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return memberID, nil
}

func memberByEmail(ctx context.Context, tx *sql.Tx, email string) (Member, error) {
	query, args, err := memberBuilder.
		Where(sq.Eq{"email": email}).
		ToSql()

	if err != nil {
		return Member{}, err
	}

	var m Member

	row := tx.QueryRowContext(ctx, query, args...)
	if err := scanMember(row, &m); err != nil {
		return Member{}, err
	}

	return m, nil
}

func memberByID(ctx context.Context, tx *sql.Tx, memberID int64) (Member, error) {
	query, args, err := memberBuilder.
		Where(sq.Eq{"id": memberID}).
		ToSql()

	if err != nil {
		return Member{}, err
	}

	var m Member
	row := tx.QueryRowContext(ctx, query, args...)
	if err := scanMember(row, &m); err != nil {
		return Member{}, err
	}

	return m, nil
}

func memberPasswordHash(ctx context.Context, tx *sql.Tx, memberID int64) ([]byte, error) {
	query, args, err := sq.
		Select("password_hash").
		From("member").
		Where(sq.Eq{"id": memberID}).
		ToSql()

	var passwordHash []byte

	err = tx.
		QueryRowContext(ctx, query, args...).
		Scan(&passwordHash)

	if err != nil {
		return nil, err
	}

	return passwordHash, nil
}

func scanMember(s Scanner, dst *Member) error {
	err := s.Scan(
		&dst.ID,
		&dst.FirstName,
		&dst.LastName,
		&dst.Email,
		&dst.ProfilePictureURL,
		&dst.Active,
		&dst.PasswordHash,
	)

	if err != nil {
		return err
	}

	return nil
}

var memberBuilder = sq.
	Select(
		"id",
		"first_name",
		"last_name",
		"email",
		"profile_picture_url",
		"active",
		"password_hash",
	).
	From("member")

func listMembers(ctx context.Context, tx *sql.Tx, params QueryParams) (MemberCollection, error) {
	builder := memberBuilder

	if params.Limit > 0 {
		builder = builder.Limit(uint64(params.Limit))
	}

	if params.Offset > 0 {
		builder = builder.Offset(uint64(params.Offset))
	}

	active := true

	if filters, ok := params.Filters["active"]; ok {
		for _, filter := range filters {
			val := strings.ToUpper(filter.Value)
			if val == "FALSE" {
				active = false
			} else if val == "TRUE" {
				active = true
			}
		}
	}

	activeVal := "TRUE"
	if !active {
		activeVal = "FALSE"
	}

	builder = builder.Where(sq.Eq{"active": activeVal})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	members := make(MemberCollection, 0)

	for rows.Next() {
		var m Member
		if err := scanMember(rows, &m); err != nil {
			return nil, err
		}

		members = append(members, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return members, nil
}

func updateMember(ctx context.Context, tx *sql.Tx, memberID int64, m Member) error {
	builder := sq.Update("member").Where(sq.Eq{"id": memberID})

	if m.FirstName != "" {
		builder = builder.Set("first_name", m.FirstName)
	}

	if m.LastName != "" {
		builder = builder.Set("last_name", m.LastName)
	}

	if m.Email != "" {
		builder = builder.Set("email", m.Email)
	}

	if m.ProfilePictureURL != "" {
		builder = builder.Set("profile_picture_url", m.ProfilePictureURL)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	res, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected <= 0 {
		return ErrNotFound
	}

	return nil
}

func approveMember(ctx context.Context, tx *sql.Tx, memberID int64) error {
	query, args, err := sq.
		Update("member").
		Set("active", "TRUE").
		Where(sq.Eq{"id": memberID}).
		ToSql()

	if err != nil {
		return err
	}

	res, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected <= 0 {
		return ErrNotFound
	}

	return nil
}

func deleteMember(ctx context.Context, tx *sql.Tx, memberID int64) error {
	query, args, err := sq.
		Delete("member").
		Where(sq.Eq{"id": memberID}).
		ToSql()

	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func deleteMemberTeams(ctx context.Context, tx *sql.Tx, memberID int64) error {
	query, args, err := sq.
		Delete("members_teams").
		Where(sq.Eq{"member_id": memberID}).
		ToSql()

	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func memberCount(ctx context.Context, tx *sql.Tx) (int, error) {
	var count int
	query, args, err := sq.
		Select("COUNT(*)").
		From("member").
		ToSql()

	if err != nil {
		return 0, err
	}

	err = tx.
		QueryRowContext(ctx, query, args...).
		Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (m Member) ToInternal() member.Member {
	return member.Member{
		ID:                m.ID,
		FirstName:         m.FirstName,
		LastName:          m.LastName,
		Email:             m.Email,
		PasswordHash:      m.PasswordHash,
		ProfilePictureURL: m.ProfilePictureURL,

		Active: m.Active,
	}
}
