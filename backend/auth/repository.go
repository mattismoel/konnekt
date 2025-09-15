package auth

import "context"

type SessionRepo interface {
	InsertSession(context.Context, Session) (int64, error)
	GetSession(context.Context, SessionID) (Session, error)
	DeleteSession(context.Context, SessionID) error
}
