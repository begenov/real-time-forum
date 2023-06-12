package repository

import (
	"context"
	"database/sql"

	"github.com/begenov/real-time-forum/internal/domain"
)

type ChatRepo struct {
	db *sql.DB
}

func NewChatRepo(db *sql.DB) *ChatRepo {
	return &ChatRepo{
		db: db,
	}
}

func (r *ChatRepo) Create(ctx context.Context, msg domain.Message) error {
	stmt := `INSERT INTO messages(from_user_id, to_user_id, message, create_at) VALUES ($1, $2, $3, $4)`
	_, err := r.db.ExecContext(ctx, stmt, msg.FromUserID, msg.ToUserID, msg.Text, msg.CreateAt)

	if err != nil {
		return err
	}
	return nil
}

func (r *ChatRepo) GetMessages(ctx context.Context, from_user_id, to_user_id int, lastMessageID, limit int) ([]domain.Message, error) {
	return nil, nil
}

func (r *ChatRepo) ReadMessage(ctx context.Context, to_user_id, messageID int) (domain.Message, error) {
	return domain.Message{}, nil
}

func (r *ChatRepo) GetChats() {
}
