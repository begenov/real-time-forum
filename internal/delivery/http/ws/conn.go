package ws

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/begenov/real-time-forum/internal/domain"
	"github.com/gorilla/websocket"
)

type conn struct {
	clientID int
	conn     *websocket.Conn
	mu       sync.Mutex
}

func (c *conn) getEvent() (domain.WSEvent, error) {
	var event domain.WSEvent

	_, msg, err := c.conn.ReadMessage()
	if err != nil {
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			log.Printf("error get event %v", err)
		}
		return event, err
	}

	err = json.Unmarshal(msg, &event)

	return event, err
}
