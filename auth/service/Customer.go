package service

import (
	"github.com/mrbelka12000/netfix/auth/models"
	"github.com/mrbelka12000/netfix/auth/repository"
)

type srvCustomer struct {
	repo *repository.Repository
}

func newCustomer(repo *repository.Repository) *srvCustomer {
	return &srvCustomer{repo}
}

func (sc *srvCustomer) RegisterCustomer(customer *models.Customer) error {
	return sc.repo.RegisterCustomer(customer)
}
