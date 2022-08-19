package app

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/mrbelka12000/netfix/users/internal/delivery"
	"github.com/mrbelka12000/netfix/users/internal/repository"
	"github.com/mrbelka12000/netfix/users/internal/service"
)

const consumersAmount = 5

func Initialize() {
	repo := repository.NewRepo()
	srv := service.NewService(repo)
	del := delivery.NewDelivery(srv)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	exit := make(chan struct{}, consumersAmount)
	wg := &sync.WaitGroup{}

	wg.Add(consumersAmount)
	go del.ConsumerForCustomer(exit, wg)
	go del.ConsumerForGetCompany(exit, wg)
	go del.ConsumerForGetCustomer(exit, wg)
	go del.ConsumerForLogin(exit, wg)
	go del.ConsumerForCompany(exit, wg)

	log.Println("users service started")
	<-done
	defer func() {
		conn := repository.GetConnection()
		err := conn.Close()
		if err != nil {
			log.Printf("close connection to db error: %v \n", err)
		} else {
			log.Println("connection to db successfully canceled")
		}
		close(exit)
		close(done)
	}()
	exitProgramm(exit)

	wg.Wait()

	log.Println("users service exited properly")
}

func exitProgramm(exit chan struct{}) {
	exit <- struct{}{}
	exit <- struct{}{}
	exit <- struct{}{}
	exit <- struct{}{}
	exit <- struct{}{}

	<-exit
	<-exit
	<-exit
	<-exit
	<-exit
}
