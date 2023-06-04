package ws

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type Handler struct {
	hub *Hub
}

func NewHandler(h *Hub) *Handler {
	return &Handler{
		hub: h,
	}
}

type CreateRoomReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	// var req CreateRoomReq
	// if err := c.ShouldBindJSON(&req); err != nil {

	// 	return
	// }

	// h.hub.Rooms[req.ID] = &Room{
	// 	ID:      req.ID,
	// 	Name:    req.Name,
	// 	Clients: make(map[string]*Client),
	// }

}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) JoinRoom(w http.ResponseWriter, r *http.Request) {
	// conn, err := upgrader.Upgrade(w, r, nil)
	// if err != nil {
	// 	// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// roomID := c.Param("roomId")
	// clientID := c.Query("userId")
	// username := c.Query("username")

	// cl := &Client{
	// 	Conn:     conn,
	// 	Message:  make(chan *Message, 10),
	// 	ID:       clientID,
	// 	RoomID:   roomID,
	// 	Username: username,
	// }

	// m := &Message{
	// 	Content:  "A new user has joined the room",
	// 	RoomID:   roomID,
	// 	Username: username,
	// }

	// h.hub.Register <- cl
	// h.hub.Broadcast <- m

	// go cl.WriteMessage()
	// cl.ReadMessage(h.hub)
}

type RoomRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) GetRooms(w http.ResponseWriter, r *http.Request) {
	rooms := make([]RoomRes, 0)

	for _, r := range h.hub.Rooms {
		rooms = append(rooms, RoomRes{
			// ID:   r.ID,
			Name: r.Name,
		})
	}

}

type ClientRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (h *Handler) GetClients(w http.ResponseWriter, r *http.Request) {
	// var clients []ClientRes
	// roomId := c.Param("roomId")

	// if _, ok := h.hub.Rooms[roomId]; !ok {
	// 	clients = make([]ClientRes, 0)

	// }

	// for _, c := range h.hub.Rooms[roomId].Clients {
	// 	clients = append(clients, ClientRes{
	// 		ID:       c.ID,
	// 		Username: c.Username,
	// 	})
	// }

}
