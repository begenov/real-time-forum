package service

import (
	"context"

	"github.com/begenov/real-time-forum/internal/domain"
	"github.com/begenov/real-time-forum/internal/repository"
)

type ChatService struct {
	repo repository.Chat
}

func NewChatService(repo repository.Chat) *ChatService {
	return &ChatService{
		repo: repo,
	}
}

func (s *ChatService) Create(ctx context.Context, msg domain.Message) error {
	if err := s.repo.Create(ctx, msg); err != nil {
		return err
	}

	return nil
}

func (s *ChatService) GetMessages() {
}

func (s *ChatService) ReadMessage() {
}

func (s *ChatService) GetChats() {
}
