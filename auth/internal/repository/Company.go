package repository

import (
	"log"

	"github.com/mrbelka12000/netfix/auth/models"
)

type repoCompany struct{}

func newCompany() *repoCompany {
	return &repoCompany{}
}

func (rc *repoCompany) RegisterCompany(company *models.Company) error {
	conn := GetConnection()

	_, err := conn.Exec(`
	INSERT INTO company
		(id, workfield)
	VALUES 
		($1,$2)
	`, company.ID, company.WorkField)
	if err != nil {
		log.Println("company register error: " + err.Error())
		return err
	}

	log.Println("company successfully created")
	return nil
}