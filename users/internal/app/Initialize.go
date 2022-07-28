package app

import (
	"github.com/mrbelka12000/netfix/users/internal/delivery"
	"github.com/mrbelka12000/netfix/users/internal/repository"
	"github.com/mrbelka12000/netfix/users/internal/service"
)

func Initialize() {
	repo := repository.NewRepo()
	srv := service.NewService(repo)
	del := delivery.NewDelivery(srv)

	go del.ConsumerForCustomer()
	go del.ConsumerForGetCompany()
	go del.ConsumerForGetCustomer()
	go del.ConsumerForLogin()
	del.ConsumerForCompany()
}
