package service

import "github.com/mrbelka12000/netfix/basic/internal/repository"

type srvGeneral struct {
	repo *repository.Repository
}

func newGeneral(repo *repository.Repository) *srvGeneral {
	return &srvGeneral{repo}
}
