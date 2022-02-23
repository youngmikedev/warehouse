package auth

import (
	"fmt"
	"math/rand"
	"time"
)

const tokenLength = 30

type RandTokenManager struct {
	signingKey  string
	rs          rand.Source
	atExpiresAt time.Duration
	rtExpiresAt time.Duration
}

func NewTokenManager(signingKey string, accessTokenExpiresAt, refreshTokenExpiresAt time.Duration) *RandTokenManager {
	return &RandTokenManager{
		signingKey:  signingKey,
		rs:          rand.NewSource(time.Now().Unix()),
		atExpiresAt: accessTokenExpiresAt,
		rtExpiresAt: refreshTokenExpiresAt,
	}
}

func (m *RandTokenManager) NewAccessToken() string {
	return m.generateRandomString(tokenLength)
}

func (m *RandTokenManager) NewRefreshToken() string {
	return m.generateRandomString(tokenLength)
}

func (m *RandTokenManager) generateRandomString(len int) string {
	b := make([]byte, len)

	r := rand.New(m.rs)

	_, err := r.Read(b)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%x", b)
}

func (m *RandTokenManager) ValidateAccessToken(createdAt time.Time) bool {
	return createdAt.Add(m.atExpiresAt).Before(time.Now())
}

func (m *RandTokenManager) ValidateRefreshToken(createdAt time.Time) bool {
	return createdAt.Add(m.rtExpiresAt).Before(time.Now())
}
