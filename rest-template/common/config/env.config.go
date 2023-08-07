package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	if err := godotenv.Load(".env"); err != nil {
		panic("Cannot find .env file.")
	}

	return os.Getenv(key)
}
