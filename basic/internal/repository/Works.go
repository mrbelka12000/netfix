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
