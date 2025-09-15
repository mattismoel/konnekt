package konnekt

import "context"

type Member struct {
	ID          int64  `json:"id"`
	Email       string `json:"email"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	AvatarURL   string `json:"avatarUrl"`
	SpecialRole string `json:"specialRole"`
	Approved    bool   `json:"approved"`
}

type CreateMember struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	AvatarURL string `json:"avatarUrl"`
	Password  string `json:"password"`
}

type MemberRepo interface {
	MemberByID(context.Context, int64) (Member, error)
	MemberByEmail(context.Context, string) (Member, error)
	InsertMember(context.Context, CreateMember) (int64, error)
	MemberPasswordHash(context.Context, int64) ([]byte, error)
}
