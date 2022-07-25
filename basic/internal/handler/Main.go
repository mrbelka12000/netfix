package handler

import "net/http"

func (h *Handler) Main(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte("OK"))
}
