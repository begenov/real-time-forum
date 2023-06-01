package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/begenov/real-time-forum/internal/domain"
	"github.com/gorilla/mux"
)

func (h *Handler) InitCommentRouter(router *mux.Router) {
	router.HandleFunc("/api/v1/post/comment", h.getAllComment)
	router.HandleFunc("/api/v1/post/comment/{id}", h.getCommentByID)
	router.HandleFunc("/api/v1/post/comment/create", h.createComment)
	router.HandleFunc("/api/v1/post/comment/update/{id}", h.updateComment)
	router.HandleFunc("/api/v1/post/comment/delete/{id}", h.deleteComment)
}

type commentInput struct {
	Text string `json:"text"`
}

func (h *Handler) createComment(w http.ResponseWriter, r *http.Request) {
	var inp commentInput
	if err := json.NewDecoder(r.Body).Decode(&inp); err != nil {
		msg := fmt.Sprintf("error decode: %v", err)
		h.errorResponse(w, msg, http.StatusBadRequest)
		return
	}

	h.service.Comment.Create(r.Context(), domain.Comment{
		Text:     inp.Text,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	})
}

func (h *Handler) getAllComment(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) getCommentByID(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) deleteComment(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) updateComment(w http.ResponseWriter, r *http.Request) {
}
