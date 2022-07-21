package main

import (
	"github.com/mrbelka12000/netfix/auth/app"
	"github.com/mrbelka12000/netfix/auth/config"
)

func main() {
	cfg := config.GetConf()
	app.Initialize(cfg)
}
