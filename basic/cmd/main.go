package main

import (
	"github.com/mrbelka12000/netfix/basic/database"
	"github.com/mrbelka12000/netfix/basic/internal/app"
)

// @title net-fix
// @version 1.0
// @description API for hiring companies by their type of work

// @host localhost:8081
// @BasePath /
// @contact.email karshyga.beknur@gmail.com

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name session
func main() {
	database.Up()
	app.Initialize()
}
