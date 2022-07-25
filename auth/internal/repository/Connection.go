package repository

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
	"github.com/mrbelka12000/netfix/auth/config"
)

var (
	conn *sql.DB
	once sync.Once
)

// GetConnection singleton implementation.
func GetConnection() *sql.DB {
	once.Do(func() {
		conn = connectToDB()
		if conn == nil {
			panic("where is my connection!!!!!!!!!")
		}
	})

	return conn
}

func connectToDB() *sql.DB {
	db, err := sql.Open("postgres", getConnectionString())
	if err != nil {
		log.Println("db open error: " + err.Error())
		return nil
	}

	err = db.Ping()
	if err != nil {
		log.Println("ping error: " + err.Error())
		return nil
	}

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
