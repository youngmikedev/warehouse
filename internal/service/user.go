package service

import (
	"context"
	"errors"

	"github.com/rs/zerolog"
	"github.com/youngmikedev/warehouse/internal/auth"
	"github.com/youngmikedev/warehouse/internal/domain"
	"github.com/youngmikedev/warehouse/internal/repository"
)

type cache interface {
	Set(key string, value ShortSession)
	Get(key string) (ShortSession, bool)
	Delete(key string)
}
type UserService struct {
	cache        cache
	tokenManager auth.TokenManager
	hashManager  auth.HashManager
	repo         *repository.Repositories
	log          *zerolog.Logger
}

func NewUserService(repo *repository.Repositories, cache cache, tokenManager auth.TokenManager, hashManager auth.HashManager, logger *zerolog.Logger) *UserService {
	sl := logger.With().Str("service", "user").Logger()
	return &UserService{
		cache:        cache,
		tokenManager: tokenManager,
		hashManager:  hashManager,
		repo:         repo,
		log:          &sl,
	}
}

func (s *UserService) SignUp(ctx context.Context, user domain.User, password string) error {
	pwd, err := s.hashManager.Hash(password)
	if err != nil {
		s.log.Error().
			Err(err).
			Str("func", "SignUp.Hash").
			Msg("failed hash password")
		return domain.ErrInternal
	}
	_, err = s.repo.User.Create(ctx, user, pwd)
	if err != nil {
		return checkAppError(
			"SignUp.Create",
			err,
			s.log.Error().
				Str("email", user.Email),
		)
	}

	return nil
}

type SignInResponse struct {
	User         domain.User
	AccessToken  string
	RefreshToken string
}

func (s *UserService) SignIn(ctx context.Context, login, password string) (SignInResponse, error) {
	u, hpwd, err := s.repo.User.GetByLogin(ctx, login)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			s.log.Info().
				Str("login", login).
				Str("invalid", "email").
				Msg("failed login")
			return SignInResponse{}, domain.ErrInvalidLoginOrPassword
		}
		return SignInResponse{}, checkAppError(
			"SignIn.GetByLogin",
			err,
			s.log.Error().
				Str("email", login),
		)
	}

	if !s.hashManager.Validate(hpwd, password) {
		s.log.Info().
			Str("login", login).
			Str("invalid", "password").
			Msg("failed login")
		return SignInResponse{}, domain.ErrInvalidLoginOrPassword
	}

	at := s.tokenManager.NewAccessToken()
	rt := s.tokenManager.NewRefreshToken()

	_, err = s.repo.User.CreateSession(ctx, domain.Session{
		UserID:       u.ID,
		AccessToken:  at,
		RefreshToken: rt,
	})
	if err != nil {
		return SignInResponse{}, checkAppError(
			"SignIn.CreateSession",
			err,
			s.log.Error().
				Str("email", login).
				Int("uid", u.ID),
		)
	}

	return SignInResponse{AccessToken: at, RefreshToken: rt, User: u}, nil
}

func (s *UserService) Get(ctx context.Context, userID int) (domain.User, error) {
	u, _, err := s.repo.User.Get(ctx, userID)
	if err != nil {
		return domain.User{}, checkAppError(
			"Get.Get",
			err,
			s.log.Error().
				Int("uid", userID),
		)
	}

	return u, nil
}

func (s *UserService) Update(ctx context.Context, user domain.User, password string) (err error) {
	if password != "" {
		password, err = s.hashManager.Hash(password)
		if err != nil {
			s.log.Error().
				Err(err).
				Str("func", "Update.Hash").
				Msg("failed hash password")
			return domain.ErrInternal
		}
	}

	if err = s.repo.User.Update(ctx, user, password); err != nil {
		return checkAppError(
			"Update.Update",
			err,
			s.log.Error().
				Int("uid", user.ID),
		)
	}

	return nil
}

func (s *UserService) CheckAccessToken(ctx context.Context, token string) (sesID, userID int, err error) {
	if ses, ok := s.cache.Get(token); ok {
		if s.tokenManager.ValidateAccessToken(ses.UpdatedAt) {
			return ses.SessionID, ses.UserID, nil
		}
	}

	ses, err := s.repo.GetSessionByAccess(ctx, token)
	if err != nil {
		return 0, 0, checkAppError(
			"CheckAccessToken.GetSessionByAccess",
			err,
			s.log.Error(),
		)
	}

	if !s.tokenManager.ValidateAccessToken(ses.UpdatedAt) {
		ses.Disabled = true
		if err = s.repo.User.UpdateSession(ctx, ses); err != nil {
			return 0, 0, checkAppError(
				"CheckAccessToken.UpdateSession",
				err,
				s.log.Error().
					Int("uid", ses.UserID).
					Int("session id", ses.ID),
			)
		}
		return 0, 0, domain.ErrTokenExpired
	}

	s.cache.Set(token, ShortSession{
		UserID:      ses.UserID,
		SessionID:   ses.ID,
		AccessToken: token,
		UpdatedAt:   ses.UpdatedAt,
	})

	return ses.ID, ses.UserID, err
}

func (s *UserService) RefreshTokens(ctx context.Context, oldRefreshToken string) (accessToken, refreshToken string, err error) {
	ses, err := s.repo.GetSessionByRefresh(ctx, oldRefreshToken)
	if err != nil {
		return "", "", checkAppError(
			"RefreshTokens.GetSessionByRefresh",
			err,
			s.log.Error(),
		)
	}

	if s.tokenManager.ValidateRefreshToken(ses.UpdatedAt) {
		ses.Disabled = true
		if err = s.repo.User.UpdateSession(ctx, ses); err != nil {
			return "", "", checkAppError(
				"RefreshTokens.DisableSession",
				err,
				s.log.Error().
					Int("uid", ses.UserID).
					Int("session id", ses.ID),
			)
		}
		return "", "", domain.ErrTokenExpired
	}

	ses.AccessToken = s.tokenManager.NewAccessToken()
	ses.RefreshToken = s.tokenManager.NewRefreshToken()

	if err = s.repo.User.UpdateSession(ctx, ses); err != nil {
		return "", "", checkAppError(
			"RefreshTokens.UpdateSession",
			err,
			s.log.Error().
				Int("uid", ses.UserID).
				Int("session id", ses.ID),
		)
	}

	return ses.AccessToken, ses.RefreshToken, nil
}

func (s *UserService) LogOut(ctx context.Context, accessToken string) error {
	ses, err := s.repo.GetSessionByAccess(ctx, accessToken)
	if err != nil {
		return checkAppError(
			"LogOut.GetSessionByAccess",
			err,
			s.log.Error(),
		)
	}

	ses.Disabled = true
	if err = s.repo.User.UpdateSession(ctx, ses); err != nil {
		return checkAppError(
			"LogOut.UpdateSession",
			err,
			s.log.Error().
				Int("uid", ses.UserID).
				Int("session id", ses.ID),
		)
	}

	return nil
}
