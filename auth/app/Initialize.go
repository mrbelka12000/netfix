package app

import (
	"github.com/mrbelka12000/netfix/auth/delivery"
	"github.com/mrbelka12000/netfix/auth/repository"
	"github.com/mrbelka12000/netfix/auth/service"
)

func Initialize() {
	repo := repository.NewRepo()
	srv := service.NewService(repo)
	del := delivery.NewDelivery(srv)

	go del.ConsumerForCustomer()
	del.ConsumerForCompany()
	//del.Produce()
}
