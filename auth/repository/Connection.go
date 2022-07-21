package repository

import (
	"database/sql"
	"fmt"
	"os"
)

var (
	alive bool
	conn  *sql.DB
)

func GetConnection() *sql.DB {
	if alive {
		return conn
	}
	conn = connectToDB()
	if !alive {
		panic("where is my connection!!!!!!!!!")
	}
	return conn
}

func connectToDB() *sql.DB {
	db, err := sql.Open("postgres", getConnectionString())
	if err != nil {
		return nil
	}

	err = db.Ping()
	if err != nil {
		return nil
	}

	alive = true
	return db
}

func getConnectionString() string {
	connStrForHeroku := os.Getenv("DATABASE_URL")
	if connStrForHeroku != "" {
		return connStrForHeroku
	}

	connStrForDocker := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"))
	return connStrForDocker
}
