package auth

import (
	"errors"
	"slices"
)

var (
	ErrMissingPermissions = errors.New("One or more permissions are missing")
)

type Permission struct {
	ID          int64
	Name        string
	DisplayName string
	Description string
}

type PermissionCollection []Permission

// Returns all permission names of the collection.
func (c PermissionCollection) Names() []string {
	names := make([]string, 0)
	for _, perm := range c {
		names = append(names, perm.Name)
	}

	return names
}

// Returns whether or not the permission collection has includes all required
// permissions.
func (c PermissionCollection) ContainsAll(requiredPerms ...string) error {
	permNames := c.Names()

	for _, requiredPerm := range requiredPerms {
		if !slices.Contains(permNames, requiredPerm) {
			return ErrMissingPermissions
		}
	}

	return nil
}
