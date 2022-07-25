package handler

import (
	"github.com/gorilla/mux"
	"github.com/mrbelka12000/netfix/basic/models"
	"log"
	"net/http"
	"strconv"

	"github.com/mrbelka12000/netfix/basic/tools"
)

// GetWorkFields example
// @Summary get all work fields
// @ID  work fields
// @Tags service
// @Accept  json
// @Produce  json
// @Success 200 {object} workFields "okey"
// @Failure 400,404,405,500
// @Router /workfields [get]
func (h *Handler) GetWorkFields(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	works, err := h.srv.GetWorkFields()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte(tools.MakeJsonString(works)))
}

// GetWork example
// @Summary get work by id
// @ID  get work
// @Tags service
// @Accept  json
// @Produce  json
// @Success 200 {object} work "okey"
// @Failure 400,404,405,500
// @Router  /service/{id} [get]
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

type workFields struct {
	*models.SwaggerWorkFields
}

type work struct {
	*models.Work
}
