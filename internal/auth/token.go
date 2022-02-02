package auth

import (
	"fmt"
	"math/rand"
	"time"
)

const tokenLength = 30

type RandTokenManager struct {
	signingKey  string
	atExpiresAt time.Duration
	rtExpiresAt time.Duration
}

func NewJWTManager(signingKey string, accessTokenExpiresAt, refreshTokenExpiresAt time.Duration) *RandTokenManager {
	return &RandTokenManager{
		signingKey:  signingKey,
		atExpiresAt: accessTokenExpiresAt,
		rtExpiresAt: refreshTokenExpiresAt,
	}
}

func (m *RandTokenManager) NewAccessToken() string {
	return generateRandomString(tokenLength)
}

func (m *RandTokenManager) NewRefreshToken() string {
	return generateRandomString(tokenLength)
}

func generateRandomString(len int) string {
	b := make([]byte, len)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	_, err := r.Read(b)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%x", b)
}

func (m *RandTokenManager) ValidateAccessToken(createdAt time.Time) bool {
	return createdAt.Add(m.atExpiresAt).After(time.Now())
}

func (m *RandTokenManager) ValidateRefreshToken(createdAt time.Time) bool {
	return createdAt.Add(m.rtExpiresAt).After(time.Now())
}
