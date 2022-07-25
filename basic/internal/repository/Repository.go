package repository

import "github.com/mrbelka12000/netfix/basic/models"

type Company interface {
	CreateWork(work *models.CreateWork) error
	GetWorkStatus(workID int) (bool, error)
}

type Customer interface {
	ApplyForWork(apply *models.ApplyForWork) error
}

type General interface{}

type WorkFields interface {
	GetWorkFields() (*models.WorkFields, error)
}

type Repository struct {
	Company
	Customer
	General
	WorkFields
}

func NewRepo() *Repository {
	return &Repository{
		Company:    newCompany(),
		Customer:   newCustomer(),
		General:    newGeneral(),
		WorkFields: newWorks(),
	}
}
