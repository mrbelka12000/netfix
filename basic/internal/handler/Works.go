package handler

import (
	"net/http"

	"github.com/mrbelka12000/netfix/basic/tools"
)

func (h *Handler) GetWorkFields(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	works, err := h.srv.GetWorkFields()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte(tools.MakeJsonString(works)))
}
