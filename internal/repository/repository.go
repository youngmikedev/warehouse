package repository

import (
	"context"

	"github.com/youngmikedev/warehouse/internal/domain"
	"github.com/youngmikedev/warehouse/internal/repository/postgres"
	"github.com/youngmikedev/warehouse/internal/repository/postgres/ent"
)

type Repositories struct {
	User
	Product
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

type Product interface {
	Create(ctx context.Context, uid int, product domain.Product) (int, error)
	Update(ctx context.Context, uid int, product domain.Product) error
	Get(ctx context.Context, uid, id int) (domain.Product, error)
	GetManyByFilter(ctx context.Context, filter domain.GetManyProductsFilter) (domain.GetManyProductsResponse, error)
}

func NewPostgresRepositories(client *ent.Client) *Repositories {
	return &Repositories{
		User:    postgres.NewUsersRepo(client),
		Product: postgres.NewProductRepo(client),
	}
}
