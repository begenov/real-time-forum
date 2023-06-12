package v1

import (
	"github.com/begenov/real-time-forum/internal/delivery/http/ws"
	"github.com/begenov/real-time-forum/internal/service"
	"github.com/begenov/real-time-forum/pkg/logger"
	"github.com/gorilla/mux"
)

type Handler struct {
	service *service.Service
	log     *logger.Log
	ws      *ws.Handler
}

func NewHandler(service *service.Service, log *logger.Log, ws *ws.Handler) *Handler {
	return &Handler{
		service: service,
		log:     log,
		ws:      ws,
	}
}

func (h *Handler) InitRouters(router *mux.Router) {
	h.initUserRouter(router)
	h.initCategoryRouter(router)
	h.initPostRouter(router)
	h.initCommentRouter(router)
	h.initWSRouter(router)
}
