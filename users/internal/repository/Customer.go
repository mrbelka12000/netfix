package repository

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/mrbelka12000/netfix/users/models"
)

type repoCustomer struct{}

func newCustomer() *repoCustomer {
	return &repoCustomer{}
}

const layout = "2006-01-02"

func (rc *repoCustomer) RegisterCustomer(customer *models.Customer, tx *sql.Tx) error {

	date, err := time.Parse(layout, customer.Birth)
	if err != nil {
		log.Println("invalid birth date: " + err.Error())
		tx.Rollback()
		return errors.New("invalid birth date")
	}

	_, err = tx.Exec(`
	INSERT INTO customer
		(id, birth)
	VALUES 
		($1,$2)
`, customer.ID, date)
	if err != nil {
		tx.Rollback()
		log.Println("customer register error: " + err.Error())
		return err
	}

	log.Println("customer successfully created")
	return nil
}

func (rc *repoCustomer) GetByID(id int) (*models.General, error) {
	conn := GetConnection()
	comp := &models.General{}

	err := conn.QueryRow(`
		SELECT 
		    customer.birth, general.id, general.username, general.email
		FROM
			customer join general on customer.id=general.id
		WHERE
		    customer.id=$1
`, id).Scan(&comp.Birth, &comp.ID, &comp.Username, &comp.Email)
	if err != nil {
		log.Println("get company error: " + err.Error())
		return nil, err
	}

	return comp, nil
}
