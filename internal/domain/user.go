package domain

import "time"

type User struct {
	ID        int
	Name      string
	Email     string
	CreatedAt time.Time
}

type Session struct {
	ID           int
	UserID       int
	RefreshToken string
	ExpiresAt    time.Duration
	UpdatedAt    time.Time
	CreatedAt    time.Time
}
