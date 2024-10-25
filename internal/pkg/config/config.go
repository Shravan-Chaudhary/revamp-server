package config

import (
	"flag"
	"log"
	"os"

	"github.com/Shravan-Chaudhary/revamp-server/internal/pkg/types"
	"github.com/ilyakaznacheev/cleanenv"
)


func MustLoad() *types.Config {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "" , "path to config file")
		flag.Parse()

		configPath = *flags
	}

	// Check if path exists
	if configPath == "" {
		log.Fatal("Config path is required")
	}

	// Check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file not found %s", configPath)
	}

	var cfg types.Config

	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
		log.Fatalf("Failed to read config: %s", err.Error())
	}

	return &cfg
}