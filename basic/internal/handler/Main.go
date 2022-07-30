package handler

import (
	"log"
	"net/http"

	"github.com/mrbelka12000/netfix/basic/tools"
)

// Main example
// @Summary get all works
// @ID  works
// @Tags service
// @Accept  json
// @Produce  json
// @Success 200 {object} works "okey"
// @Failure 400,404,405,500
// @Router / [get]
func (h *Handler) Main(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	works, err := h.srv.GetAll()
	if err != nil {
		log.Println("get all works error: " + err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte(tools.MakeJsonString(works)))
}

type works []work
