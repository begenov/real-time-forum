package v1

import (
	"context"
	"errors"
	"net/http"
	"time"
)

type user string

const userID user = "userID"

func (h *Handler) userIdentity(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), userID, 0)))
			} else {
				h.handleError(w, http.StatusInternalServerError, "Failed to get cookie")
			}
			return
		}

		user, err := h.service.User.GetUserByToken(r.Context(), cookie.Value)
		if err != nil {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), userID, 0)))
			return
		}

		if user.ExpiresAt.Before(time.Now()) {
			if err := h.service.User.DeleteSession(r.Context(), cookie.Value); err != nil {
				h.handleError(w, http.StatusInternalServerError, "Failed to delete session")
				return
			}

			return
		}

		next(w, r.WithContext(context.WithValue(r.Context(), userID, user.Id)))
	}
}
