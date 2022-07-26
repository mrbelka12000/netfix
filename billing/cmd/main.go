package main

import (
	"github.com/mrbelka12000/netfix/billing/database"
	"github.com/mrbelka12000/netfix/billing/internal/app"
)

func main() {
	database.Up()
	app.Initialize()
}
