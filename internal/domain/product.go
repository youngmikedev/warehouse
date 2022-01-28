package domain

import "time"

type Product struct {
	ID        int
	Price     int
	Article   string
	Name      string
	CreatedAt time.Time
}
