package domain

import "time"

type Comment struct {
	Id       int
	PostID   int
	UserID   int
	Text     string
	CreateAt time.Time
	UpdateAt time.Time
}
