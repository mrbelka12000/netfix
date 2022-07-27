package repository

import (
	"log"

	"github.com/mrbelka12000/netfix/basic/models"
)

type repoWorks struct{}

func newWorks() *repoWorks {
	return &repoWorks{}
}

func (rw *repoWorks) GetWorkFields() (*models.WorkFields, error) {
	wf := &models.WorkFields{}

	conn := GetConnection()

	rows, err := conn.Query(`
		SELECT 
			WORKFIELD
		FROM
			WORKFIELDS
	`)
	if err != nil {
		log.Println("get workfields error: " + err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		workField := ""
		err = rows.Scan(&workField)
		if err != nil {
			log.Println("scan error: " + err.Error())
			return nil, err
		}
		wf.WorkFileds = append(wf.WorkFileds, workField)
	}

	err = rows.Err()
	if err != nil {
		log.Println("hz oshibka: " + err.Error())
		return nil, err
	}
	return wf, nil
}

func (rw *repoWorks) IsExists(workField string) bool {
	conn := GetConnection()

	check := ""
	err := conn.QueryRow(`
	SELECT 
		WORKFIELD
	FROM
		WORKFIELDS
	WHERE
		workfield = $1
`, workField).Scan(&check)
	if err != nil {
		log.Println("can not find work field: " + err.Error())
		return false
	}

	return workField == check
}

func (rw *repoWorks) GetByID(id int) (*models.Work, error) {
	conn := GetConnection()
	w := &models.Work{}
	err := conn.QueryRow(`
	SELECT
	    name, workfield, description, price, date,CompanyID 
	FROM 
	    works
	WHERE
	    Id =$1
`, id).Scan(&w.Name, &w.WorkField, &w.Description, &w.Price, &w.Date, &w.CompanyID)
	if err != nil {
		log.Println("get work error: " + err.Error())
		return nil, err
	}

	return w, nil
}
