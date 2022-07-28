package service

import (
	"database/sql"
	"log"

	"github.com/mrbelka12000/netfix/users/internal/repository"
	"github.com/mrbelka12000/netfix/users/models"
	"golang.org/x/crypto/bcrypt"
)

type srvGeneral struct {
	repo *repository.Repository
}

func newGeneral(repo *repository.Repository) *srvGeneral {
	return &srvGeneral{repo}
}

func (sg *srvGeneral) Register(general *models.General, tx *sql.Tx) (int, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(general.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("generate hash error: " + err.Error())
		return 0, err
	}

	general.Password = string(hash)
	return sg.repo.Register(general, tx)
}

func (sg *srvGeneral) Login(l *models.Login) error {
	return sg.repo.Login(l)
}
