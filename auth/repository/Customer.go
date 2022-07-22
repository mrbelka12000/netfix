package repository

import (
	"errors"
	"github.com/mrbelka12000/netfix/auth/models"
	"log"
	"time"
)

type repoCustomer struct{}

func newCustomer() *repoCustomer {
	return &repoCustomer{}
}

const layout = "2006-01-02"

func (rc *repoCustomer) RegisterCustomer(customer *models.Customer) error {
	conn := GetConnection()

	date, err := time.Parse(layout, customer.Birth)
	if err != nil {
		log.Println("invalid birth date: " + err.Error())
		return errors.New("invalid birth date")
	}

	_, err = conn.Exec(`
	INSERT INTO customer
		(id, birth)
	VALUES 
		($1,$2)
`, customer.ID, date)
	if err != nil {
		log.Println("customer register error: " + err.Error())
		return err
	}

	return nil
}
