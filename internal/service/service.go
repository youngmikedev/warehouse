package service

import (
	"context"
	"errors"

	"github.com/imranzahaev/warehouse/internal/domain"
	"github.com/rs/zerolog"
)

type User interface {
	SignUp(ctx context.Context, user domain.User, password string) error
	SignIn(ctx context.Context, login, password string) (SignInResponse, error)
	Get(ctx context.Context, userID int) (domain.User, error)
	Update(ctx context.Context, user domain.User, password string) (err error)
	RefreshTokens(ctx context.Context, oldRefreshToken string) (accessToken, refreshToken string, err error)
	LogOut(ctx context.Context, accessToken string) error
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
