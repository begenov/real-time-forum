package domain

import "time"

type Post struct {
	ID          int
	UserID      int      `json:"userid"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Category    []string `json:"category"`
	CreateAt    time.Time
	UpdateAt    time.Time
}
