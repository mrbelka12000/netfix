package repository

import "github.com/mrbelka12000/netfix/basic/models"

type Company interface {
	CreateWork(work *models.Work) error
	GetWorkStatus(workID int) (bool, error)
}

type Customer interface {
	ApplyForWork(apply *models.WorkActions) error
	FinishWork(work *models.WorkActions) error
}

type General interface{}

type Work interface {
	GetWorkFields() (*models.WorkFields, error)
	IsExists(workField string) bool
	GetByID(id int) (*models.Work, error)
}

type Repository struct {
	Company
	Customer
	General
	Work
}

func NewRepo() *Repository {
	return &Repository{
		Company:  newCompany(),
		Customer: newCustomer(),
		General:  newGeneral(),
		Work:     newWorks(),
	}
}
