package repository

import (
	"github.com/mrbelka12000/netfix/basic/models"
	"log"
)

type repoCustomer struct {
}

func newCustomer() *repoCustomer {
	return &repoCustomer{}
}

func (rc *repoCustomer) ApplyForWork(apply *models.ApplyForWork) error {
	conn := GetConnection()

	tx, err := conn.Begin()
	if err != nil {
		log.Println("tx create error: " + err.Error())
		return err
	}
	defer tx.Commit()
	_, err = tx.Exec(`
	INSERT INTO apply
		(customerid, workID)
	VALUES 
		($1,$2)
`, apply.CustomerID, apply.WorkID)
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
`, true, apply.WorkID)
	if err != nil {
		tx.Rollback()
		log.Println("apply for work error: " + err.Error())
		return err
	}

	log.Println("successfully applied for work")
	return nil
}
