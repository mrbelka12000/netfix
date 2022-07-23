package service

import (
	"github.com/mrbelka12000/netfix/basic/models"
	"github.com/mrbelka12000/netfix/basic/repository"
)

type Company interface {
	CreateWork(work *models.CreateWork) error
}

type Customer interface {
}

type General interface {
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
