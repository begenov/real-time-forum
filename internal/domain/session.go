package domain

import "time"

type Session struct {
	Id        int
	UserID    int
	Token     string
	ExpiresAt time.Time
}
