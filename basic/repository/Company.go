package repository

import (
	"github.com/mrbelka12000/netfix/basic/models"
	"log"
	"time"
)

type repoCompany struct {
}

func newCompany() *repoCompany {
	return &repoCompany{}
}

func (rc *repoCompany) CreateWork(work *models.CreateWork) error {

	conn := GetConnection()

	loc, _ := time.LoadLocation("Asia/Almaty")

	now := time.Now().In(loc)
	_, err := conn.Exec(`
		INSERT INTO works
			(name, workfield, description, price, date, companyID)
		VALUES 
			($1,$2,$3,$4,$5,$6)
	`, work.Name, work.WorkField, work.Description, work.Price, now, work.CompanyID)

	if err != nil {
		log.Println("error creating work: " + err.Error())
		return err
	}

	log.Println("work successfully created")
	return nil
}
