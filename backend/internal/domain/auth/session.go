package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"encoding/hex"
	"errors"
	"time"
)

var (
	ErrNoSession      = errors.New("No such session")
	ErrInvalidSession = errors.New("Session is invalid")
)

type SessionToken string
type SessionID string

type Session struct {
	ID        SessionID
	MemberID  int64
	ExpiresAt time.Time
}

func NewSession(token SessionToken, memberID int64, lifetime time.Duration) Session {
	expiry := time.Now().Add(lifetime)

	return Session{
		ID:        token.SessionID(),
		MemberID:  memberID,
		ExpiresAt: expiry,
	}
}

func NewSessionToken() (SessionToken, error) {
	bytes := make([]byte, 20)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	encoder := base32.StdEncoding.WithPadding(base32.NoPadding)
	token := encoder.EncodeToString(bytes)
	return SessionToken(token), nil
}

func (token SessionToken) SessionID() SessionID {
	hash := sha256.Sum256([]byte(token))

	sessionID := hex.EncodeToString(hash[:])

	return SessionID(sessionID)
}

// Returns whether or not the session has passed its expiry date.
func (s Session) IsExpired() bool {
	if time.Now().After(s.ExpiresAt) {
		return true
	}

	return false
}

// Returns whether or not the session is refreshable given a refresh buffer duration.
// If the time of calling is within the sessions refresh buffer, the session
// is considered refreshable.
func (s Session) IsRefreshable(buffer time.Duration) bool {
	now := time.Now()

	if now.After(s.ExpiresAt) {
		return false
	}

	if now.After(s.ExpiresAt.Add(-buffer.Abs())) {
		return true
	}

	return false
}
