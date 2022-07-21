package repository

import (
	"errors"

	"github.com/mrbelka12000/netfix/auth/models"
)

type repoCustomer struct{}

func newCustomer() *repoCustomer {
	return &repoCustomer{}
}

func (rc *repoCustomer) RegisterCustomer(customer *models.Customer) error {
	conn := GetConnection()
	if conn == nil {
		return errors.New("db problemm")
	}
	return nil
}
