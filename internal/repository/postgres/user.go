package postgres

import (
	"context"
	"errors"
	"strings"

	"github.com/imranzahaev/warehouse/internal/domain"
	"github.com/imranzahaev/warehouse/internal/repository/postgres/ent"
	"github.com/imranzahaev/warehouse/internal/repository/postgres/ent/session"
	entuser "github.com/imranzahaev/warehouse/internal/repository/postgres/ent/user"
)

type UsersRepo struct {
	client *ent.Client
}

func NewUsersRepo(client *ent.Client) *UsersRepo {
	return &UsersRepo{
		client: client,
	}
}

// Create user with hashed password
func (r *UsersRepo) Create(ctx context.Context, user domain.User, password string) (int, error) {
	res, err := r.client.User.
		Create().
		SetEmail(user.Email).
		SetName(user.Name).
		SetPassword(password).
		Save(ctx)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "duplicate key value violates unique constraint"):
			return 0, domain.ErrUserAlreadyExists

		case ent.IsValidationError(err):
			e := err.(*ent.ValidationError)
			return 0, domain.NewValidationError(e.Name, errors.Unwrap(e.Unwrap()))
		}

		return 0, err
	}

	return res.ID, nil
}

// Update user by id
func (r *UsersRepo) Update(ctx context.Context, user domain.User, password string) error {
	if user.ID == 0 {
		return domain.AppError{Err: errors.New("empty user id")}
	}

	u := r.client.User.UpdateOneID(user.ID)
	if password != "" {
		u = u.SetPassword(password)
	}
	if user.Email != "" {
		u = u.SetEmail(user.Email)
	}
	if user.Name != "" {
		u = u.SetName(user.Name)
	}

	_, err := u.Save(ctx)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "duplicate key value violates unique constraint"):
			return domain.ErrUserAlreadyExists

		case ent.IsValidationError(err):
			e := err.(*ent.ValidationError)
			return domain.NewValidationError(e.Name, errors.Unwrap(e.Unwrap()))
		case ent.IsNotFound(err):
			return domain.ErrUserNotFound
		}

		return err
	}

	return nil
}

func (r *UsersRepo) Get(ctx context.Context, id int) (domain.User, error) {
	u, err := r.client.User.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return domain.User{}, domain.ErrUserNotFound
		}
		return domain.User{}, err
	}

	return convertUserToDomain(u), nil
}

// GetByLogin expects the login will be email
func (r *UsersRepo) GetByLogin(ctx context.Context, login string) (domain.User, error) {
	u, err := r.client.User.Query().
		Where(entuser.Email(login)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return domain.User{}, domain.ErrUserNotFound
		}
		return domain.User{}, err
	}

	return convertUserToDomain(u), nil
}

func (r *UsersRepo) CreateSession(ctx context.Context, session domain.Session) (int, error) {
	s, err := r.client.Session.
		Create().
		SetAccessToken(session.AccessToken).
		SetRefreshToken(session.RefreshToken).
		SetOwnerID(session.UserID).
		Save(ctx)
	if err != nil {
		switch {
		case ent.IsValidationError(err):
			e := err.(*ent.ValidationError)
			return 0, domain.NewValidationError(e.Name, errors.Unwrap(e.Unwrap()))
		case ent.IsConstraintError(err) && strings.Contains(err.Error(), "violates foreign key constraint"):
			return 0, domain.ErrUserNotFound
		}
		return 0, err
	}
	return s.ID, nil
}

func (r *UsersRepo) UpdateSession(ctx context.Context, session domain.Session) error {
	if session.ID == 0 {
		return domain.AppError{Err: errors.New("empty session id")}
	}

	u := r.client.Session.UpdateOneID(session.ID)
	if session.AccessToken != "" {
		u = u.SetAccessToken(session.AccessToken)
	}
	if session.RefreshToken != "" {
		u = u.SetRefreshToken(session.RefreshToken)
	}
	u = u.SetDisabled(session.Disabled)

	_, err := u.Save(ctx)
	if err != nil {
		switch {
		case ent.IsValidationError(err):
			e := err.(*ent.ValidationError)
			return domain.NewValidationError(e.Name, errors.Unwrap(e.Unwrap()))
		case ent.IsNotFound(err):
			return domain.ErrSessionNotFound
		}
		return err
	}

	return nil
}

func (r *UsersRepo) GetSessionByAccess(ctx context.Context, token string) (domain.Session, error) {
	s, err := r.client.Session.
		Query().
		Where(session.AccessToken(token)).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return domain.Session{}, domain.ErrSessionNotFound
		}
		return domain.Session{}, err
	}

	user, err := s.QueryOwner().Only(ctx)
	if err != nil {
		return domain.Session{}, err
	}

	return convertSessionToDomain(user.ID, s), nil
}

func (r *UsersRepo) GetSessionByRefresh(ctx context.Context, token string) (domain.Session, error) {
	s, err := r.client.Session.
		Query().
		Where(session.RefreshToken(token)).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return domain.Session{}, domain.ErrSessionNotFound
		}
		return domain.Session{}, err
	}

	user, err := s.QueryOwner().Only(ctx)
	if err != nil {
		return domain.Session{}, err
	}

	return convertSessionToDomain(user.ID, s), nil
}

func convertUserToDomain(user *ent.User) domain.User {
	return domain.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}

func convertSessionToDomain(userID int, session *ent.Session) domain.Session {
	return domain.Session{
		ID:           session.ID,
		UserID:       userID,
		AccessToken:  session.AccessToken,
		RefreshToken: session.RefreshToken,
		CreatedAt:    session.CreatedAt,
	}
}
