package service

import (
	"github.com/mrbelka12000/netfix/basic/models"
	"github.com/mrbelka12000/netfix/basic/repository"
)

type srvCompany struct {
	repo *repository.Repository
}

func newCompany(repo *repository.Repository) *srvCompany {
	return &srvCompany{repo}
}

func (sc *srvCompany) CreateWork(work *models.CreateWork) error {
	return sc.repo.CreateWork(work)
}

func (sc *srvCompany) GetWorkStatus(workID int) (bool, error) {
	return sc.repo.GetWorkStatus(workID)
}
