package service

import (
	"context"
	"errors"
	"time"

	"github.com/begenov/real-time-forum/internal/domain"
	"github.com/begenov/real-time-forum/internal/repository"
)

type PostService struct {
	repo repository.Post
}

func NewPostService(repo repository.Post) *PostService {
	return &PostService{
		repo: repo,
	}
}

func (s *PostService) Create(ctx context.Context, post domain.Post) error {
	err := s.repo.Create(ctx, post)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostService) Update(ctx context.Context, post domain.Post) error {
	p, err := s.repo.GetPostByID(ctx, post.ID)
	if err != nil {
		return err
	}

	if p.UserID != post.UserID {
		return errors.New("no error")
	}

	if post.Title == "" {
		post.Title = p.Title
	}

	if post.Description == "" {
		post.Description = p.Description
	}

	if len(post.Category) == 0 {
		post.Category = p.Category
	}

	post.UpdateAt = time.Now()

	err = s.repo.Update(ctx, post)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostService) GetPostById(ctx context.Context, id int) (domain.Post, error) {
	return s.repo.GetPostByID(ctx, id)
}

func (s *PostService) GetAllPosts(ctx context.Context) ([]domain.Post, error) {
	return s.repo.GetAllPosts(ctx)
}

func (s *PostService) Delete(ctx context.Context, id int, userID int) error {
	p, err := s.repo.GetPostByID(ctx, id)
	if err != nil {
		return err
	}

	if p.UserID != userID {
		return errors.New("no error")
	}

	return s.repo.DeletePost(ctx, id)
}
