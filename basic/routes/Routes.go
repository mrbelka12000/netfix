package routes

import (
	"github.com/gorilla/mux"
	"github.com/mrbelka12000/netfix/basic/handler"
)

func SetUpMux(h *handler.Handler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", h.Main)
	r.HandleFunc("/register/company", h.RegisterCompany)
	r.HandleFunc("/register/customer", nil)

	r.HandleFunc("/service", nil)
	r.HandleFunc("/service/apply", nil)
	return r
}
