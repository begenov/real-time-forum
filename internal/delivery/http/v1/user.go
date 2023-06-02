package v1

import (
	"net/http"

	"github.com/begenov/real-time-forum/internal/domain"
	"github.com/gorilla/mux"
)

func (h *Handler) InitUserRouter(router *mux.Router) {
	router.HandleFunc("/api/v1/sign-up", h.signUp).Methods("POST")
	router.HandleFunc("/api/v1/sign-in", h.signIn).Methods("POST")
	router.HandleFunc("/api/v1/log-out", h.logOut).Methods("POST")
	router.HandleFunc("/api/v1/check-user", h.checkUser).Methods("GET")
}

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var inp domain.UserInput
	err := parseInput(r.Body, inp)
	if err != nil {
		h.handleError(w, http.StatusBadRequest, "Failed to decode JSON: "+err.Error())
		return
	}

	if err := h.service.User.SignUp(r.Context(), domain.User{
		Nickname:  inp.Nickname,
		Age:       inp.Age,
		Gender:    inp.Gender,
		FirstName: inp.FirstName,
		LastName:  inp.LastName,
		Email:     inp.Email,
		Password:  inp.Password,
	}); err != nil {
		h.handleError(w, http.StatusBadRequest, "Bad Request "+err.Error())
		return
	}

	h.writeJSONResponse(w, http.StatusOK, map[string]string{
		"msg": "User created successfully",
	})
}

type signIn struct {
	Auth     string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	var inp signIn

	if err := parseInput(r.Body, inp); err != nil {
		h.handleError(w, http.StatusBadRequest, "Failed to decode JSON: "+err.Error())
		return
	}

	session, err := h.service.User.SignIn(r.Context(), inp.Auth, inp.Password)
	if err != nil {
		h.handleError(w, http.StatusBadRequest, "Bad Request "+err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Path:     "/",
		Name:     "session",
		Expires:  session.ExpiresAt,
		Value:    session.Token,
		HttpOnly: true,
	})

	h.writeJSONResponse(w, http.StatusOK, map[string]string{
		"token": session.Token,
	})
}

func (h *Handler) logOut(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		h.handleError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if err := h.service.User.DeleteSession(r.Context(), cookie.Value); err != nil {
		h.handleError(w, http.StatusBadRequest, "Failed to log out "+err.Error())
		return
	}

	cookie.MaxAge = -1
	cookie.Value = ""
	cookie.Name = "session"
	cookie.Path = "/"
	http.SetCookie(w, cookie)

	h.writeJSONResponse(w, http.StatusOK, "Log out success")
}

func (h *Handler) checkUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		h.handleError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	user, err := h.service.User.GetUserByToken(r.Context(), cookie.Value)
	if err != nil {
		h.handleError(w, http.StatusBadRequest, "Bad Request "+err.Error())
		return
	}

	h.writeJSONResponse(w, http.StatusOK, user.Nickname)
}
