package domain

import "time"

type Post struct {
	ID          int
	UserID      int      `json:"userid"`
	Author      string   `json:"author"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Category    []string `json:"category"`
	CategoryID  []int
	CreateAt    time.Time
	UpdateAt    time.Time
}

type Category struct {
	Id       int
	Category string
}
