package service

import (
	"context"
	"time"

	"github.com/mattismoel/konnekt/internal/domain/auth"
	"github.com/mattismoel/konnekt/internal/domain/user"
	"github.com/mattismoel/konnekt/internal/query"
	"golang.org/x/crypto/bcrypt"
)

const (
	SESSION_LIFETIME       = 30 * 24 * time.Hour // 30 day expiry.
	SESSION_REFRESH_BUFFER = 15 * 24 * time.Hour // 15 day refresh buffer.
)

type AuthService struct {
	userRepo user.Repository
	authRepo auth.Repository
}

func NewAuthService(userRepo user.Repository, authRepo auth.Repository) (*AuthService, error) {
	return &AuthService{
		userRepo: userRepo,
		authRepo: authRepo,
	}, nil
}

func (srv AuthService) Register(ctx context.Context, email string, password []byte, passwordConfirm []byte, firstName, lastName string) (auth.SessionToken, time.Time, error) {
	// Return if user already exists.
	_, err := srv.userRepo.ByEmail(ctx, email)
	if err == nil {
		return "", time.Time{}, err
	}

	if err := auth.DoPasswordsMatch(password, passwordConfirm); err != nil {
		return "", time.Time{}, err
	}

	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", time.Time{}, err
	}

	u, err := user.NewUser(
		user.WithEmail(email),
		user.WithFirstName(firstName),
		user.WithLastName(lastName),
		user.WithPasswordHash(hash),
	)

	if err != nil {
		return "", time.Time{}, err
	}

	userID, err := srv.userRepo.Insert(ctx, u)
	if err != nil {
		return "", time.Time{}, err
	}

	role, err := srv.authRepo.RoleByName(ctx, "member")
	if err != nil {
		return "", time.Time{}, err
	}
	err = srv.authRepo.AddUserRoles(ctx, userID, role.ID)
	if err != nil {
		return "", time.Time{}, err
	}

	token, expiry, err := srv.createSession(ctx, userID)
	if err != nil {
		return "", time.Time{}, err
	}

	return token, expiry, nil
}

func (srv AuthService) Login(ctx context.Context, email string, password []byte) (auth.SessionToken, time.Time, error) {
	usr, err := srv.validateUser(ctx, email, password)
	if err != nil {
		return "", time.Time{}, err
	}

	err = srv.clearUserSession(ctx, usr.ID)
	if err != nil {
		return "", time.Time{}, err
	}

	token, expiry, err := srv.createSession(ctx, usr.ID)
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

	err = srv.authRepo.DeleteUserSession(ctx, session.UserID)
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

// Checks whether or not a user has all required permissions.
func (srv AuthService) HasPermission(ctx context.Context, userID int64, permNames ...string) error {
	userPermissions, err := srv.userPermissions(ctx, userID)
	if err != nil {
		return err
	}

	err = userPermissions.ContainsAll(permNames...)
	if err != nil {
		return err
	}

	return nil
}

func (srv AuthService) UserRoles(ctx context.Context, userID int64) ([]auth.Role, error) {
	_, err := srv.userRepo.ByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	roles, err := srv.authRepo.UserRoles(ctx, userID)
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (srv AuthService) userPermissions(ctx context.Context, userID int64) (auth.PermissionCollection, error) {
	usr, err := srv.userRepo.ByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	userRoles, err := srv.authRepo.UserRoles(ctx, usr.ID)
	if err != nil {
		return nil, err
	}

	userPerms := auth.PermissionCollection(make([]auth.Permission, 0))

	for _, role := range userRoles {
		rolePerms, err := srv.authRepo.RolePermissions(ctx, role.ID)
		if err != nil {
			return nil, err
		}

		// Add the role perms to the users permissions.
		userPerms = append(userPerms, rolePerms...)
	}

	return userPerms, nil

}

func (srv AuthService) validateUser(ctx context.Context, email string, password []byte) (user.User, error) {
	// Return early if user does not exist.
	usr, err := srv.userRepo.ByEmail(ctx, email)
	if err != nil {
		return user.User{}, err
	}

	hash, err := srv.userRepo.PasswordHash(ctx, usr.ID)
	if err != nil {
		return user.User{}, err
	}

	if err := hash.Matches(password); err != nil {
		return user.User{}, err
	}

	return usr, err
}

func (srv AuthService) clearUserSession(ctx context.Context, userID int64) error {
	err := srv.authRepo.DeleteUserSession(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}

func (srv AuthService) createSession(ctx context.Context, userID int64) (auth.SessionToken, time.Time, error) {
	token, err := auth.NewSessionToken()
	if err != nil {
		return "", time.Time{}, err
	}

	session := auth.NewSession(token, userID, SESSION_LIFETIME)

	err = srv.authRepo.InsertSession(ctx, session)
	if err != nil {
		return "", time.Time{}, err
	}

	return token, session.ExpiresAt, nil
}

func (srv AuthService) ListRoles(ctx context.Context, q query.ListQuery) (query.ListResult[auth.Role], error) {
	result, err := srv.authRepo.ListRoles(ctx, q)
	if err != nil {
		return query.ListResult[auth.Role]{}, err
	}

	return result, nil
}

type CreateRole struct {
	Name        string
	DisplayName string
	Description string
}

func (srv AuthService) CreateRole(ctx context.Context, load CreateRole) (auth.Role, error) {
	r, err := auth.NewRole(
		auth.WithName(load.Name),
		auth.WithDisplayName(load.DisplayName),
		auth.WithDescription(load.Description),
	)

	if err != nil {
		return auth.Role{}, err
	}

	roleID, err := srv.authRepo.InsertRole(ctx, r)
	if err != nil {
		return auth.Role{}, err
	}

	role, err := srv.authRepo.RoleByID(ctx, roleID)
	if err != nil {
		return auth.Role{}, err
	}

	return role, nil
}

func (srv AuthService) DeleteRole(ctx context.Context, roleID int64) error {
	err := srv.authRepo.DeleteRole(ctx, roleID)
	if err != nil {
		return err
	}

	return nil
}

func (srv AuthService) ListPermissions(ctx context.Context, q query.ListQuery) (query.ListResult[auth.Permission], error) {
	result, err := srv.authRepo.ListPermissions(ctx, q)
	if err != nil {
		return query.ListResult[auth.Permission]{}, err
	}

	return result, nil
}
