package handler

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"

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

func (h *Handler) GetWork(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		log.Println("missing work id")
		http.Error(w, "missing work id", 400)
		return
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 400)
		return
	}

	work, err := h.srv.GetByID(intID)
	if err != nil {
		log.Println("get work by id error: " + err.Error())
		http.Error(w, err.Error(), 400)
		return
	}

	w.Write([]byte(tools.MakeJsonString(work)))
}
