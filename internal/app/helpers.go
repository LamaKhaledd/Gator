package app

import (
	"fmt"

	"gator/internal/config"
)

func LoadConfig() (config.Config, error) {
	return config.Read()
}

func SetAndSaveUser(cfg *config.Config, username string) error {
	return cfg.SetUser(username)
}

func PrintConfig(cfg config.Config) {
	fmt.Printf("Current Config: %+v\n", cfg)
}
