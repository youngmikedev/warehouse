package auth

import "time"

type TokenManager interface {
	NewAccessToken() string
	NewRefreshToken() string
	ValidateAccessToken(createdAt time.Time) bool
	ValidateRefreshToken(createdAt time.Time) bool
}

type HashManager interface {
	Hash(str string) (string, error)
	Validate(hashedPassword, password string) bool
}
