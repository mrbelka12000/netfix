package main

import (
	"github.com/mrbelka12000/netfix/users/database"
	"github.com/mrbelka12000/netfix/users/internal/app"
)

func main() {
	database.Up()
	app.Initialize()
}
