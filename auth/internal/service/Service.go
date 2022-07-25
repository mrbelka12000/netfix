package service

import (
	"github.com/mrbelka12000/netfix/auth/internal/repository"
	"github.com/mrbelka12000/netfix/auth/models"
)

type Company interface {
	RegisterCompany(company *models.Company) error
}

type Customer interface {
	RegisterCustomer(customer *models.Customer) error
}

type General interface {
	Register(general *models.General) (int, error)
}

type Service struct {
	Company
	Customer
	General
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Company:  newCompany(repo),
		Customer: newCustomer(repo),
		General:  newGeneral(repo),
	}
}
