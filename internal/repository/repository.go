package repository

import (
	"context"
	"database/sql"

	"github.com/begenov/real-time-forum/internal/domain"
)

type Authorization interface {
	Create(ctx context.Context, user domain.User) error
	GetByID(ctx context.Context, id int) (domain.User, error)
	GetByNickname(ctx context.Context, nickname string) (domain.User, error)
	GetByEmail(ctx context.Context, email string) (domain.User, error)
	UpdatePassword(ctx context.Context, password string, id int) error
}

type Session interface {
	Create(ctx context.Context, session domain.Session) error
	GetSessionByUserID(ctx context.Context, userID int) (domain.Session, error)
	Update(ctx context.Context, session domain.Session) error
	Delete(ctx context.Context, value string) error
}

type Repository struct {
	Authorization Authorization
	Session       Session
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthorization(db),
		Session:       NewSession(db),
	}
}
