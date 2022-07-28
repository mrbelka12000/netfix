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
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}
}
