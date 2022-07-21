package repository

import "github.com/mrbelka12000/netfix/auth/models"

type repoCompany struct{}

func newCompany() *repoCompany {
	return &repoCompany{}
}

func (rc *repoCompany) RegisterCompany(company *models.Company) error {
	return nil
}
