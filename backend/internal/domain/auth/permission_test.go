package auth_test

import (
	"errors"
	"slices"
	"testing"

	"github.com/mattismoel/konnekt/internal/domain/auth"
)

var basePerms = []auth.Permission{
	{
		ID:          1,
		Name:        "view:event",
		DisplayName: "Event (view)",
		Description: "Permission to view events.",
	},
	{
		ID:          2,
		Name:        "delete:event",
		DisplayName: "Event (delete)",
		Description: "Permission to delete events.",
	},
	{
		ID:          2,
		Name:        "delete:artist",
		DisplayName: "Artist (delete)",
		Description: "Permission to delete artists.",
	},
}

func TestNames(t *testing.T) {
	type test struct {
		pc        auth.PermissionCollection
		wantNames []string
	}

	tests := map[string]test{
		"Valid permissions": {
			pc:        basePerms,
			wantNames: []string{"view:event", "delete:event", "delete:artist"},
		},
		"Empty permissions": {
			pc:        make(auth.PermissionCollection, 0),
			wantNames: make([]string, 0),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			names := tt.pc.Names()
			if !slices.Equal(tt.wantNames, names) {
				t.Fatalf("got %+v, want %+v\n", names, tt.wantNames)
			}
		})
	}
}

func TestContainsAll(t *testing.T) {
	type test struct {
		pc        auth.PermissionCollection
		permNames []string
		wantErr   error
	}

	tests := map[string]test{
		"All contained": {
			pc:        basePerms,
			permNames: []string{"view:event", "delete:event"},
			wantErr:   nil,
		},
		"One contained": {
			pc:        basePerms,
			permNames: []string{"view:event", "delete:team"},
			wantErr:   auth.ErrMissingPermissions,
		},
		"None contained": {
			pc:        basePerms,
			permNames: []string{"view:member", "delete:team"},
			wantErr:   auth.ErrMissingPermissions,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			err := tt.pc.ContainsAll(tt.permNames...)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got %v, want %v\n", err, tt.wantErr)
			}
		})
	}
}
