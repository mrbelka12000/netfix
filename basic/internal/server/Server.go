package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mrbelka12000/netfix/basic/config"
)

func NewServer(r *mux.Router) *http.Server {
	cfg := config.GetConf()

	return &http.Server{
		Addr:         ":" + cfg.App.Port,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  5 * time.Second,
	}
}
