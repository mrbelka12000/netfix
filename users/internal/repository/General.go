package repository

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"

	"github.com/mrbelka12000/netfix/users/models"
)

type repoGeneral struct{}

func newGeneral() *repoGeneral {
	return &repoGeneral{}
}

func (rg *repoGeneral) Register(gen *models.General, tx *sql.Tx) (int, error) {

	err := tx.QueryRow(`
	INSERT INTO general
		(email, password, username)
	VALUES
		($1,$2,$3)
	RETURNING
		id
`, gen.Email, gen.Password, gen.Username).Scan(&gen.ID)
	if err != nil {
		tx.Rollback()
		log.Println("general user register error: " + err.Error())
		return 0, err
	}

	log.Println("general user successfully created")
	return gen.ID, nil
}

func (rg *repoGeneral) Login(l *models.Login) error {

	conn := GetConnection()

	password := ""
	err := conn.QueryRow(`
	SELECT
	    id ,password
	FROM
	    general
	WHERE
	    username = $1 or email = $2
`, l.Credential, l.Credential).Scan(&l.ID, &password)
	if err != nil {
		log.Println("get by username or email error: " + err.Error())
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(l.Password))
	if err != nil {
		log.Println(err.Error())
		return err
	}

	birth := ""
	err = conn.QueryRow(`
	SELECT 
	    birth
	FROM
	    customer
	WHERE
	    ID=$1
`, l.ID).Scan(&birth)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			l.UserType = models.Cmp
			return nil
		}
		log.Println("scan error: " + err.Error())
		return err
	}

	if birth == "" {
		log.Println("this is company")
		l.UserType = models.Cmp
		return nil
	}

	l.UserType = models.Cust
	return nil
}
