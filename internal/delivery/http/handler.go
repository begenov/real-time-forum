package http

import (
	"encoding/json"
	"net/http"

	v1 "github.com/begenov/real-time-forum/internal/delivery/http/v1"
	"github.com/begenov/real-time-forum/internal/service"
	"github.com/begenov/real-time-forum/pkg/logger"
	"github.com/gorilla/mux"
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

func (h *Handler) Init() http.Handler {
	router := mux.NewRouter()
	h.initRouter(router)
	return h.setCorsHeaders(router)
}

func (h *Handler) initRouter(router *mux.Router) {
	v1 := v1.NewHandler(h.service, h.log)
	router.Use(h.logRequest)
	v1.InitUserRouter(router)
	v1.InitPostRouter(router)
	v1.InitCommentRouter(router)
}

func (h *Handler) handleError(w http.ResponseWriter, statusCode int, err interface{}) {
	h.log.Error(err)

	errorResponse := map[string]interface{}{
		"error": err,
	}

	h.writeJSONResponse(w, statusCode, errorResponse)
}

func (h *Handler) writeJSONResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		h.handleError(w, http.StatusInternalServerError, "Failed to encode JSON response")
		return
	}
}
