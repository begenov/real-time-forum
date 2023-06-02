package v1

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) InitCategoryRouter(router *mux.Router) {
	router.HandleFunc("/api/v1/post/categories", h.userIdentity(h.getCategories)).Methods("GET")
}

func (h *Handler) getCategories(w http.ResponseWriter, r *http.Request) {
	res := []string{"Rust", "Go", "Python"}
	h.writeJSONResponse(w, http.StatusOK, map[string][]string{
		"categories": res,
	})
}
