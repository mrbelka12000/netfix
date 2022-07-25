package main

import (
	"github.com/mrbelka12000/netfix/basic/config"
	"github.com/mrbelka12000/netfix/basic/internal/app"
)

func main() {
	config.GetConf()
	app.Initialize()
}
