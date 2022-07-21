package repository

import "github.com/mrbelka12000/netfix/auth/models"

type repoGeneral struct{}

func newGeneral() *repoGeneral {
	return &repoGeneral{}
}

func (ng *repoGeneral) Register(general *models.General) (int, error) {
	return 0, nil
}
