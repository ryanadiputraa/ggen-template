package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	PostgresDSN string
}

func LoadConfig() (config Config, err error) {
	if err = godotenv.Load(); err != nil {
		return
	}

	config = Config{
		Port:        os.Getenv("PORT"),
		PostgresDSN: os.Getenv("POSTGRES_DSN"),
	}
	return
}
