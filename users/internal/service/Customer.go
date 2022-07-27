package service

import (
	"database/sql"
	"github.com/mrbelka12000/netfix/users/internal/repository"
	"github.com/mrbelka12000/netfix/users/models"
)

type srvCustomer struct {
	repo *repository.Repository
}

func newCustomer(repo *repository.Repository) *srvCustomer {
	return &srvCustomer{repo}
}

func (sc *srvCustomer) RegisterCustomer(customer *models.Customer, tx *sql.Tx) error {
	return sc.repo.RegisterCustomer(customer, tx)
}

func (sc *srvCustomer) GetByID(id int) (*models.General, error) {
	return sc.repo.Customer.GetByID(id)
}
