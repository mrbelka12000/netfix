package service

import (
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

func (sg *srvGeneral) Register(general *models.General) (int, error) {
	id, err := sg.repo.Register(general)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return id, nil
}
