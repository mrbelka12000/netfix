package repository

import "github.com/mrbelka12000/netfix/basic/models"

type Company interface {
	CreateWork(work *models.CreateWork) error
}

type Customer interface {
}

type General interface {
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
