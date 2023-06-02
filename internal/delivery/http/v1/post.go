package v1

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/begenov/real-time-forum/internal/domain"
	"github.com/gorilla/mux"
)

func (h *Handler) InitPostRouter(router *mux.Router) {
	router.HandleFunc("/api/v1/post", h.userIdentity(h.getAllPosts)).Methods("GET")
	router.HandleFunc("/api/v1/post/{id}", h.userIdentity(h.getPostByID)).Methods("GET")
	router.HandleFunc("/api/v1/post/create", h.userIdentity(h.createPost)).Methods("POST")
	router.HandleFunc("/api/v1/post/update/{id}", h.userIdentity(h.updatePost)).Methods("PUT")
	router.HandleFunc("/api/v1/post/delete/{id}", h.userIdentity(h.deletePost)).Methods("DELETE")
}

type postInput struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Category    []string `json:"category"`
}

func (h *Handler) createPost(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(userID)
	if userID, ok := userId.(int); !ok || userID <= 0 {
		h.handleError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var inp postInput
	if err := parseInput(r.Body, &inp); err != nil {
		h.handleError(w, http.StatusBadRequest, "Failed to decode JSON: "+err.Error())
		return
	}

	post := domain.Post{
		Title:       inp.Title,
		Description: inp.Description,
		Category:    inp.Category,
		CreateAt:    time.Now(),
		UserID:      userId.(int),
		UpdateAt:    time.Now(),
	}

	if err := h.service.Post.Create(r.Context(), post); err != nil {
		h.handleError(w, http.StatusBadRequest, "Failed to create post: "+err.Error())
		return
	}

	h.writeJSONResponse(w, http.StatusOK, map[string]string{
		"msg": "Post created successfully",
	})
}

func (h *Handler) getAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.service.Post.GetAllPosts(r.Context())
	if err != nil {
		h.handleError(w, http.StatusBadRequest, "Failed to get posts")
		return
	}

	h.writeJSONResponse(w, http.StatusOK, posts)
}

func (h *Handler) getPostByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		h.handleError(w, http.StatusBadRequest, "ID not found in URL path")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.handleError(w, http.StatusBadRequest, fmt.Sprintf("Invalid ID: %v", err))
		return
	}

	post, err := h.service.Post.GetPostById(r.Context(), id)
	if err != nil {
		h.handleError(w, http.StatusBadRequest, fmt.Sprintf("Failed to get post: %v", err))
		return
	}

	h.writeJSONResponse(w, http.StatusOK, post)
}

func (h *Handler) updatePost(w http.ResponseWriter, r *http.Request) {
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
		h.handleError(w, http.StatusBadRequest, fmt.Sprintf("Invalid ID: %v", err))
		return
	}

	var inp postInput
	if err := parseInput(r.Body, &inp); err != nil {
		h.handleError(w, http.StatusBadRequest, fmt.Sprintf("Failed to decode JSON: %v", err))
		return
	}

	if err := h.service.Post.Update(r.Context(), domain.Post{
		Title:       inp.Title,
		Description: inp.Description,
		Category:    inp.Category,
		UpdateAt:    time.Now(),
		ID:          id,
		UserID:      userId.(int),
	}); err != nil {
		h.handleError(w, http.StatusBadRequest, fmt.Sprintf("Failed to update post: %v", err))
		return
	}

	h.writeJSONResponse(w, http.StatusOK, map[string]string{
		"msg": "Update success",
	})
}

func (h *Handler) deletePost(w http.ResponseWriter, r *http.Request) {
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
		h.handleError(w, http.StatusBadRequest, fmt.Sprintf("Invalid ID: %v", err))
		return
	}

	if err := h.service.Post.Delete(r.Context(), id, userId.(int)); err != nil {
		h.handleError(w, http.StatusBadRequest, fmt.Sprintf("Failed to delete post: %v", err))
		return
	}

	h.writeJSONResponse(w, http.StatusOK, map[string]string{
		"msg": "Delete success",
	})
}
