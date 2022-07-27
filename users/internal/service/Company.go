package service

import (
	"database/sql"
	"github.com/mrbelka12000/netfix/users/internal/repository"
	"github.com/mrbelka12000/netfix/users/models"
)

type srvCompany struct {
	repo *repository.Repository
}

func newCompany(repo *repository.Repository) *srvCompany {
	return &srvCompany{repo}
}

func (sc *srvCompany) RegisterCompany(company *models.Company, tx *sql.Tx) error {
	return sc.repo.RegisterCompany(company, tx)
}

func (sc *srvCompany) GetByID(id int) (*models.General, error) {
	return sc.repo.Company.GetByID(id)
}
