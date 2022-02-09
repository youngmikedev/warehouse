package service

import (
	"context"
	"errors"

	"github.com/imranzahaev/warehouse/internal/auth"
	"github.com/imranzahaev/warehouse/internal/domain"
	"github.com/imranzahaev/warehouse/internal/repository"
	"github.com/rs/zerolog"
)

type Services struct {
	User
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

func NewServices(
	repos *repository.Repositories,
	tokenManager auth.TokenManager,
	hasher auth.HashManager,
	logger *zerolog.Logger,
) *Services {
	return &Services{
		User: NewUserService(repos, tokenManager, hasher, logger),
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
