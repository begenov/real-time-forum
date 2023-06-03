package repository

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/begenov/real-time-forum/internal/domain"
)

func TestAuthorizationRepo_Create(t *testing.T) {
	type fields struct {
		db *sql.DB
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
			r := &AuthorizationRepo{
				db: tt.fields.db,
			}
			if err := r.Create(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("AuthorizationRepo.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAuthorizationRepo_GetByID(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AuthorizationRepo{
				db: tt.fields.db,
			}
			got, err := r.GetByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthorizationRepo.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthorizationRepo.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthorizationRepo_GetByNickname(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx      context.Context
		nickname string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AuthorizationRepo{
				db: tt.fields.db,
			}
			got, err := r.GetByNickname(tt.args.ctx, tt.args.nickname)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthorizationRepo.GetByNickname() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthorizationRepo.GetByNickname() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthorizationRepo_GetByEmail(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx   context.Context
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AuthorizationRepo{
				db: tt.fields.db,
			}
			got, err := r.GetByEmail(tt.args.ctx, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthorizationRepo.GetByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthorizationRepo.GetByEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthorizationRepo_UpdatePassword(t *testing.T) {
	type fields struct {
		db *sql.DB
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
			r := &AuthorizationRepo{
				db: tt.fields.db,
			}
			if err := r.UpdatePassword(tt.args.ctx, tt.args.password, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("AuthorizationRepo.UpdatePassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
