package v1

import "github.com/gorilla/mux"

func (h *Handler) initWSRouter(router *mux.Router) {
	router.HandleFunc("/ws", h.userIdentity(h.ws.ServeWS))
}
