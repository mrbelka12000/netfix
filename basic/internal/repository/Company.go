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

	tx, err := conn.Begin()
	if err != nil {
		log.Println("tx create error: " + err.Error())
		return err
	}
	defer tx.Commit()
	now := time.Now().In(loc)
	err = tx.QueryRow(`
		INSERT INTO works
			(name, workfield, description, price, date, companyID)
		VALUES 
			($1,$2,$3,$4,$5,$6)
		RETURNING id
	`, work.Name, work.WorkField, work.Description, work.Price, now, work.CompanyID).Scan(&work.ID)

	if err != nil {
		tx.Rollback()
		log.Println("error creating work: " + err.Error())
		return err
	}

	_, err = tx.Exec(`
	INSERT INTO workstatus
		(workid, status)
	VALUES 
		($1,$2)
`, work.ID, false)
	if err != nil {
		tx.Rollback()
		log.Println("work status create error: " + err.Error())
		return err
	}
	log.Println("work successfully created")
	return nil
}

func (rc *repoCompany) GetWorkStatus(workID int) (bool, error) {
	conn := GetConnection()
	var status bool
	err := conn.QueryRow(`
	SELECT 
	    Status
	FROM 
	    workstatus
	WHERE
	    Workid=$1
`, workID).Scan(&status)
	if err != nil {
		log.Println("get status error: " + err.Error())
		return false, err
	}

	return status, nil
}
