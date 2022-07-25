package main

import (
	"github.com/mrbelka12000/netfix/auth/config"
	"github.com/mrbelka12000/netfix/auth/database"
	"github.com/mrbelka12000/netfix/auth/internal/app"
)

func main() {
	database.Up()
	config.GetConf()
	app.Initialize()
}
