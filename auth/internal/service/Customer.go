package service

import (
	"github.com/mrbelka12000/netfix/auth/internal/repository"
	"github.com/mrbelka12000/netfix/auth/models"
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
