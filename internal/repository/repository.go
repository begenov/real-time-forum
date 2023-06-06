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
	AllUsers(ctx context.Context) ([]domain.Users, error)
}

type Session interface {
	Create(ctx context.Context, session domain.Session) error
	GetSessionByUserID(ctx context.Context, userID int) (domain.Session, error)
	Update(ctx context.Context, session domain.Session) error
	Delete(ctx context.Context, value string) error
	GetUserIDByToken(ctx context.Context, value string) (domain.Session, error)
}

type Post interface {
	Create(ctx context.Context, post domain.Post) error
	Update(ctx context.Context, post domain.Post) error
	GetPostByID(ctx context.Context, id int) (domain.Post, error)
	GetAllPosts(ctx context.Context) ([]domain.Post, error)
	DeletePost(ctx context.Context, id int) error
}

type Category interface {
	CreateCategory(ctx context.Context, category string) error
	GetAllCategories(ctx context.Context) ([]domain.Category, error)
	DeleteCategory(ctx context.Context, category string) error
	GetCategoryPostCategoryID(ctx context.Context, categories []string) ([]int, error)
}

type Comment interface {
	Create(ctx context.Context, comment domain.Comment) error
	Update(ctx context.Context, comment domain.Comment) error
	Delete(ctx context.Context, id int) error
	GetAllComment(ctx context.Context) ([]domain.Comment, error)
	GetCommentByID(ctx context.Context, id int) (domain.Comment, error)
}

type Repository struct {
	Authorization Authorization
	Session       Session
	Post          Post
	Category      Category
	Comment       Comment
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthorization(db),
		Session:       NewSession(db),
		Post:          NewPostRepo(db),
		Category:      NewCategoryRepo(db),
		Comment:       NewCommentRepo(db),
	}
}
