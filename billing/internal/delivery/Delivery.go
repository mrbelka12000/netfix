package delivery

import "github.com/mrbelka12000/netfix/billing/internal/service"

type Delivery struct {
	srv *service.Service
}

func NewDelivery(srv *service.Service) *Delivery {
	return &Delivery{srv}
}
