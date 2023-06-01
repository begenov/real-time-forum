package service

import (
	"context"

	"github.com/begenov/real-time-forum/internal/config"
	"github.com/begenov/real-time-forum/internal/domain"
	"github.com/begenov/real-time-forum/internal/repository"
	"github.com/begenov/real-time-forum/pkg/auth"
	"github.com/begenov/real-time-forum/pkg/hash"
)

type User interface {
	SignUp(ctx context.Context, user domain.User) error
	SignIn(ctx context.Context, email string, password string) (domain.Session, error)
	UpdatePassword(ctx context.Context, password string, id int) error
	GetUserByID(ctx context.Context, id int) (domain.User, error)
	DeleteSession(ctx context.Context, value string) error
}

type Post interface {
	Create(ctx context.Context, post domain.Post) error
	Update(ctx context.Context, post domain.Post) error
	GetPostById(ctx context.Context, id int) (domain.Post, error)
	GetAllPosts(ctx context.Context) ([]domain.Post, error)
	Delete(ctx context.Context, id int) error
}

type Service struct {
	User User
	Post Post
}

func NewService(repo *repository.Repository, hash hash.PasswordHasher, manager auth.TokenManager, cfg *config.Config) *Service {
	return &Service{
		User: NewUserService(repo.Authorization, repo.Session, hash, manager, cfg.Token),
		Post: NewPostService(repo.Post),
	}
}
