package main

import (
	"log"
	"magazine/internal/app"
	"magazine/internal/config"
)

func main() {
	log.Println("Config initializing...")
	config := config.MustLoad()

	app := app.NewApp(config)
	app.Router.Run()
}
