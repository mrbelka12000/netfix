package repository

import (
	"log"

	"github.com/mrbelka12000/netfix/users/models"
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

func (rc *repoCompany) GetByID(id int) (*models.General, error) {
	conn := GetConnection()
	comp := &models.General{}

	err := conn.QueryRow(`
		SELECT 
		    company.workfield,general.id, general.username, general.email
		FROM
			company join general on company.id=general.id
		WHERE
		    company.id=$1
`, id).Scan(&comp.WorkField, &comp.ID, &comp.Username, &comp.Email)
	if err != nil {
		log.Println("get company error: " + err.Error())
		return nil, err
	}

	return comp, nil
}
