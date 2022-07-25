package app

import (
	"github.com/mrbelka12000/netfix/auth/internal/delivery"
	"github.com/mrbelka12000/netfix/auth/internal/repository"
	"github.com/mrbelka12000/netfix/auth/internal/service"
)

func Initialize() {
	repo := repository.NewRepo()
	srv := service.NewService(repo)
	del := delivery.NewDelivery(srv)

	go del.ConsumerForCustomer()
	del.ConsumerForCompany()
	// del.Produce()
}
