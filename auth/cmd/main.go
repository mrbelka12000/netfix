package main

import (
	"github.com/mrbelka12000/netfix/auth/app"
	"github.com/mrbelka12000/netfix/auth/database"
)

func main() {
	database.Up()
	app.Initialize()
}
