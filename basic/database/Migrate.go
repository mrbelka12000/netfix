package database

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/mrbelka12000/netfix/basic/config"
	"github.com/mrbelka12000/netfix/basic/internal/repository"
)

func Up() {
	cfg := config.GetConf()
	dirName := cfg.App.SchemaUp
	dir, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}

	conn := repository.GetConnection()
	for _, file := range dir {
		body, err := ioutil.ReadFile(dirName + "/" + file.Name())
		if err != nil {
			log.Fatal("Cant read file: ", err)
		}

		if _, err = conn.Exec(string(body)); err != nil {
			log.Println(fmt.Sprintf("Миграция %v не может отработать по причине %v", file.Name(), err.Error()))
			continue
		}
		log.Println(fmt.Sprintf("Миграция %v отработала успешно ", file.Name()))
	}
}
