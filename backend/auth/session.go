package auth

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	SESSION_ID_LENGTH = 24
	SESSION_LIFETIME  = 1 * 24 * time.Hour // 1 day session lifetime.
)

var (
	ErrInvalidToken = errors.New("Token is invalid")
)

type SessionID string
type SessionToken string

type Session struct {
	ID         SessionID    `json:"id"`
	CreatedAt  time.Time    `json:"createdAt"`
	SecretHash []byte       `json:"-"`
	Token      SessionToken `json:"-"`
}

func CreateSession() (Session, error) {
	id, err := generateSecureRandomBytes()
	if err != nil {
		return Session{}, fmt.Errorf("Could not generate session ID: %v", err)
	}

	secret, err := generateSecureRandomBytes()
	if err != nil {
		return Session{}, fmt.Errorf("Could not generate session secret: %v", err)
	}

	secretHash := hashSecret(secret)
	token := fmt.Sprintf("%s.%s", id, secret)

	return Session{
		ID:         SessionID(id),
		SecretHash: secretHash,
		CreatedAt:  time.Now(),
		Token:      SessionToken(token),
	}, nil
}

func (st SessionToken) SessionID() (SessionID, error) {
	tokenParts := strings.Split(string(st), ".")
	if len(tokenParts) != 2 {
		return "", ErrInvalidToken
	}

	return SessionID(tokenParts[0]), nil
}

func generateSecureRandomBytes() ([]byte, error) {
	const alphabet = "abcdefghijkmnpqrstuvwxyz23456789"

	b := make([]byte, SESSION_ID_LENGTH)
	_, err := rand.Read(b)
	if err != nil {
		return nil, fmt.Errorf("Could not read random bytes: %v", err)
	}

	var id []byte
	for i := range SESSION_ID_LENGTH {
		id = append(id, alphabet[b[i]>>3])
	}

	return id, nil
}

func hashSecret(secret []byte) []byte {
	hash := sha256.New()
	hash.Write(secret)
	return hash.Sum(nil)
}

func (st SessionToken) Validate(secretHash []byte) error {
	parts := strings.Split(string(st), ".")
	if len(parts) != 2 {
		return ErrInvalidToken
	}

	tokenSecretHash := hashSecret([]byte(parts[0]))
	if bytes.Equal(tokenSecretHash, secretHash) {
		return ErrInvalidToken
	}

	return nil
}
