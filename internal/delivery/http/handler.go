package http

import (
	"net/http"

	v1 "github.com/begenov/real-time-forum/internal/delivery/http/v1"
	"github.com/begenov/real-time-forum/internal/delivery/http/ws"
	"github.com/begenov/real-time-forum/internal/domain"
	"github.com/begenov/real-time-forum/internal/service"
	"github.com/begenov/real-time-forum/pkg/logger"
	"github.com/gorilla/mux"
)

type Handler struct {
	service *service.Service
	log     *logger.Log
	wsEvent chan *domain.WSEvent
}

func NewHandler(service *service.Service, log *logger.Log) *Handler {
	return &Handler{
		service: service,
		log:     log,
		wsEvent: make(chan *domain.WSEvent),
	}

}

func (h *Handler) Init() http.Handler {
	router := mux.NewRouter()
	h.initRouter(router)
	return h.setCorsHeaders(router)
}

func (h *Handler) initRouter(router *mux.Router) {
	ws := ws.NewHandler(h.service, h.wsEvent)
	v1 := v1.NewHandler(h.service, h.log, ws)
	router.Use(h.logRequest)
	v1.InitRouters(router)
}
