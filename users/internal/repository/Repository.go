package repository

import (
	"database/sql"

	"github.com/mrbelka12000/netfix/users/models"
)

type Company interface {
	RegisterCompany(company *models.Company, tx *sql.Tx) error
	GetByID(id int) (*models.General, error)
}

type Customer interface {
	RegisterCustomer(customer *models.Customer, tx *sql.Tx) error
	GetByID(id int) (*models.General, error)
}

type General interface {
	Register(general *models.General, tx *sql.Tx) (int, error)
	Login(l *models.Login) error
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
