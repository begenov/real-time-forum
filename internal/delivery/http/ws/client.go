package ws

import (
	"sync"

	"github.com/begenov/real-time-forum/internal/domain"
)

type client struct {
	conns []*conn
	mu    sync.Mutex
	domain.User
}
