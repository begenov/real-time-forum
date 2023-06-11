package v1

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) initCategoryRouter(router *mux.Router) {
	router.HandleFunc("/api/v1/categories", h.userIdentity(h.getCategories)).Methods("GET")
}

func (h *Handler) getCategories(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(userID)
	if userID, ok := userId.(int); !ok || userID <= 0 {
		h.handleError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	res, err := h.service.Post.GetAllCategories(context.Background())
	if err != nil {
		log.Println(err)
		h.handleError(w, http.StatusBadRequest, "Bad Request")
		return
	}

	h.writeJSONResponse(w, http.StatusOK, res)
}
