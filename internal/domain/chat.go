package domain

import "time"

type Message struct {
	ID         int       `json:"id"`
	FromUserID int       `json:"from_user_id"`
	ToUserID   int       `json:"to_user_id"`
	Text       string    `json:"text"`
	CreateAt   time.Time `json:"create_at"`
}
