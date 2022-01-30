package service

type UserService struct {
	// tokenManager auth.TokenManager
	// hasher       auth.HashManager
	// repo         *repository.Repositories
}

// func NewUserService(repo *repository.Repositories, tokenManager auth.TokenManager, hasher auth.HashManager) *UserService {
// 	return &UserService{
// 		tokenManager: tokenManager,
// 		hasher:       hasher,
// 		repo:         repo,
// 	}
// }

// func (s *UserService) SignUp(ctx context.Context, user *domain.User) (token, refreshToken string, err error) {
// 	if err = s.repo.User.CheckUserExist(ctx, user.Email); err != nil {
// 		if err == errors.ErrUserAlreadyExists {
// 			log.Error().
// 				Err(err).
// 				Str("email", user.Email).
// 				Send()
// 			return "", "", errors.ErrUserAlreadyExists
// 		}
// 		log.Error().
// 			Err(err).
// 			Msg("failed check user exist")
// 		return "", "", errors.ErrInternal
// 	}

// 	user.CreatedAt = time.Now()
// 	user.Password, err = s.hasher.Hash(user.Password)
// 	if err != nil {
// 		log.Error().
// 			Err(err).
// 			Msg("failed hash password")
// 		return "", "", errors.ErrInternal
// 	}

// 	if err = s.repo.User.Create(ctx, user); err != nil {
// 		log.Error().
// 			Err(err).
// 			Msg("failed create user")
// 		return "", "", errors.ErrInternal
// 	}

// 	rt := s.tokenManager.NewRefreshToken()
// 	t, err := s.tokenManager.NewAccessToken(user.ID.String(), s.tokenManager.GetDefaultTokenExpiresAt())
// 	if err != nil {
// 		log.Error().
// 			Err(err).
// 			Msg("failed create access token")
// 		return "", "", errors.ErrInternal
// 	}

// 	if err = s.repo.User.SetSession(ctx, &domain.Session{
// 		UserID:       user.ID,
// 		RefreshToken: rt,
// 		ExpiresAt:    s.tokenManager.GetDefaultRefreshTokenExpiresAt(),
// 		UpdatedAt:    time.Now(),
// 		CreatedAt:    time.Now(),
// 	}); err != nil {
// 		log.Error().
// 			Err(err).
// 			Msg("failed set session")
// 		return "", "", errors.ErrInternal
// 	}

// 	return t, rt, nil
// }
