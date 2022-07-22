package app

import (
	"github.com/mrbelka12000/netfix/basic/handler"
	"github.com/mrbelka12000/netfix/basic/repository"
	"github.com/mrbelka12000/netfix/basic/routes"
	"github.com/mrbelka12000/netfix/basic/server"
	"github.com/mrbelka12000/netfix/basic/service"
	"log"
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
