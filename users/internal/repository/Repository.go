package repository

import "github.com/mrbelka12000/netfix/users/models"

type Company interface {
	RegisterCompany(company *models.Company) error
	GetByID(id int) (*models.General, error)
}

type Customer interface {
	RegisterCustomer(customer *models.Customer) error
	GetByID(id int) (*models.General, error)
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
