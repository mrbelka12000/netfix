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

	res, err := conn.Exec(`
		UPDATE
			apply
		SET 
		    enddate=$1
		WHERE
		    workId = $2 and customerId= $3
`, tools.GetUnixDate(), work.WorkID, work.CustomerID)
	if err != nil {
		log.Println("finish work error: " + err.Error())
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Println("rows affected error: " + err.Error())
		return err
	}

	if rows == 0 {
		log.Println("no works to finish")
		return errors.New("no works to finish")
	}

	return nil
}
