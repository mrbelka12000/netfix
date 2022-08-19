package app

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/mrbelka12000/netfix/billing/internal/delivery"
	"github.com/mrbelka12000/netfix/billing/internal/repository"
	"github.com/mrbelka12000/netfix/billing/internal/service"
	"github.com/mrbelka12000/netfix/billing/internal/stepfuncs"
)

const consumersAmount = 4

func Initialize() {
	repo := repository.NewRepo()
	srv := service.NewService(repo)
	d := delivery.NewDelivery(srv)
	bil := make(chan []byte)
	exit := make(chan struct{}, consumersAmount)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	wg := &sync.WaitGroup{}

	wg.Add(consumersAmount)
	go d.ConsumerForGetWallet(exit, wg)
	go d.ConsumerForWallets(exit, wg)
	go d.ConsumerForBilling(bil, exit, wg)
	go stepfuncs.Billing(bil, exit, wg)

	log.Println("billing server started")
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
		close(bil)
		close(done)
	}()

	exitProgramm(exit)

	wg.Wait()
	log.Println("billing service exited properly")
}

func exitProgramm(exit chan struct{}) {
	exit <- struct{}{}
	exit <- struct{}{}
	exit <- struct{}{}
	exit <- struct{}{}

	<-exit
	<-exit
	<-exit
	<-exit
}
