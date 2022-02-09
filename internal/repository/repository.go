package repository

import (
	"context"

	"github.com/imranzahaev/warehouse/internal/domain"
	"github.com/imranzahaev/warehouse/internal/repository/postgres"
	"github.com/imranzahaev/warehouse/internal/repository/postgres/ent"
)

type Repositories struct {
	User
}

type User interface {
	Create(ctx context.Context, user domain.User, password string) (int, error)
	Update(ctx context.Context, user domain.User, password string) error
	Get(ctx context.Context, id int) (u domain.User, password string, err error)
	GetByLogin(ctx context.Context, login string) (u domain.User, password string, err error)
	CreateSession(ctx context.Context, session domain.Session) (int, error)
	UpdateSession(ctx context.Context, session domain.Session) error
	GetSessionByAccess(ctx context.Context, token string) (domain.Session, error)
	GetSessionByRefresh(ctx context.Context, token string) (domain.Session, error)
}

func NewPostgresRepositories(client *ent.Client) *Repositories {
	return &Repositories{
		User: postgres.NewUsersRepo(client),
	}
}
