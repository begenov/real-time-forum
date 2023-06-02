package v1

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/begenov/real-time-forum/internal/domain"
	"github.com/gorilla/mux"
)

func (h *Handler) InitCommentRouter(router *mux.Router) {
	router.HandleFunc("/api/v1/post/comment", h.getAllComment)
	router.HandleFunc("/api/v1/post/comment/{id}", (h.getCommentByID))
	router.HandleFunc("/api/v1/post/comment/create/{id}", h.userIdentity(h.createComment))
	router.HandleFunc("/api/v1/post/comment/update/{id}", h.userIdentity(h.updateComment))
	router.HandleFunc("/api/v1/post/comment/delete/{id}", h.userIdentity(h.deleteComment))
}

type commentInput struct {
	Text string `json:"text"`
}

func (h *Handler) createComment(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(userID)
	if userId.(int) <= 0 {
		log.Println("comment")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, `{
			"msg":"Status Unauthorized"
		}`)
		return
	}

	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		msg := "ID not found in URL path"
		h.errorResponse(w, msg, http.StatusBadRequest)
		return
	}

	postID, err := strconv.Atoi(idStr)
	if err != nil {
		msg := fmt.Sprintf("%v", err)
		h.errorResponse(w, msg, http.StatusBadRequest)
		return
	}

	var inp commentInput
	if err := json.NewDecoder(r.Body).Decode(&inp); err != nil {
		msg := fmt.Sprintf("error decode: %v", err)

		h.errorResponse(w, msg, http.StatusBadRequest)
		return
	}

	if err := h.service.Comment.Create(r.Context(), domain.Comment{
		Text:     inp.Text,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
		UserID:   userId.(int),
		PostID:   postID,
	}); err != nil {

		msg := fmt.Sprintf("error decode: %v", err)
		h.errorResponse(w, msg, http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{
		"msg":"Create Comment"
	}`)
}

func (h *Handler) getAllComment(w http.ResponseWriter, r *http.Request) {
	comments, err := h.service.Comment.GetAllComment(r.Context())
	if err != nil {
		msg := fmt.Sprintf("%v", err)
		h.errorResponse(w, msg, http.StatusBadRequest)
		return
	}
	buf, err := json.Marshal(&comments)
	if err != nil {
		msg := fmt.Sprintf("%v", err)
		h.errorResponse(w, msg, http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(buf)
}

func (h *Handler) getCommentByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		msg := "ID not found in URL path"
		h.errorResponse(w, msg, http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		msg := fmt.Sprintf("%v", err)
		h.errorResponse(w, msg, http.StatusBadRequest)
		return
	}
	comment, err := h.service.Comment.GetCommentById(r.Context(), id)
	if err != nil {
		msg := fmt.Sprintf("%v", err)
		h.errorResponse(w, msg, http.StatusBadRequest)
		return
	}
	body, err := json.Marshal(comment)
	if err != nil {
		msg := fmt.Sprintf("%v", err)
		h.errorResponse(w, msg, http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)

}

func (h *Handler) deleteComment(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(userID)
	if userId.(int) <= 0 {
		log.Println("comment")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, `{
			"msg":"Status Unauthorized"
		}`)
		return
	}
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		msg := "ID not found in URL path"
		h.errorResponse(w, msg, http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		msg := fmt.Sprintf("%v", err)
		h.errorResponse(w, msg, http.StatusBadRequest)
		return
	}
	if err := h.service.Comment.Delete(r.Context(), id, userId.(int)); err != nil {
		msg := fmt.Sprintf("%v", err)
		h.errorResponse(w, msg, http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{
		"msg":"Delete success"
	}`)
}

func (h *Handler) updateComment(w http.ResponseWriter, r *http.Request) {
	var inp commentInput
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		msg := "ID not found in URL path"
		h.errorResponse(w, msg, http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		msg := fmt.Sprintf("%v", err)
		h.errorResponse(w, msg, http.StatusBadRequest)
		return
	}
	if err := h.service.Comment.Update(r.Context(), domain.Comment{
		Text:     inp.Text,
		UpdateAt: time.Now(),
		Id:       id,
	}); err != nil {
		msg := fmt.Sprintf("%v", err)
		h.errorResponse(w, msg, http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{
		"msg":"Update success"
	}`)
}
