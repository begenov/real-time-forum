package v1

import (
	"github.com/begenov/real-time-forum/internal/service"
	"github.com/begenov/real-time-forum/pkg/logger"
)

type Handler struct {
	service *service.Service
	log     *logger.Log
}

func NewHandler(service *service.Service, log *logger.Log) *Handler {
	return &Handler{
		service: service,
		log:     log,
	}
}
