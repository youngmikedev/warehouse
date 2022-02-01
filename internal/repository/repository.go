package repository

import (
	"context"

	"github.com/imranzahaev/warehouse/internal/domain"
)

type Repositories struct {
	User
}

type User interface {
	Create(ctx context.Context, user domain.User, pwd string) error
	Update(ctx context.Context, user domain.User, pwd string) error
	Get(ctx context.Context, id int) (domain.User, error)
	SetSession(ctx context.Context, session domain.Session) error
	GetSession(ctx context.Context, userID int, refreshToken string) (domain.Session, error)
}
