package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/mattismoel/konnekt/internal/domain/auth"
	"github.com/mattismoel/konnekt/internal/domain/member"
	"github.com/mattismoel/konnekt/internal/domain/team"
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

	memberTeams, err := memberTeams(ctx, tx, memberID)
	if err != nil {
		return member.Member{}, err
	}

	memeberPerms, err := memberPermissions(ctx, tx, memberID)
	if err != nil {
		return member.Member{}, err
	}

	if err := tx.Commit(); err != nil {
		return member.Member{}, err
	}

	return m.ToInternal(memberTeams.ToInternal(), memeberPerms.ToInternal()), nil
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
		memberTeams, err := memberTeams(ctx, tx, dbMember.ID)
		if err != nil {
			return query.ListResult[member.Member]{}, err
		}

		memberPerms, err := memberPermissions(ctx, tx, dbMember.ID)
		if err != nil {
			return query.ListResult[member.Member]{}, err
		}

		teams, perms := memberTeams.ToInternal(), memberPerms.ToInternal()

		members = append(members, dbMember.ToInternal(teams, perms))
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
		FirstName: m.FirstName,
		LastName:  m.LastName,
		Email:     m.Email,
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

	memberTeams, err := memberTeams(ctx, tx, m.ID)
	if err != nil {
		return member.Member{}, err
	}

	memberPerms, err := memberPermissions(ctx, tx, m.ID)
	if err != nil {
		return member.Member{}, err
	}

	if err := tx.Commit(); err != nil {
		return member.Member{}, err
	}

	return m.ToInternal(memberTeams.ToInternal(), memberPerms.ToInternal()), nil
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
	query := `
	INSERT OR IGNORE INTO member (email, first_name, last_name, password_hash, profile_picture_url) 
	VALUES (@email, @first_name, @last_name, @password_hash, @profile_picture_url)`

	res, err := tx.ExecContext(ctx, query,
		sql.Named("email", m.Email),
		sql.Named("first_name", m.FirstName),
		sql.Named("last_name", m.LastName),
		sql.Named("password_hash", m.PasswordHash),
		sql.Named("profile_picture_url", m.ProfilePictureURL),
	)

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
	query := "select id from member where email = @email"

	var id int64

	err := tx.QueryRowContext(ctx, query, sql.Named("email", email)).Scan(&id)
	if err != nil {
		return Member{}, err
	}

	m, err := memberByID(ctx, tx, id)

	return m, nil
}

func memberByID(ctx context.Context, tx *sql.Tx, memberID int64) (Member, error) {
	query := `
  SELECT email, first_name, last_name, active, password_hash, profile_picture_url FROM member 
  WHERE id = @member_id`

	var email, firstName, lastName, profilePictureURL string
	var active bool
	var passwordHash []byte

	err := tx.QueryRowContext(ctx, query, sql.Named("member_id", memberID)).Scan(
		&email, &firstName, &lastName, &active, &passwordHash, &profilePictureURL,
	)

	if err != nil {
		return Member{}, err
	}

	return Member{
		ID:                memberID,
		Email:             email,
		FirstName:         firstName,
		LastName:          lastName,
		ProfilePictureURL: profilePictureURL,
		Active:            active,
		PasswordHash:      passwordHash,
	}, nil
}

func memberPasswordHash(ctx context.Context, tx *sql.Tx, memberID int64) ([]byte, error) {
	query := `SELECT password_hash FROM member WHERE id = @id`

	var passwordHash []byte

	err := tx.QueryRowContext(ctx, query, sql.Named("id", memberID)).Scan(&passwordHash)
	if err != nil {
		return nil, err
	}

	return passwordHash, nil
}

func listMembers(ctx context.Context, tx *sql.Tx, q QueryParams) (MemberCollection, error) {
	dbQuery, err := NewQuery(`
		SELECT 
			id, 
			first_name, 
			last_name, 
			email, 
			profile_picture_url,
			active,
			password_hash
		FROM member`)

	if err != nil {
		return nil, err
	}

	if err := dbQuery.WithLimit(q.Limit); err != nil {
		return nil, err
	}

	if err := dbQuery.WithOffset(q.Offset); err != nil {
		return nil, err
	}

	active := true

	if filters, ok := q.Filters["active"]; ok {
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

	err = dbQuery.AddFilter("active", query.Equal, activeVal)
	if err != nil {
		return nil, err
	}

	queryStr, args := dbQuery.Build()

	rows, err := tx.QueryContext(ctx, queryStr, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	members := make(MemberCollection, 0)

	for rows.Next() {
		var id int64
		var firstName, lastName, email, profilePictureURL string
		var active bool
		var passwordhash []byte

		err := rows.Scan(&id, &firstName, &lastName, &email, &profilePictureURL, &active, &passwordhash)
		if err != nil {
			return nil, err
		}

		members = append(members, Member{
			ID:                id,
			FirstName:         firstName,
			LastName:          lastName,
			ProfilePictureURL: profilePictureURL,
			Active:            active,
			Email:             email,
			PasswordHash:      passwordhash,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return members, nil
}

func updateMember(ctx context.Context, tx *sql.Tx, memberID int64, m Member) error {
	query := `
		update member set
		first_name = CASE 
			WHEN @first_name = '' THEN first_name 
			ELSE @first_name
		END,
		last_name = CASE 
			WHEN @last_name = '' THEN last_name 
			ELSE @last_name
		END,
		email = CASE 
			WHEN @email = '' THEN email 
			ELSE @email
		END
		WHERE id = @member_id`

	res, err := tx.ExecContext(ctx, query,
		sql.Named("first_name", m.FirstName),
		sql.Named("last_name", m.LastName),
		sql.Named("email", m.Email),
		sql.Named("member_id", memberID),
	)

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
	query := `UPDATE member SET active = "TRUE" WHERE id = @member_id`

	res, err := tx.ExecContext(ctx, query, sql.Named("member_id", memberID))
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
	query := "DELETE FROM member WHERE id = @member_id"

	_, err := tx.ExecContext(ctx, query, sql.Named("member_id", memberID))
	if err != nil {
		return err
	}

	return nil
}

func deleteMemberTeams(ctx context.Context, tx *sql.Tx, memberID int64) error {
	query := "DELETE FROM members_teams WHERE member_id = @member_id"

	_, err := tx.ExecContext(ctx, query, sql.Named("member_id", memberID))
	if err != nil {
		return err
	}

	return nil
}

func memberCount(ctx context.Context, tx *sql.Tx) (int, error) {
	var count int
	query := "SELECT COUNT(*) FROM member"

	err := tx.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (m Member) ToInternal(teams []team.Team, perms auth.PermissionCollection) member.Member {
	return member.Member{
		ID:                m.ID,
		FirstName:         m.FirstName,
		LastName:          m.LastName,
		Email:             m.Email,
		PasswordHash:      m.PasswordHash,
		ProfilePictureURL: m.ProfilePictureURL,

		Active: m.Active,

		Teams:       teams,
		Permissions: perms,
	}
}
