package postgres

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/youngmikedev/warehouse/internal/domain"
)

func TestUsersRepo(t *testing.T) {
	usersCreate(t)
	usersUpdate(t)
	usersGet(t)
	sessionCreate(t)
	sessionUpdate(t)
	sessionGetByAccess(t)
}

func usersCreate(t *testing.T) {
	type args struct {
		user     domain.User
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
		errType error
	}{
		{
			name: "1.Valid",
			args: args{
				user: domain.User{
					Name:  "Tom",
					Email: "example@ecample.org",
				},
				password: "strongpass",
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "2.Empty email",
			args: args{
				user: domain.User{
					Name: "Tom",
				},
				password: "strongpass",
			},
			want:    0,
			wantErr: true,
			errType: &domain.ValidationError{
				Field: "email",
				Err:   errors.New("value is less than the required length"),
			},
		},
		{
			name: "3.Invalid duplicate",
			args: args{
				user: domain.User{
					Name:  "Tom",
					Email: "example@ecample.org",
				},
				password: "strongpass",
			},
			want:    0,
			wantErr: true,
			errType: domain.ErrUserAlreadyExists,
		},
		{
			name: "4.Invalid email",
			args: args{
				user: domain.User{
					Name:  "Tom",
					Email: "example",
				},
				password: "strongpass",
			},
			want:    0,
			wantErr: true,
			errType: &domain.ValidationError{
				Field: "email",
				Err:   errors.New("invalid email address"),
			},
		},
		{
			name: "5.Empty email and name",
			args: args{
				user:     domain.User{},
				password: "strongpass",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "6.Empty name",
			args: args{
				user: domain.User{
					Email: "example1@ecample.org",
				},
				password: "strongpass",
			},
			want:    0,
			wantErr: true,
			errType: &domain.ValidationError{
				Field: "name",
				Err:   errors.New("value is less than the required length"),
			},
		},
		{
			name: "7.Empty password",
			args: args{
				user: domain.User{
					Name:  "Tom",
					Email: "example1@ecample.org",
				},
				password: "",
			},
			want:    0,
			wantErr: true,
			errType: &domain.ValidationError{
				Field: "password",
				Err:   errors.New("value is less than the required length"),
			},
		},
		{
			name: "8.Valid",
			args: args{
				user: domain.User{
					Name:  "Ivan",
					Email: "another@ecample.org",
				},
				password: "strongpass",
			},
			want:    3,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run("Create "+tt.name, func(t *testing.T) {
			r := &UsersRepo{
				client: db,
			}
			got, err := r.Create(context.TODO(), tt.args.user, tt.args.password)
			switch {
			case (err != nil) != tt.wantErr:
				t.Errorf("UsersRepo.Create() error = %v, wantErr %v", err, tt.wantErr)
				return

			case err != nil && tt.errType != nil && !(err.Error() == tt.errType.Error()):
				t.Errorf("UsersRepo.Create() error = %v, want = %v", err, tt.errType)
				return
			}
			if got != tt.want {
				t.Errorf("UsersRepo.Create() = %v, want %v", got, tt.want)
			}
		})
	}

	// Update
}

func usersUpdate(t *testing.T) {
	type args struct {
		user     domain.User
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		errType error
	}{
		{
			name: "1.Valid",
			args: args{
				user: domain.User{
					ID:    1,
					Name:  "Tom_new",
					Email: "example_new@ecample.org",
				},
				password: "strongpass_new",
			},
			wantErr: false,
		},
		{
			name: "2.Invalid duplicate",
			args: args{
				user: domain.User{
					ID:    3,
					Name:  "Tom",
					Email: "example_new@ecample.org",
				},
				password: "strongpass",
			},
			wantErr: true,
			errType: domain.ErrUserAlreadyExists,
		},
		{
			name: "3.Invalid email",
			args: args{
				user: domain.User{
					ID:    1,
					Email: "example",
				},
			},
			wantErr: true,
			errType: &domain.ValidationError{
				Field: "email",
				Err:   errors.New("invalid email address"),
			},
		},
		{
			name: "4.Empty id",
			args: args{
				user: domain.User{
					Name: "Ivan",
				},
			},
			wantErr: true,
			errType: errors.New("empty user id"),
		},
		{
			name: "5.Invalid id",
			args: args{
				user: domain.User{
					ID:   1000,
					Name: "Ivan",
				},
			},
			wantErr: true,
			errType: domain.ErrUserNotFound,
		},
		{
			name: "6.Valid one field",
			args: args{
				user: domain.User{
					ID:   1,
					Name: "Jerry",
				},
				password: "strongpass_new",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run("Update "+tt.name, func(t *testing.T) {
			r := &UsersRepo{
				client: db,
			}
			err := r.Update(context.TODO(), tt.args.user, tt.args.password)
			switch {
			case (err != nil) != tt.wantErr:
				t.Errorf("UsersRepo.Update() error = %v, wantErr %v", err, tt.wantErr)
				return

			case err != nil && tt.errType != nil && !(err.Error() == tt.errType.Error()):
				t.Errorf("UsersRepo.Update() error = %v, want = %v", err, tt.errType)
				return
			}
		})
	}
}

func usersGet(t *testing.T) {
	tests := []struct {
		name    string
		userID  int
		want    domain.User
		wantPwd string
		wantErr bool
		errType error
	}{
		{
			name:   "1.Valid",
			userID: 1,
			want: domain.User{
				ID:    1,
				Name:  "Jerry",
				Email: "example_new@ecample.org",
			},
			wantPwd: "strongpass_new",
		},
		{
			name:    "2.Invalid id",
			userID:  1000,
			wantErr: true,
			errType: domain.ErrUserNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UsersRepo{
				client: db,
			}
			got, gotPwd, err := r.Get(context.TODO(), tt.userID)
			switch {
			case (err != nil) != tt.wantErr:
				t.Errorf("UsersRepo.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			case err != nil && tt.errType != nil && !(err.Error() == tt.errType.Error()):
				t.Errorf("UsersRepo.Get() error = %v, want = %v", err, tt.errType)
				return
			}
			if got.ID != tt.want.ID ||
				got.Email != tt.want.Email ||
				got.Name != tt.want.Name ||
				gotPwd != tt.wantPwd {
				t.Errorf("UsersRepo.Get() (usr = %v, pwd = %v), want (usr = %v, pwd = %v)", got, gotPwd, tt.want, tt.wantPwd)
			}
		})
	}
}

func usersGetByLogin(t *testing.T) {
	tests := []struct {
		name    string
		login   string
		want    domain.User
		wantPwd string
		wantErr bool
		errType error
	}{
		{
			name:  "1.Valid",
			login: "example_new@ecample.org",
			want: domain.User{
				ID:    1,
				Name:  "Jerry",
				Email: "example_new@ecample.org",
			},
			wantPwd: "strongpass_new",
		},
		{
			name:    "2.Invalid id",
			login:   "a@a.a",
			wantErr: true,
			errType: domain.ErrUserNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UsersRepo{
				client: db,
			}
			got, gotPwd, err := r.GetByLogin(context.TODO(), tt.login)
			switch {
			case (err != nil) != tt.wantErr:
				t.Errorf("UsersRepo.GetByLogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			case err != nil && tt.errType != nil && !(err.Error() == tt.errType.Error()):
				t.Errorf("UsersRepo.GetByLogin() error = %v, want = %v", err, tt.errType)
				return
			}
			if got.ID != tt.want.ID ||
				got.Email != tt.want.Email ||
				got.Name != tt.want.Name ||
				gotPwd != tt.wantPwd {
				t.Errorf("UsersRepo.GetByLogin() (usr = %v, pwd = %v), want (usr = %v, pwd = %v)", got, gotPwd, tt.want, tt.wantPwd)
			}
		})
	}
}

func sessionCreate(t *testing.T) {
	tests := []struct {
		name    string
		session domain.Session
		want    int
		wantErr bool
		errType error
	}{
		{
			name: "1.Valid",
			session: domain.Session{
				UserID:       1,
				AccessToken:  "access_token",
				RefreshToken: "refresh_token",
				ExpiresAt:    time.Hour,
			},
			want: 1,
		},
		{
			name: "2.Invalid user id",
			session: domain.Session{
				UserID:       1000,
				AccessToken:  "access_token",
				RefreshToken: "refresh_token",
			},
			wantErr: true,
			errType: domain.ErrUserNotFound,
		},
		{
			name: "3.Invalid refresh token",
			session: domain.Session{
				UserID:       1,
				AccessToken:  "access_token",
				RefreshToken: "",
				ExpiresAt:    time.Hour,
			},
			wantErr: true,
			errType: &domain.ValidationError{
				Field: "refresh_token",
				Err:   errors.New("value is less than the required length"),
			},
		},
		{
			name: "4.Empty access token",
			session: domain.Session{
				UserID:       1,
				AccessToken:  "",
				RefreshToken: "refresh_token",
				ExpiresAt:    time.Hour,
			},
			wantErr: true,
			errType: &domain.ValidationError{
				Field: "access_token",
				Err:   errors.New("value is less than the required length"),
			},
		},
	}
	for _, tt := range tests {
		t.Run("Session create "+tt.name, func(t *testing.T) {
			r := &UsersRepo{
				client: db,
			}
			got, err := r.CreateSession(context.TODO(), tt.session)
			switch {
			case (err != nil) != tt.wantErr:
				t.Errorf("UsersRepo.CreateSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			case err != nil && tt.errType != nil && !(err.Error() == tt.errType.Error()):
				t.Errorf("UsersRepo.CreateSession() error = %v, want = %v", err, tt.errType)
				return
			}
			if got != tt.want {
				t.Errorf("UsersRepo.CreateSession() = %v, want %v", got, tt.want)
			}
		})
	}
}

func sessionUpdate(t *testing.T) {
	tests := []struct {
		name    string
		session domain.Session
		wantErr bool
		errType error
	}{
		{
			name: "1.Valid",
			session: domain.Session{
				ID:           1,
				AccessToken:  "new_access_token",
				RefreshToken: "new_refresh_token",
				Disabled:     true,
			},
			wantErr: false,
		},
		{
			name: "2.Valid one field",
			session: domain.Session{
				ID:       1,
				Disabled: false,
			},
			wantErr: false,
		},
		{
			name: "3.Empty id",
			session: domain.Session{
				AccessToken:  "new_access_token",
				RefreshToken: "new_refresh_token",
				Disabled:     true,
			},
			wantErr: true,
			errType: errors.New("empty session id"),
		},
		{
			name: "4.Invalid id",
			session: domain.Session{
				ID:           1000,
				AccessToken:  "new_access_token",
				RefreshToken: "new_refresh_token",
				Disabled:     true,
			},
			wantErr: true,
			errType: domain.ErrSessionNotFound,
		},
	}
	for _, tt := range tests {
		t.Run("Update "+tt.name, func(t *testing.T) {
			r := &UsersRepo{
				client: db,
			}
			err := r.UpdateSession(context.TODO(), tt.session)
			switch {
			case (err != nil) != tt.wantErr:
				t.Errorf("UsersRepo.UpdateSession() error = %v, wantErr %v", err, tt.wantErr)
				return

			case err != nil && tt.errType != nil && !(err.Error() == tt.errType.Error()):
				t.Errorf("UsersRepo.UpdateSession() error = %v, want = %v", err, tt.errType)
				return
			}
		})
	}
}

func sessionGetByAccess(t *testing.T) {
	tests := []struct {
		name    string
		token   string
		want    domain.Session
		wantErr bool
		errType error
	}{
		{
			name:  "1.Valid",
			token: "new_access_token",
			want: domain.Session{
				ID:           1,
				UserID:       1,
				AccessToken:  "new_access_token",
				RefreshToken: "new_refresh_token",
				Disabled:     false,
			},
		},
		{
			name:    "2.Invalid token",
			token:   "access_token",
			wantErr: true,
			errType: domain.ErrSessionNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UsersRepo{
				client: db,
			}
			got, err := r.GetSessionByAccess(context.TODO(), tt.token)
			switch {
			case (err != nil) != tt.wantErr:
				t.Errorf("UsersRepo.GetByAccess() error = %v, wantErr %v", err, tt.wantErr)
				return
			case err != nil && tt.errType != nil && !(err.Error() == tt.errType.Error()):
				t.Errorf("UsersRepo.GetByAccess() error = %v, want = %v", err, tt.errType)
				return
			}
			if got.ID != tt.want.ID ||
				got.UserID != tt.want.UserID ||
				got.AccessToken != tt.want.AccessToken ||
				got.RefreshToken != tt.want.RefreshToken ||
				got.Disabled != tt.want.Disabled {
				t.Errorf("UsersRepo.GetByAccess() = %v, want %v", got, tt.want)
			}
		})
	}
}
