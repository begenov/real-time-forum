package v1

import "github.com/gorilla/mux"

func (h *Handler) initWSRouter(router *mux.Router) {
	router.HandleFunc("/api/v1/ws", h.userIdentity(h.ws.ServeWS))
}
