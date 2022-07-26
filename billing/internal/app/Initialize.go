package app

import (
	"github.com/mrbelka12000/netfix/billing/internal/delivery"
	"github.com/mrbelka12000/netfix/billing/internal/repository"
	"github.com/mrbelka12000/netfix/billing/internal/service"
)

func Initialize() {
	repo := repository.NewRepo()
	srv := service.NewService(repo)
	d := delivery.NewDelivery(srv)

	d.ConsumerForWallets()
}
