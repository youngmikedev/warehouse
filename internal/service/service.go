package service

import (
	"context"
	"errors"

	"github.com/rs/zerolog"
	"github.com/youngmikedev/warehouse/internal/auth"
	"github.com/youngmikedev/warehouse/internal/domain"
	"github.com/youngmikedev/warehouse/internal/repository"
)

type Services struct {
	User
	Product
}

type User interface {
	SignUp(ctx context.Context, user domain.User, password string) error
	SignIn(ctx context.Context, login, password string) (SignInResponse, error)
	Get(ctx context.Context, userID int) (domain.User, error)
	Update(ctx context.Context, user domain.User, password string) (err error)
	CheckAccessToken(ctx context.Context, token string) (sesID, userID int, err error)
	RefreshTokens(ctx context.Context, oldRefreshToken string) (accessToken, refreshToken string, err error)
	LogOut(ctx context.Context, accessToken string) error
}

type Product interface {
	Create(ctx context.Context, uid int, product domain.Product) (id int, err error)
	Update(ctx context.Context, uid int, product domain.Product) (err error)
	Get(ctx context.Context, uid int, id int) (domain.Product, error)
	GetManyByFilter(ctx context.Context, filter domain.GetManyProductsFilter) (domain.GetManyProductsResponse, error)
}

func NewServices(
	repos *repository.Repositories,
	cache cache,
	tokenManager auth.TokenManager,
	hasher auth.HashManager,
	logger *zerolog.Logger,
) *Services {
	return &Services{
		User:    NewUserService(repos, cache, tokenManager, hasher, logger),
		Product: NewProductService(repos, logger),
	}
}

func checkAppError(l *zerolog.Logger, err error, fname string) error {
	if !errors.As(err, &domain.AppError{}) {
		l.Error().
			Err(err).
			Str("func", fname).
			Msg("internal error")
		return domain.ErrInternal
	}

	return err
}
