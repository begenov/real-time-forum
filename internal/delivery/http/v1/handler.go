package v1

import (
	"github.com/begenov/real-time-forum/internal/service"
	"github.com/begenov/real-time-forum/pkg/logger"
	"github.com/gorilla/websocket"
)

type Handler struct {
	service           *service.Service
	log               *logger.Log
	activeConnections []*websocket.Conn
}

func NewHandler(service *service.Service, log *logger.Log) *Handler {
	return &Handler{
		service: service,
		log:     log,
	}
}
