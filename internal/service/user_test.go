package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/begenov/real-time-forum/internal/config"
	"github.com/begenov/real-time-forum/internal/domain"
	"github.com/begenov/real-time-forum/internal/repository"
	"github.com/begenov/real-time-forum/pkg/auth"
	"github.com/begenov/real-time-forum/pkg/hash"
)

func TestUserService_SignUp(t *testing.T) {
	type fields struct {
		auth    repository.Authorization
		session repository.Session
		hash    hash.PasswordHasher
		cfg     config.Token
		manager auth.TokenManager
	}
	type args struct {
		ctx  context.Context
		user domain.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserService{
				auth:    tt.fields.auth,
				session: tt.fields.session,
				hash:    tt.fields.hash,
				cfg:     tt.fields.cfg,
				manager: tt.fields.manager,
			}
			if err := s.SignUp(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("UserService.SignUp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserService_SignIn(t *testing.T) {
	type fields struct {
		auth    repository.Authorization
		session repository.Session
		hash    hash.PasswordHasher
		cfg     config.Token
		manager auth.TokenManager
	}
	type args struct {
		ctx      context.Context
		email    string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Session
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserService{
				auth:    tt.fields.auth,
				session: tt.fields.session,
				hash:    tt.fields.hash,
				cfg:     tt.fields.cfg,
				manager: tt.fields.manager,
			}
			got, err := s.SignIn(tt.args.ctx, tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.SignIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.SignIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_UpdatePassword(t *testing.T) {
	type fields struct {
		auth    repository.Authorization
		session repository.Session
		hash    hash.PasswordHasher
		cfg     config.Token
		manager auth.TokenManager
	}
	type args struct {
		ctx      context.Context
		password string
		id       int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserService{
				auth:    tt.fields.auth,
				session: tt.fields.session,
				hash:    tt.fields.hash,
				cfg:     tt.fields.cfg,
				manager: tt.fields.manager,
			}
			if err := s.UpdatePassword(tt.args.ctx, tt.args.password, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UserService.UpdatePassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
