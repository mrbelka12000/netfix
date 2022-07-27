package repository

import (
	"database/sql"
	"github.com/mrbelka12000/netfix/users/models"
	"log"
)

type repoGeneral struct{}

func newGeneral() *repoGeneral {
	return &repoGeneral{}
}

func (ng *repoGeneral) Register(gen *models.General, tx *sql.Tx) (int, error) {

	err := tx.QueryRow(`
	INSERT INTO general
		(email, password, username)
	VALUES
		($1,$2,$3)
	RETURNING
		id
`, gen.Email, gen.Password, gen.Username).Scan(&gen.ID)
	if err != nil {
		tx.Rollback()
		log.Println("general user register error: " + err.Error())
		return 0, err
	}

	log.Println("general user successfully created")
	return gen.ID, nil
}
