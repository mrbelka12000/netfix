package service

import "github.com/mrbelka12000/netfix/basic/repository"

type srvCompany struct {
	repo *repository.Repository
}

func newCompany(repo *repository.Repository) *srvCompany {
	return &srvCompany{repo}
}
