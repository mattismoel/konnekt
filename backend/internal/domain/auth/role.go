package auth

import (
	"errors"
	"strings"
)

var (
	ErrRoleIDInvalid          = errors.New("Role ID must be a valid positive integer")
	ErrRoleNameInvalid        = errors.New("Role name must be a valid non-empty string")
	ErrRoleDisplayNameInvalid = errors.New("Role display name must be a valid non-empty string")
	ErrRoleDescriptionInvalid = errors.New("Role description must be a valid non-empty string")
)

type Role struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
}

type RoleCollection []Role

type roleCfgFunc func(r *Role) error

func NewRole(cfgs ...roleCfgFunc) (Role, error) {
	r := &Role{}

	for _, cfg := range cfgs {
		if err := cfg(r); err != nil {
			return Role{}, err
		}
	}

	return *r, nil
}

func WithID(id int64) roleCfgFunc {
	return func(r *Role) error {
		if id <= 0 {
			return ErrRoleIDInvalid
		}

		r.ID = id
		return nil
	}
}

func WithName(name string) roleCfgFunc {
	name = strings.TrimSpace(name)
	return func(r *Role) error {
		if name == "" {
			return ErrRoleNameInvalid
		}

		r.Name = name
		return nil
	}
}

func WithDisplayName(displayName string) roleCfgFunc {
	displayName = strings.TrimSpace(displayName)
	return func(r *Role) error {
		if displayName == "" {
			return ErrRoleDisplayNameInvalid
		}

		r.DisplayName = displayName
		return nil
	}
}
func WithDescription(desc string) roleCfgFunc {
	desc = strings.TrimSpace(desc)
	return func(r *Role) error {
		if desc == "" {
			return ErrRoleDescriptionInvalid
		}

		r.Description = desc
		return nil
	}
}
