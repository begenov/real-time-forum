package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/begenov/real-time-forum/internal/domain"
	"github.com/gorilla/mux"
)

func (h *Handler) InitUserRouter(router *mux.Router) {
	router.HandleFunc("/api/v1/sign-up", h.signUp).Methods("POST")
	router.HandleFunc("/api/v1/sign-in", h.signIn).Methods("POST")
	router.HandleFunc("/api/v1/log-out", h.logOut).Methods("POST")
}

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {

	var inp domain.UserInput

	if err := json.NewDecoder(r.Body).Decode(&inp); err != nil {
		msg := fmt.Sprintf("error decode: %v", err)
		h.errorResponse(w, msg, http.StatusBadRequest)
		return
	}

	if err := h.service.User.SignUp(context.Background(), domain.User{
		Nickname:  inp.Nickname,
		Age:       inp.Age,
		Gender:    inp.Gender,
		FirstName: inp.FirstName,
		LastName:  inp.LastName,
		Email:     inp.Email,
		Password:  inp.Password,
	}); err != nil {
		msg := fmt.Sprintf(" %v", err)
		h.errorResponse(w, msg, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{
		"msg":"Create User"
	}`)
}

type singIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	var inp singIn

	if err := json.NewDecoder(r.Body).Decode(&inp); err != nil {
		msg := fmt.Sprintf("error decode: %v", err)
		h.errorResponse(w, msg, http.StatusBadRequest)
		return
	}

	session, err := h.service.User.SignIn(context.Background(), inp.Email, inp.Password)

	if err != nil {
		msg := fmt.Sprintf("%v", err)
		h.errorResponse(w, msg, http.StatusBadRequest)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Path:     "/",
		Name:     "session",
		Expires:  session.ExpiresAt,
		Value:    session.Token,
		HttpOnly: true,
	})
	var token struct {
		Token string `json:"token"`
	}
	token.Token = session.Token
	data, err := json.Marshal(&token)
	if err != nil {
		msg := fmt.Sprintf("%v", err)
		h.errorResponse(w, msg, http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *Handler) logOut(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err == http.ErrNoCookie {
		log.Fatal(err)
		return
	}
	if err != nil {
		log.Fatal(err)
		return
	}

	if err := h.service.User.DeleteSession(context.Background(), cookie.Value); err != nil {
		log.Fatal(err, "error")
		return
	}

	cookie.MaxAge = -1
	cookie.Value = ""
	cookie.Name = "session"
	cookie.Path = "/"
	http.SetCookie(w, cookie)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Log out success"))
}
