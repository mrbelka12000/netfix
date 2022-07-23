package main

import (
	"github.com/mrbelka12000/netfix/basic/app"
	"github.com/mrbelka12000/netfix/basic/database"
)

func main() {
	database.Up()
	app.Initialize()
}
