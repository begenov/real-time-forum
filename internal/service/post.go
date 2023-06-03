package service

import (
	"context"
	"errors"
	"time"

	"github.com/begenov/real-time-forum/internal/domain"
	"github.com/begenov/real-time-forum/internal/repository"
)

type PostService struct {
	repo     repository.Post
	category repository.Category
}

func NewPostService(repo repository.Post, category repository.Category) *PostService {
	return &PostService{
		repo:     repo,
		category: category,
	}
}

func (s *PostService) Create(ctx context.Context, post domain.Post) error {
	id, err := s.category.GetCategoryPostCategoryID(ctx, post.Category)
	if err != nil {
		return err
	}
	post.CategoryID = id

	err = s.repo.Create(ctx, post)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostService) Update(ctx context.Context, post domain.Post) error {
	// Retrieve the existing post from the repository
	existingPost, err := s.repo.GetPostByID(ctx, post.ID)
	if err != nil {
		return err
	}

	if post.Title != "" {
		existingPost.Title = post.Title
	}
	if post.Description != "" {
		existingPost.Description = post.Description
	}
	if len(post.Category) > 0 {
		id, err := s.category.GetCategoryPostCategoryID(ctx, post.Category)
		if err != nil {
			return err
		}

		existingPost.CategoryID = id
	}

	// Set the update timestamp
	existingPost.UpdateAt = time.Now()

	// Save the updated post using the repository
	err = s.repo.Update(ctx, existingPost)
	if err != nil {
		return err
	}

	// Perform any other operations or return the appropriate response
	// ...

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
