package konnekt

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type SessionToken []byte
type SessionID string

type Session struct {
	ID        SessionID `json:"id"`
	UserID    int64     `json:"userId"`
	ExpiresAt time.Time `json:"expiresAt"`
}

func (s Session) HasExpired() bool {
	now := time.Now()

	return now.After(s.ExpiresAt) || now.Equal(s.ExpiresAt)
}

// Returns whether or not a session should be extended.
//
// A session is set to be extended, when the current datetime is within
// the duration of expiration, provided by the extend buffer duration.
func (s Session) ShouldExtend(extendBuffer time.Duration) bool {
	now := time.Now()

	return now.After(s.ExpiresAt.Add(-extendBuffer))
}

// Generates a session id based on the session token.
func (t SessionToken) GenerateSessionID() string {
	hash := sha256.Sum256(t)

	sessionId := hex.EncodeToString(hash[:])
	return sessionId
}
