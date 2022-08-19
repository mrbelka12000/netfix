package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mrbelka12000/netfix/basic/internal/handler"
	"github.com/mrbelka12000/netfix/basic/internal/repository"
	"github.com/mrbelka12000/netfix/basic/internal/routes"
	"github.com/mrbelka12000/netfix/basic/internal/server"
	"github.com/mrbelka12000/netfix/basic/internal/service"
	"github.com/mrbelka12000/netfix/basic/redis"
)

func Initialize() {
	repo := repository.NewRepo()
	srv := service.NewService(repo)
	h := handler.NewHandler(srv)
	router := routes.SetUpMux(h)
	s := server.NewServer(router)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {

		err := s.ListenAndServe()
		if err != nil {
			log.Println(err.Error())
			return
		}
	}()
	log.Println("server started")
	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer func() {
		conn := repository.GetConnection()
		err := conn.Close()
		if err != nil {
			log.Printf("close connection to db error: %v \n", err)
		} else {
			log.Println("connection to db successfully canceled")
		}

		err = redis.CloseRedis()
		if err != nil {
			log.Printf("close redis connection error: %v \n", err)
		} else {
			log.Println("redis connection closed")
		}

		close(done)

		cancel()
	}()

	err := s.Shutdown(ctx)
	if err != nil {
		log.Printf("Server Shutdown Failed: %v", err)
		return
	}

	log.Println("basic service exited properly")
}
