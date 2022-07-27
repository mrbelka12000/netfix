package app

import (
	"github.com/mrbelka12000/netfix/billing/internal/delivery"
	"github.com/mrbelka12000/netfix/billing/internal/repository"
	"github.com/mrbelka12000/netfix/billing/internal/service"
	"github.com/mrbelka12000/netfix/billing/internal/stepfuncs"
)

func Initialize() {
	repo := repository.NewRepo()
	srv := service.NewService(repo)
	d := delivery.NewDelivery(srv)
	bil := make(chan []byte)
	exit := make(chan struct{})

	go d.ConsumerForGetWallet()
	go d.ConsumerForWallets()
	go d.ConsumerForBilling(bil)
	stepfuncs.Billing(bil, exit)
}
