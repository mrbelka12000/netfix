package service

import (
	"errors"
	"log"

	"github.com/mrbelka12000/netfix/auth/models"
	"github.com/mrbelka12000/netfix/auth/repository"
)

type srvGeneral struct {
	repo *repository.Repository
}

func newGeneral(repo *repository.Repository) *srvGeneral {
	return &srvGeneral{repo}
}

func (sg *srvGeneral) Register(general *models.General, user string) (int, error) {
	id, err := sg.repo.Register(general)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	switch user {
	case models.Cmp:
		err := sg.repo.Company.RegisterCompany(&models.Company{ID: id})
		if err != nil {
			log.Println(err)
			return 0, err
		}
	case models.Cust:
		err := sg.repo.Customer.RegisterCustomer(&models.Customer{ID: id})
		if err != nil {
			log.Println(err)
			return 0, err
		}
	default:
		log.Println("invalid user type")
		return 0, errors.New("invalid user type")
	}
	return id, nil
}
