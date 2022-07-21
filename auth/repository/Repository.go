package repository

import "github.com/mrbelka12000/netfix/auth/models"

type Company interface {
	RegisterCompany(company *models.Company) error
}

type Customer interface {
	RegisterCustomer(customer *models.Customer) error
}

type General interface {
	Register(general *models.General) (int, error)
}

type Repository struct {
	Company
	Customer
	General
}

func NewRepo() *Repository {
	return &Repository{
		Company:  newCompany(),
		Customer: newCustomer(),
		General:  newGeneral(),
	}
}
