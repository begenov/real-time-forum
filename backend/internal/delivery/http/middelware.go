package http

import (
	"fmt"
	"net/http"
)

func (h *Handler) setCorsHeaders(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4000")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Cookie, Content-Length, Accept-Encoding, X-CSRF-Token, charset, Credentials, Accept, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		handler.ServeHTTP(w, r)
	})
}
func (h *Handler) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.log.Info(fmt.Sprintf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.String()))

		defer func() {
			if err := recover(); err != nil {
				h.handleError(w, http.StatusInternalServerError, "Internal Server Error")
			}
		}()

		next.ServeHTTP(w, r)
	})
}
