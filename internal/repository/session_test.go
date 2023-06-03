package repository

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/begenov/real-time-forum/internal/domain"
)

func TestSessionRepo_Create(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx     context.Context
		session domain.Session
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
			r := &SessionRepo{
				db: tt.fields.db,
			}
			if err := r.Create(tt.args.ctx, tt.args.session); (err != nil) != tt.wantErr {
				t.Errorf("SessionRepo.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSessionRepo_GetSessionByUserID(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx    context.Context
		userID int
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
			r := &SessionRepo{
				db: tt.fields.db,
			}
			got, err := r.GetSessionByUserID(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("SessionRepo.GetSessionByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SessionRepo.GetSessionByUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSessionRepo_Update(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx     context.Context
		session domain.Session
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
			r := &SessionRepo{
				db: tt.fields.db,
			}
			if err := r.Update(tt.args.ctx, tt.args.session); (err != nil) != tt.wantErr {
				t.Errorf("SessionRepo.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
