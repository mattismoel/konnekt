package server

import (
	"context"

	konnekt "github.com/mattismoel/konnekt/backend"
)

type MemberRepo interface {
	MemberByID(context.Context, int64) (konnekt.Member, error)
	MemberByEmail(context.Context, string) (konnekt.Member, error)
	InsertMember(context.Context, konnekt.CreateMember) (int64, error)
	MemberPasswordHash(context.Context, int64) ([]byte, error)
}
