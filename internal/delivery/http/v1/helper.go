package v1

import (
	"encoding/json"
	"io"
	"net/http"
)

func parseInput(body io.Reader, inp interface{}) error {
	err := json.NewDecoder(body).Decode(&inp)
	return err
}

func (h *Handler) handleError(w http.ResponseWriter, statusCode int, err interface{}) {
	h.log.Error(err)

	errorResponse := map[string]interface{}{
		"error": err,
	}

	h.writeJSONResponse(w, statusCode, errorResponse)
}

func (h *Handler) writeJSONResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		h.handleError(w, http.StatusInternalServerError, "Failed to encode JSON response")
		return
	}
}
