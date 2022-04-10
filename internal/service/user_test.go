package service

import (
	"context"
	"os"
	"testing"

	"github.com/rs/zerolog"
	authmock "github.com/youngmikedev/warehouse/internal/auth/mocks"
	"github.com/youngmikedev/warehouse/internal/domain"
	"github.com/youngmikedev/warehouse/internal/repository"
	repomocks "github.com/youngmikedev/warehouse/internal/repository/mocks"
)

func TestUserService_SignUp(t *testing.T) {
	ctx := context.TODO()
	type fields struct {
		tokenManager func() *authmock.TokenManager
		hashManager  func() *authmock.HashManager
		repo         func() *repomocks.User
	}
	type args struct {
		user     domain.User
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		errType error
	}{
		{
			name: "1.Valid",
			args: args{
				user: domain.User{
					Name:  "Tom",
					Email: "example@example.org",
				},
				password: "strongpassword",
			},
			fields: fields{
				repo: func() *repomocks.User {
					m := new(repomocks.User)
					m.On("Create", ctx, domain.User{Name: "Tom", Email: "example@example.org"}, "hashedpassword").Return(0, nil)
					return m
				},
				tokenManager: func() *authmock.TokenManager {
					m := new(authmock.TokenManager)
					return m
				},
				hashManager: func() *authmock.HashManager {
					m := new(authmock.HashManager)
					m.On("Hash", "strongpassword").Return("hashedpassword", nil)
					return m
				},
			},
			wantErr: false,
		},
		{
			name: "1.Invalid email",
			args: args{
				user: domain.User{
					Name:  "Tom",
					Email: "example@example.org",
				},
				password: "strongpassword",
			},
			fields: fields{
				repo: func() *repomocks.User {
					m := new(repomocks.User)
					m.On("Create", ctx, domain.User{Name: "Tom", Email: "example@example.org"}, "hashedpassword").Return(0, nil)
					return m
				},
				tokenManager: func() *authmock.TokenManager {
					m := new(authmock.TokenManager)
					return m
				},
				hashManager: func() *authmock.HashManager {
					m := new(authmock.HashManager)
					m.On("Hash", "strongpassword").Return("hashedpassword", nil)
					return m
				},
			},
			wantErr: false,
		},
	}
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm := tt.fields.tokenManager()
			hm := tt.fields.hashManager()
			um := tt.fields.repo()
			repo := repository.Repositories{User: um}

			s := &UserService{
				tokenManager: tm,
				hashManager:  hm,
				repo:         &repo,
				log:          &logger,
			}
			if err := s.SignUp(ctx, tt.args.user, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("UserService.SignUp() error = %v, wantErr %v", err, tt.wantErr)
			}
			tm.AssertExpectations(t)
			hm.AssertExpectations(t)
			um.AssertExpectations(t)
		})
	}
}
