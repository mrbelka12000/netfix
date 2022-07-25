package service

import (
	"github.com/mrbelka12000/netfix/basic/internal/repository"
	"github.com/mrbelka12000/netfix/basic/models"
)

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
	IsExists(workField string) bool
}

type Service struct {
	Company
	Customer
	General
	WorkFields
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Company:    newCompany(repo),
		Customer:   newCustomer(repo),
		General:    newGeneral(repo),
		WorkFields: newWorks(repo),
	}
}
