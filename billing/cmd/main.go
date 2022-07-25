package main

import (
	"github.com/mrbelka12000/netfix/billing/config"
	"github.com/mrbelka12000/netfix/billing/internal/app"
)

func main() {
	config.GetConf()
	app.Initialize()
}
