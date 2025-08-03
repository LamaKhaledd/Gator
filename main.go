package main

import (
	"log"

	"gator/internal/app"
)

func main() {
	cfg, err := app.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if err := app.SetAndSaveUser(&cfg, "Lama"); err != nil {
		log.Fatalf("Failed to set user: %v", err)
	}

	app.PrintConfig(cfg)
}
