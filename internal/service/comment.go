package service

import (
	"context"
	"time"

	"github.com/begenov/real-time-forum/internal/domain"
	"github.com/begenov/real-time-forum/internal/repository"
)

type CommentService struct {
	repo repository.Comment
}

func NewCommentService(repo repository.Comment) *CommentService {
	return &CommentService{
		repo: repo,
	}
}

func (s *CommentService) Create(ctx context.Context, comment domain.Comment) error {
	return s.repo.Create(ctx, comment)
}

func (s *CommentService) Update(ctx context.Context, comment domain.Comment) error {
	c, err := s.repo.GetCommentByID(ctx, comment.Id)
	if err != nil {
		return err
	}

	if comment.Text == "" {
		comment.Text = c.Text
	}

	comment.UpdateAt = time.Now()

	err = s.repo.Update(ctx, comment)
	return err
}

func (s *CommentService) GetCommentById(ctx context.Context, id int) (domain.Comment, error) {
	return s.repo.GetCommentByID(ctx, id)
}

func (s *CommentService) GetAllComment(ctx context.Context) ([]domain.Comment, error) {
	return s.repo.GetAllComment(ctx)
}

func (s *CommentService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
