package repository

import (
	"errors"
	"log"

	"github.com/mrbelka12000/netfix/basic/models"
	"github.com/mrbelka12000/netfix/basic/tools"
)

type repoCustomer struct{}

func newCustomer() *repoCustomer {
	return &repoCustomer{}
}

const active = false

func (rc *repoCustomer) ApplyForWork(work *models.WorkActions) error {
	conn := GetConnection()

	tx, err := conn.Begin()
	if err != nil {
		log.Println("tx create error: " + err.Error())
		return err
	}
	defer tx.Commit()
	_, err = tx.Exec(`
	INSERT INTO apply
		(customerid, workID, Startdate)
	VALUES 
		($1,$2,$3)
`, work.CustomerID, work.WorkID, tools.GetUnixDate())
	if err != nil {
		tx.Rollback()
		log.Println("apply for work error: " + err.Error())
		return err
	}

	_, err = tx.Exec(`
		update
			workstatus
		set
			status=$1
		where
			workid=$2
`, true, work.WorkID)
	if err != nil {
		tx.Rollback()
		log.Println("apply for work error: " + err.Error())
		return err
	}

	log.Println("successfully applied for work")
	return nil
}

func (rc *repoCustomer) FinishWork(work *models.WorkActions) error {
	conn := GetConnection()
	now := tools.GetUnixDate()

	tx, err := conn.Begin()
	if err != nil {
		log.Println("tx creation error: " + err.Error())
		return err
	}
	defer tx.Commit()

	err = tx.QueryRow(`
		UPDATE
			apply
		SET 
		    enddate=$1
		WHERE
		    workId = $2 and customerId= $3
		RETURNING 
			ID, startdate
`, now, work.WorkID, work.CustomerID).Scan(&work.ID, &work.StartDate)
	if err != nil {
		tx.Rollback()
		log.Println("repo finish work error: " + err.Error())
		return err
	}

	if work.ID == 0 {
		tx.Rollback()
		return errors.New("no items to update")
	}

	_, err = tx.Exec(`
		UPDATE 
			workstatus
		SET
		    status=$1
		WHERE 
		    workid=$2
`, active, work.WorkID)
	if err != nil {
		log.Println("update work status error: " + err.Error())
		tx.Rollback()
		return err
	}

	work.EndDate = now
	return nil
}
