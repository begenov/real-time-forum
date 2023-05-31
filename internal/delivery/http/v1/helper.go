package v1

import "net/http"

func (h *Handler) errorResponse(w http.ResponseWriter, msg string, code int) {
	w.WriteHeader(code)
	w.Write([]byte(msg))
}
