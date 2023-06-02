package v1

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"
)

type user string

const userID user = "userID"

func (h *Handler) userIdentity(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil {
			switch {
			case errors.Is(err, http.ErrNoCookie):
				next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), userID, 0)))
			case errors.Is(err, cookie.Valid()):
				log.Println("invalid cookie value")
			}
			log.Println("failed to get cookie")
			return
		}

		u, err := h.service.User.GetUserByToken(r.Context(), cookie.Value)
		if err != nil {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), userID, 0)))
			return
		}
		if u.ExpiresAt.Before(time.Now()) {
			if err := h.service.User.DeleteSession(r.Context(), cookie.Value); err != nil {
				return
			}
			http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
			return
		}
		log.Println(cookie.Value)
		next(w, r.WithContext(context.WithValue(r.Context(), userID, u.Id)))
	}
}
