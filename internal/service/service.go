package service

import (
	"context"
	"errors"

	"github.com/imranzahaev/warehouse/internal/domain"
	"github.com/rs/zerolog"
)

type User interface {
	SignUp(ctx context.Context, user domain.User, pwd string) (token, refreshToken string, err error)
	SignIn(ctx context.Context, login, password string) (user domain.User, token, refreshToken string, err error)
	Get(ctx context.Context, userID int) (domain.User, error)
	Update(ctx context.Context, user domain.User) error
	RefreshTokens(ctx context.Context, userID int, oldToken string) (token, refreshToken string, err error)
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
