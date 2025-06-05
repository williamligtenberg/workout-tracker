package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser string
	DBPass string
	DBName string
	DBHost string
	DBPort string
}

func LoadConfig() (Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return Config{}, fmt.Errorf("error loading .env: %w", err)
	}

	return Config{
		DBUser: os.Getenv("DBUSER"),
		DBPass: os.Getenv("DBPASS"),
		DBName: os.Getenv("DBNAME"),
		DBHost: "127.0.0.1",
		DBPort: "3306",
	}, nil
}
