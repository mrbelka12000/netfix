package service

import (
	"github.com/mrbelka12000/netfix/auth/internal/repository"
	"github.com/mrbelka12000/netfix/auth/models"
)

type srvCompany struct {
	repo *repository.Repository
}

func newCompany(repo *repository.Repository) *srvCompany {
	return &srvCompany{repo}
}

func (sc *srvCompany) RegisterCompany(company *models.Company) error {
	return sc.repo.RegisterCompany(company)
}
