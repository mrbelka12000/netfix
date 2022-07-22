package repository

import (
	"github.com/mrbelka12000/netfix/auth/models"
	"log"
)

type repoGeneral struct{}

func newGeneral() *repoGeneral {
	return &repoGeneral{}
}

func (ng *repoGeneral) Register(gen *models.General) (int, error) {
	conn := GetConnection()

	err := conn.QueryRow(`
	INSERT INTO general
		(email, password, username)
	VALUES
		($1,$2,$3)
	RETURNING
		id
`, gen.Email, gen.Password, gen.Username).Scan(&gen.ID)
	if err != nil {
		log.Println("General register error: " + err.Error())
		return 0, err
	}

	return gen.ID, nil
}
