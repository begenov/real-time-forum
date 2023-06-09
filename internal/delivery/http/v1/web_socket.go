package v1

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func (h *Handler) InitWebSoketRoute(router *mux.Router) {
	router.HandleFunc("/api/v1/web-socket", h.userIdentity(h.webSoketHandler))
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) webSoketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.handleError(w, http.StatusBadRequest, "Failed to upgrade connection to WebSocket:"+err.Error())
		return
	}

	h.activeConnections = append(h.activeConnections, conn)

	go h.handleWebSocketMessages(conn)
}

type Message struct {
	FromUserID int    `json:"from_user_id"`
	ToUserID   int    `json:"to_user_id"`
	Text       string `json:"text"`
}

func (h *Handler) handleWebSocketMessages(conn *websocket.Conn) {
	defer func() {
		h.removeConnection(conn)
		conn.Close()
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			h.log.Error("Failed to read WebSocket message: ", err)
			break
		}

		var message Message
		if err := json.Unmarshal(msg, &message); err != nil {
			if e, ok := err.(*json.SyntaxError); ok {
				h.log.Error("syntax error at byte offset", e.Offset)
				break
			}
			h.log.Error(err)
			break
		}

		h.broadcastMessage(message, conn)
		h.validateMessage(message)

	}
}

func (h *Handler) broadcastMessage(message Message, c *websocket.Conn) {
	for _, conn := range h.activeConnections {
		if conn == c {
			continue
		}
		err := conn.WriteJSON(message)
		if err != nil {
			h.log.Error("Failed to send WebSocket message: ", err)
		}
	}
}

func (h *Handler) validateMessage(message Message) {
	if message.FromUserID <= 0 {
		h.log.Error("Invalid FromUserID in the message")
		return
	}

	if message.ToUserID <= 0 {
		h.log.Error("Invalid ToUserID in the message")
		return
	}

	if len(message.Text) == 0 {
		h.log.Error("Empty Text field in the message")
		return
	}

	h.log.Info("Validated message: %+v", message)
}

func (h *Handler) removeConnection(conn *websocket.Conn) {
	for i, c := range h.activeConnections {
		if c == conn {
			h.activeConnections = append(h.activeConnections[:i], h.activeConnections[i+1:]...)
			break
		}
	}
}
