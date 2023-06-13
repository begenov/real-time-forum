package ws

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/begenov/real-time-forum/internal/domain"
	"github.com/begenov/real-time-forum/internal/service"
	"github.com/gorilla/websocket"
)

type Handler struct {
	clients  map[int]*client
	upgrader websocket.Upgrader
	service  *service.Service
	wsEvent  chan *domain.WSEvent
	users    map[int]domain.Users
	chanConn chan struct{}
}

func NewHandler(service *service.Service, wsEvent chan *domain.WSEvent) *Handler {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  4096,
		WriteBufferSize: 4096,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	return &Handler{
		upgrader: upgrader,
		service:  service,
		clients:  make(map[int]*client),
		wsEvent:  wsEvent,
		chanConn: make(chan struct{}),
	}
}

func (h *Handler) ServeWS(w http.ResponseWriter, r *http.Request) {
	connect, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error connection %v", err)
		return
	}

	connection := &conn{conn: connect}

	cookie, err := r.Cookie("session")
	if err != nil {
		log.Printf("error cookie ws %v", err)
		return
	}

	user, err := h.service.User.GetUserByToken(r.Context(), cookie.Value)
	if err != nil {
		log.Printf("error get user by token %v", err)
		return
	}

	c, ok := h.clients[user.Id]
	if !ok {
		c = &client{User: user}
		h.clients[user.Id] = c
	}

	c.conns = append(c.conns, connection)
	h.chanConn <- struct{}{}
	go h.handleClientMessages(user.Id, connection)
	go h.allUsers(user.Id, c.conns)
}

func (h *Handler) handleClientMessages(id int, connection *conn) {

	for {
		event, err := connection.getEvent()
		if err != nil {
			log.Println(err.Error())
			return
		}

		switch event.Type {
		case "message":
			err = h.newMessage(connection.clientID, &event)
		}

		if err != nil {
			log.Println(err)
			return
		}
	}

}

type messageInput struct {
	RecipientID int    `json:"to_user_id"`
	Message     string `json:"message"`
}

func (h *Handler) newMessage(clientID int, event *domain.WSEvent) error {
	var inp messageInput

	body, err := json.Marshal(event.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &inp); err != nil {
		return err
	}

	if err := h.service.Chat.Create(context.Background(), domain.Message{
		FromUserID: clientID,
		ToUserID:   event.RecipientID,
		Text:       inp.Message,
	}); err != nil {
		return err
	}

	return nil
}
func (h *Handler) allUsers(userID int, cons []*conn) {

	for {
		select {
		case <-h.chanConn:
			users, err := h.service.User.AllUsers(context.Background(), userID)
			if err != nil {
				log.Println(err.Error())
				return
			}
			for _, conn := range cons {
				conn.conn.WriteJSON(domain.WSEvent{
					Type: "online_users",
					Body: users,
				})
			}
		}

	}
}
