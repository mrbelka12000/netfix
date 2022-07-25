package app

import (
	"log"

	"github.com/mrbelka12000/netfix/basic/internal/handler"
	"github.com/mrbelka12000/netfix/basic/internal/repository"
	"github.com/mrbelka12000/netfix/basic/internal/routes"
	"github.com/mrbelka12000/netfix/basic/internal/server"
	"github.com/mrbelka12000/netfix/basic/internal/service"
)

func Initialize() {
	repo := repository.NewRepo()
	srv := service.NewService(repo)
	h := handler.NewHandler(srv)
	router := routes.SetUpMux(h)
	s := server.NewServer(router)

	err := s.ListenAndServe()
	if err != nil {
		log.Println(err.Error())
		return
	}
}
