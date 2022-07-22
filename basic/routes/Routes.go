package routes

import (
	"github.com/gorilla/mux"
	"github.com/mrbelka12000/netfix/basic/handler"
	"net/http"
)

func SetUpMux(h *handler.Handler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", h.Main).Methods(http.MethodGet)
	r.HandleFunc("/register/company", h.RegisterCompany).Methods(http.MethodPost)
	r.HandleFunc("/register/customer", h.RegisterCustomer).Methods(http.MethodPost)

	r.HandleFunc("/service", h.CreateService).Methods(http.MethodPost)
	r.HandleFunc("/service/apply", nil)
	return r
}
