package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrbelka12000/netfix/basic/internal/handler"
	_ "github.com/mrbelka12000/netfix/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetUpMux(h *handler.Handler) *mux.Router {
	r := mux.NewRouter()

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	r.HandleFunc("/", h.Main).Methods(http.MethodGet)
	r.HandleFunc("/register/company", h.RegisterCompany).Methods(http.MethodPost)
	r.HandleFunc("/register/customer", h.RegisterCustomer).Methods(http.MethodPost)
	r.HandleFunc("/login", h.Login).Methods(http.MethodPost)
	r.HandleFunc("/profile", h.Profile).Methods(http.MethodGet)

	r.HandleFunc("/service", h.CreateService).Methods(http.MethodPost)
	r.HandleFunc("/service/apply", h.ApplyForWork).Methods(http.MethodPost)
	r.HandleFunc("/service/{id}", h.GetWork).Methods(http.MethodGet)
	r.HandleFunc("/service/finish", h.FinishWork).Methods(http.MethodPost)
	// get service nado

	r.HandleFunc("/workfields", h.GetWorkFields).Methods(http.MethodGet)
	return r
}
