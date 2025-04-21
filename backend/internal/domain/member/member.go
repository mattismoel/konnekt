package member

import (
	"errors"
	"net/http"
	"net/mail"
	"net/url"
	"strings"

	"github.com/mattismoel/konnekt/internal/domain/auth"
)

var (
	ErrIDInvalid = errors.New("ID must be a positive integer")

	ErrFirstNameInvalid = errors.New("First name must be valid and non-empty")

	ErrLastNameInvalid = errors.New("Last name must be valid and non-empty")

	ErrEmailInvalid = errors.New("Email must be valid")

	ErrPasswordHashInvalid = errors.New("Password hash must be a non-empty byte array")

	ErrProfileImageURLInvalid      = errors.New("Profile image URL must be valid")
	ErrProfileImageURLInaccessible = errors.New("Profile image URL must be accessible")
)

type Member struct {
	ID              int64        `json:"id"`
	FirstName       string       `json:"firstName"`
	LastName        string       `json:"lastName"`
	Email           string       `json:"email"`
	PasswordHash    PasswordHash `json:"-"`
	Active bool `json:"active"`

	Roles       auth.RoleCollection       `json:"roles"`
	Permissions auth.PermissionCollection `json:"permissions"`

	ProfileImageURL string `json:"profileImageUrl"`
}

type cfgFunc func(m *Member) error

func NewMember(cfgs ...cfgFunc) (Member, error) {
	u := &Member{
		Active:      false,
		Roles:       make(auth.RoleCollection, 0),
		Permissions: make(auth.PermissionCollection, 0),
	}

	for _, cfg := range cfgs {
		if err := cfg(u); err != nil {
			return Member{}, err
		}
	}

	return *u, nil
}

func WithID(id int64) cfgFunc {
	return func(m *Member) error {
		if id <= 0 {
			return ErrIDInvalid
		}

		m.ID = id

		return nil
	}
}

func WithFirstName(firstName string) cfgFunc {
	firstName = strings.TrimSpace(firstName)
	return func(m *Member) error {
		if firstName == "" {
			return ErrFirstNameInvalid
		}

		m.FirstName = firstName

		return nil
	}
}

func WithLastName(lastName string) cfgFunc {
	lastName = strings.TrimSpace(lastName)

	return func(m *Member) error {
		if lastName == "" {
			return ErrLastNameInvalid
		}

		m.LastName = lastName

		return nil
	}
}

func WithEmail(email string) cfgFunc {
	email = strings.TrimSpace(email)

	return func(m *Member) error {
		if email == "" {
			return ErrEmailInvalid
		}

		mail, err := mail.ParseAddress(email)
		if err != nil {
			return ErrEmailInvalid
		}

		m.Email = mail.Address

		return nil
	}
}

func WithPasswordHash(hash []byte) cfgFunc {
	return func(m *Member) error {
		if len(hash) <= 0 {
			return ErrPasswordHashInvalid
		}

		m.PasswordHash = hash

		return nil
	}
}

func WithPermissions(perms auth.PermissionCollection) cfgFunc {
	return func(m *Member) error {
		m.Permissions = perms
		return nil
	}
}

func WithRoles(roles auth.RoleCollection) cfgFunc {
	return func(m *Member) error {
		m.Roles = roles
		return nil
	}
}

func WithProfileImageURL(imageUrl string) cfgFunc {
	return func(m *Member) error {
		u, err := url.Parse(imageUrl)
		if err != nil {
			return ErrProfileImageURLInvalid
		}

		resp, err := http.Get(u.String())
		if err != nil {
			return err
		}

		if resp.StatusCode < 200 || resp.StatusCode >= 400 {
			return ErrProfileImageURLInaccessible
		}

		m.ProfileImageURL = u.String()

		return nil
	}
}
