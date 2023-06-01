package http

import (
	v1 "github.com/begenov/real-time-forum/internal/delivery/http/v1"
	"github.com/begenov/real-time-forum/internal/service"
	"github.com/gorilla/mux"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Init() *mux.Router {
	router := mux.NewRouter()
	h.initRouter(router)
	return router
}

func (h *Handler) initRouter(router *mux.Router) {
	v1 := v1.NewHandler(h.service)

	v1.InitUserRouter(router)
	v1.InitPostRouter(router)
	v1.InitCommentRouter(router)
}
