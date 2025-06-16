package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"encoding/hex"
	"errors"
	"fmt"
	"time"
)

var (
	ErrNoSession       = errors.New("No such session")
	ErrInvalidSession  = errors.New("Session is invalid")
	ErrInvalidMemberID = errors.New("Member ID must be valid")
)

const (
	MAX_LIFETIME = 30 * 24 * time.Hour // 30 days.
)

type sessionCfgFunc func(s *Session) error
type SessionToken string
type SessionID string

type Session struct {
	ID        SessionID
	MemberID  int64
	ExpiresAt time.Time
}

func NewSession(cfgs ...sessionCfgFunc) (Session, error) {
	s := Session{}

	if err := s.WithCfgs(cfgs...); err != nil {
		return Session{}, err
	}

	return s, nil
}

func (s *Session) WithCfgs(cfgs ...sessionCfgFunc) error {
	for _, cfg := range cfgs {
		if err := cfg(s); err != nil {
			return err
		}
	}

	return nil
}

func WithToken(t SessionToken) sessionCfgFunc {
	return func(s *Session) error {
		sessionID := t.SessionID()
		s.ID = sessionID
		return nil
	}
}

func WithMemberID(id int64) sessionCfgFunc {
	return func(s *Session) error {
		if id <= 0 {
			return ErrInvalidMemberID
		}

		s.MemberID = id
		return nil
	}
}

func WithLifetime(d time.Duration) sessionCfgFunc {
	return func(s *Session) error {
		if d > MAX_LIFETIME {
			d = MAX_LIFETIME
		}

		expiry := time.Now().Add(d)

		s.ExpiresAt = expiry
		return nil
	}
}

func NewSessionToken() (SessionToken, error) {
	bytes := make([]byte, 20)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", fmt.Errorf("Could not read random bytes into buffer: %v", err)
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
