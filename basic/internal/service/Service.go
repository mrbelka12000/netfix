package service

import (
	"github.com/mrbelka12000/netfix/basic/internal/repository"
	"github.com/mrbelka12000/netfix/basic/models"
)

type Company interface {
	CreateWork(work *models.Work) error
	GetWorkStatus(workID int) (bool, error)
}

type Customer interface {
	ApplyForWork(work *models.WorkActions) error
	FinishWork(work *models.WorkActions) error
}

type General interface{}

type Work interface {
	GetWorkFields() (*models.WorkFields, error)
	IsExists(workField string) bool
	GetByID(id int) (*models.Work, error)
}

type Service struct {
	Company
	Customer
	General
	Work
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Company:  newCompany(repo),
		Customer: newCustomer(repo),
		General:  newGeneral(repo),
		Work:     newWorks(repo),
	}
}
