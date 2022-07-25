package service

import (
	"github.com/mrbelka12000/netfix/basic/internal/repository"
	"github.com/mrbelka12000/netfix/basic/models"
)

type srvWorks struct {
	repo *repository.Repository
}

func newWorks(repo *repository.Repository) *srvWorks {
	return &srvWorks{repo}
}

func (sw *srvWorks) GetWorkFields() (*models.WorkFields, error) {
	return sw.repo.GetWorkFields()
}