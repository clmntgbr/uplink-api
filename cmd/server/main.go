package main

import (
	"log"
	"uplink-api/config"
	"uplink-api/internal/app"
)

func main() {
	cfg := config.Load()
	
	application := app.New(cfg)
	
	log.Fatal(application.Start(":3000"))
}