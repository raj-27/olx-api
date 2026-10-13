package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	Env          string
	DATABASE_URL string
}

// MustLoad loads config from .env file (via godotenv) or system environment variables.
// Panics if any required value is missing.
func MustLoad() Config {
	godotenv.Load() // reads .env file into os environment

	port := os.Getenv("PORT")
	if port == "" {
		panic("Port is required.")
	}

	env := os.Getenv("ENV")
	if env == "" {
		panic("ENV is required.")
	}

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		panic("DATABASE_URL is required.")
	}

	return Config{
		Port:         port,
		Env:          env,
		DATABASE_URL: dbUrl,
	}
}
