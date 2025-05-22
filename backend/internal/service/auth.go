package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/mattismoel/konnekt/internal/domain/auth"
	"github.com/mattismoel/konnekt/internal/domain/member"
	"github.com/mattismoel/konnekt/internal/domain/team"
	"github.com/mattismoel/konnekt/internal/query"
	"golang.org/x/crypto/bcrypt"
)

const (
	SESSION_LIFETIME       = 30 * 24 * time.Hour // 30 day expiry.
	SESSION_REFRESH_BUFFER = 15 * 24 * time.Hour // 15 day refresh buffer.

)

var (
	ErrMemberInactive = errors.New("Member is not active or needs approval")
)

type AuthService struct {
	memberRepo member.Repository
	teamRepo   team.Repository
	authRepo   auth.Repository
}

func NewAuthService(memberRepo member.Repository, authRepo auth.Repository, teamRepo team.Repository) (*AuthService, error) {
	return &AuthService{
		memberRepo: memberRepo,
		teamRepo:   teamRepo,
		authRepo:   authRepo,
	}, nil
}

type RegisterLoad struct {
	Email             string
	Password          auth.Password
	PasswordConfirm   auth.Password
	FirstName         string
	LastName          string
	ProfilePictureURL string
}

func (srv AuthService) Register(ctx context.Context, load RegisterLoad) error {
	// Return if member already exists.
	_, err := srv.memberRepo.ByEmail(ctx, load.Email)
	if err == nil {
		return member.ErrAlreadyExists
	}

	if err := load.Password.Validate(); err != nil {
		return err
	}

	if err := load.Password.Matches(load.PasswordConfirm); err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword(load.Password, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	m, err := member.NewMember(
		member.WithEmail(load.Email),
		member.WithFirstName(load.FirstName),
		member.WithLastName(load.LastName),
		member.WithPasswordHash(hash),
	)

	if err != nil {
		return err
	}

	if strings.TrimSpace(load.ProfilePictureURL) != "" {
		err := m.WithCfgs(member.WithProfilePictureURL(load.ProfilePictureURL))
		if err != nil {
			return err
		}
	}

	memberID, err := srv.memberRepo.Insert(ctx, m)
	if err != nil {
		return err
	}

	team, err := srv.teamRepo.ByName(ctx, "member")
	if err != nil {
		return err
	}

	err = srv.teamRepo.AddMemberTeams(ctx, memberID, team.ID)
	if err != nil {
		return err
	}

	return nil
}

func (srv AuthService) Login(ctx context.Context, email string, password []byte) (auth.SessionToken, time.Time, error) {
	m, err := srv.validateMember(ctx, email, password)
	if err != nil {
		return "", time.Time{}, err
	}

	err = srv.clearMemberSession(ctx, m.ID)
	if err != nil {
		return "", time.Time{}, err
	}

	token, expiry, err := srv.createSession(ctx, m.ID)
	if err != nil {
		return "", time.Time{}, err
	}

	return token, expiry, nil
}

func (srv AuthService) LogOut(ctx context.Context, token auth.SessionToken) error {
	sessionID := token.SessionID()

	session, err := srv.authRepo.Session(ctx, sessionID)
	if err != nil {
		return err
	}

	err = srv.authRepo.DeleteMemberSession(ctx, session.MemberID)
	if err != nil {
		return err
	}

	return nil
}

func (srv AuthService) ValidateSession(ctx context.Context, token auth.SessionToken) (time.Time, error) {
	sessionID := token.SessionID()

	session, err := srv.authRepo.Session(ctx, sessionID)
	if err != nil {
		return time.Time{}, err
	}

	if session.IsExpired() {
		return time.Time{}, auth.ErrInvalidSession
	}

	if session.IsRefreshable(SESSION_REFRESH_BUFFER) {
		newExpiry := time.Now().Add(SESSION_LIFETIME)
		err := srv.authRepo.SetSessionExpiry(ctx, sessionID, newExpiry)
		if err != nil {
			return time.Time{}, err
		}

		return newExpiry, nil
	}

	return session.ExpiresAt, nil
}
func (srv AuthService) Session(ctx context.Context, id auth.SessionID) (auth.Session, error) {
	session, err := srv.authRepo.Session(ctx, id)
	if err != nil {
		return auth.Session{}, err
	}

	return session, nil
}

// Checks whether or not a member has all required permissions.
func (srv AuthService) HasPermission(ctx context.Context, memberID int64, permNames ...string) error {
	memberPerms, err := srv.MemberPermissions(ctx, memberID)
	if err != nil {
		return err
	}

	err = memberPerms.ContainsAll(permNames...)
	if err != nil {
		return err
	}

	return nil
}

func (srv AuthService) MemberPermissions(ctx context.Context, memberID int64) (auth.PermissionCollection, error) {
	m, err := srv.memberRepo.ByID(ctx, memberID)
	if err != nil {
		return nil, err
	}

	memberTeams, err := srv.teamRepo.MemberTeams(ctx, m.ID)
	if err != nil {
		return nil, err
	}

	memberPerms := auth.PermissionCollection(make([]auth.Permission, 0))

	for _, team := range memberTeams {
		teamPerms, err := srv.authRepo.TeamPermissions(ctx, team.ID)
		if err != nil {
			return nil, err
		}

		// Add the team perms to the members permissions.
		memberPerms = append(memberPerms, teamPerms...)
	}

	return memberPerms, nil

}

func (srv AuthService) validateMember(ctx context.Context, email string, password []byte) (member.Member, error) {
	// Return early if member does not exist.
	m, err := srv.memberRepo.ByEmail(ctx, email)
	if err != nil {
		return member.Member{}, err
	}

	if !m.Active {
		return member.Member{}, ErrMemberInactive
	}

	hash, err := srv.memberRepo.PasswordHash(ctx, m.ID)
	if err != nil {
		return member.Member{}, err
	}

	if err := hash.Matches(password); err != nil {
		return member.Member{}, err
	}

	return m, err
}

func (srv AuthService) clearMemberSession(ctx context.Context, memberID int64) error {
	err := srv.authRepo.DeleteMemberSession(ctx, memberID)
	if err != nil {
		return err
	}

	return nil
}

func (srv AuthService) createSession(ctx context.Context, memberID int64) (auth.SessionToken, time.Time, error) {
	token, err := auth.NewSessionToken()
	if err != nil {
		return "", time.Time{}, err
	}

	session := auth.NewSession(token, memberID, SESSION_LIFETIME)

	err = srv.authRepo.InsertSession(ctx, session)
	if err != nil {
		return "", time.Time{}, err
	}

	return token, session.ExpiresAt, nil
}

func (srv AuthService) ListPermissions(ctx context.Context, q query.ListQuery) (query.ListResult[auth.Permission], error) {
	result, err := srv.authRepo.ListPermissions(ctx, q)
	if err != nil {
		return query.ListResult[auth.Permission]{}, err
	}

	return result, nil
}
