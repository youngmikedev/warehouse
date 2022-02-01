package repository

import (
	"context"

	"github.com/imranzahaev/warehouse/internal/domain"
)

type Repositories struct {
	User
}

type User interface {
	Create(ctx context.Context, user domain.User, password string) (int, error)
	Update(ctx context.Context, user domain.User, password string) error
	Get(ctx context.Context, id int) (domain.User, error)
	GetByLogin(ctx context.Context, login string) (domain.User, error)
	CreateSession(ctx context.Context, session domain.Session) (int, error)
	UpdateSession(ctx context.Context, session domain.Session) error
	GetSessionByAccess(ctx context.Context, token string) (domain.Session, error)
	GetSessionByRefresh(ctx context.Context, token string) (domain.Session, error)
}
