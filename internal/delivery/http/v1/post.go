package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/begenov/real-time-forum/internal/domain"
	"github.com/gorilla/mux"
)

func (h *Handler) InitPostRouter(router *mux.Router) {
	router.HandleFunc("/api/v1/post", h.getAllPosts).Methods("GET")
	router.HandleFunc("/api/v1/post/{id}", h.getPostByID).Methods("GET")
	router.HandleFunc("/api/v1/post/create", h.createPost).Methods("POST")
	router.HandleFunc("/api/v1/post/update/{id}", h.updatePost).Methods("PUT")
	router.HandleFunc("/api/v1/post/delete/{id}", h.deletePost).Methods("DELETE")

}

func (h *Handler) getAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.service.Post.GetAllPosts(context.Background())
	if err != nil {
		msg := fmt.Sprintf("%v", err)
		h.errorResponse(w, msg, http.StatusBadRequest)
		return
	}

	buf, err := json.Marshal(&posts)
	if err != nil {
		msg := fmt.Sprintf("%v", err)
		h.errorResponse(w, msg, http.StatusBadRequest)
		return
	}
	fmt.Println(posts)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(buf)
}

type postInput struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Category    []string `json:"category"`
}

func (h *Handler) createPost(w http.ResponseWriter, r *http.Request) {
	var inp postInput
	if err := json.NewDecoder(r.Body).Decode(&inp); err != nil {
		msg := fmt.Sprintf("error decode: %v", err)
		h.errorResponse(w, msg, http.StatusBadRequest)
		return
	}
	if err := h.service.Post.Create(context.Background(), domain.Post{
		Title:       inp.Title,
		Description: inp.Description,
		Category:    inp.Category,
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
	}); err != nil {
		msg := fmt.Sprintf("%v", err)
		h.errorResponse(w, msg, http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{
		"msg":"Create Post"
	}`)
}

func (h *Handler) getPostByID(w http.ResponseWriter, r *http.Request) {
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

	post, err := h.service.Post.GetPostById(context.Background(), id)

	if err != nil {
		msg := fmt.Sprintf("%v", err)
		h.errorResponse(w, msg, http.StatusBadRequest)
		return
	}
	body, err := json.Marshal(post)
	if err != nil {
		msg := fmt.Sprintf("%v", err)
		h.errorResponse(w, msg, http.StatusBadRequest)
		return
	}

	fmt.Println(post)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)

}

func (h *Handler) updatePost(w http.ResponseWriter, r *http.Request) {
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
	var inp postInput
	if err := json.NewDecoder(r.Body).Decode(&inp); err != nil {
		msg := fmt.Sprintf("error decode: %v", err)
		h.errorResponse(w, msg, http.StatusBadRequest)
		return
	}
	if err := h.service.Post.Update(context.Background(), domain.Post{
		Title:       inp.Title,
		Description: inp.Description,
		Category:    inp.Category,
		UpdateAt:    time.Now(),
		ID:          id,
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

func (h *Handler) deletePost(w http.ResponseWriter, r *http.Request) {
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
	if err := h.service.Post.Delete(r.Context(), id); err != nil {
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
