package main

import (
	"github.com/mrbelka12000/netfix/basic/config"
	"github.com/mrbelka12000/netfix/basic/database"
	"github.com/mrbelka12000/netfix/basic/internal/app"
)

func main() {
	database.Up()
	config.GetConf()
	app.Initialize()
}
