package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/begenov/real-time-forum/internal/domain"
	"github.com/gorilla/mux"
)

func (h *Handler) InitCommentRouter(router *mux.Router) {
	router.HandleFunc("/api/v1/post/comment", h.userIdentity(h.getAllComment))
	router.HandleFunc("/api/v1/post/comment/{id}", h.userIdentity(h.getCommentByID))
	router.HandleFunc("/api/v1/post/comment/create/{id}", h.userIdentity(h.createComment))
	router.HandleFunc("/api/v1/post/comment/update/{id}", h.userIdentity(h.updateComment))
	router.HandleFunc("/api/v1/post/comment/delete/{id}", h.userIdentity(h.deleteComment))
}

type commentInput struct {
	Text string `json:"text"`
}

func (h *Handler) createComment(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(userID)
	if userID, ok := userId.(int); !ok || userID <= 0 {
		h.handleError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		h.handleError(w, http.StatusBadRequest, "ID not found in URL path")
		return
	}

	postID, err := strconv.Atoi(idStr)
	if err != nil {
		h.handleError(w, http.StatusBadRequest, "Invalid post ID")
		return
	}

	var inp commentInput
	if err := parseInput(r.Body, &inp); err != nil {
		h.handleError(w, http.StatusBadRequest, "Failed to decode request body")
		return
	}

	comment := domain.Comment{
		Text:     inp.Text,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
		UserID:   userId.(int),
		PostID:   postID,
	}

	if err := h.service.Comment.Create(r.Context(), comment); err != nil {
		h.handleError(w, http.StatusInternalServerError, "Failed to create comment")
		return
	}

	h.writeJSONResponse(w, http.StatusOK, map[string]interface{}{
		"msg": "Create Comment",
	})
}

func (h *Handler) getAllComment(w http.ResponseWriter, r *http.Request) {
	comments, err := h.service.Comment.GetAllComment(r.Context())
	if err != nil {
		h.handleError(w, http.StatusBadRequest, "Failed to get comments")
		return
	}

	h.writeJSONResponse(w, http.StatusOK, map[string]interface{}{
		"comments": comments,
	})
}

func (h *Handler) getCommentByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		h.handleError(w, http.StatusBadRequest, "ID not found in URL path")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.handleError(w, http.StatusBadRequest, "Invalid comment ID")
		return
	}

	comment, err := h.service.Comment.GetCommentById(r.Context(), id)
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, "Failed to retrieve comment")
		return
	}

	h.writeJSONResponse(w, http.StatusOK, comment)
}

func (h *Handler) deleteComment(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(userID)
	if userID, ok := userId.(int); !ok || userID <= 0 {
		h.handleError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		h.handleError(w, http.StatusBadRequest, "ID not found in URL path")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.handleError(w, http.StatusBadRequest, "Invalid ID format")
		return
	}
	err = h.service.Comment.Delete(r.Context(), id, userId.(int))
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, "Failed to delete comment")
		return
	}

	h.writeJSONResponse(w, http.StatusOK, map[string]string{
		"msg": "Comment deleted successfully",
	})
}

func (h *Handler) updateComment(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(userID)
	if userID, ok := userId.(int); !ok || userID <= 0 {
		h.handleError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		h.handleError(w, http.StatusBadRequest, "ID not found in URL path")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.handleError(w, http.StatusBadRequest, "Invalid comment ID")
		return
	}

	var inp commentInput
	if err := parseInput(r.Body, &inp); err != nil {
		h.handleError(w, http.StatusBadRequest, "Failed to decode request body")
		return
	}

	if err := h.service.Comment.Update(r.Context(), domain.Comment{
		Text:     inp.Text,
		UpdateAt: time.Now(),
		Id:       id,
		UserID:   userId.(int),
	}); err != nil {
		h.handleError(w, http.StatusInternalServerError, "Failed to update comment")
		return
	}

	h.writeJSONResponse(w, http.StatusOK, map[string]string{
		"msg": "Update success",
	})
}
