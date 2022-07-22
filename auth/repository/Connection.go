package repository

import (
	"database/sql"
	"fmt"
	"github.com/mrbelka12000/netfix/auth/config"
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
	cfg := config.GetConf()
	connStrForDocker := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		cfg.Postgres.POSTGRES_USER, cfg.Postgres.POSTGRES_PASSWORD,
		cfg.Postgres.POSTGRES_HOST, cfg.Postgres.POSTGRES_PORT,
		cfg.Postgres.POSTGRES_DB)
	return connStrForDocker
}
