package service

import (
	"github.com/mrbelka12000/netfix/basic/internal/repository"
	"github.com/mrbelka12000/netfix/basic/models"
)

type srvCustomer struct {
	repo *repository.Repository
}

func newCustomer(repo *repository.Repository) *srvCustomer {
	return &srvCustomer{repo}
}

func (sc *srvCustomer) ApplyForWork(work *models.WorkActions) error {
	return sc.repo.ApplyForWork(work)
}

func (sc *srvCustomer) FinishWork(work *models.WorkActions) error {
	return sc.repo.FinishWork(work)
}
