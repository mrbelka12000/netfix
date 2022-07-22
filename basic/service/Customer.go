package service

import "github.com/mrbelka12000/netfix/basic/repository"

type srvCustomer struct {
	repo *repository.Repository
}

func newCustomer(repo *repository.Repository) *srvCustomer {
	return &srvCustomer{repo}
}
