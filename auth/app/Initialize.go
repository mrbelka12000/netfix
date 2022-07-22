package app

import (
	"github.com/mrbelka12000/netfix/auth/config"
	"github.com/mrbelka12000/netfix/auth/delivery"
	"github.com/mrbelka12000/netfix/auth/repository"
	"github.com/mrbelka12000/netfix/auth/service"
)

func Initialize(cfg *config.Config) {
	repo := repository.NewRepo()
	srv := service.NewService(repo)
	del := delivery.NewDelivery(srv)

	go del.ConsumerForCompany(cfg)
	//go del.ConsumerForCustomer(cfg)
	del.Produce(cfg)
}
