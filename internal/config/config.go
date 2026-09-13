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

func MustLoad() Config {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		panic("Port is required.")
	}

	env := os.Getenv("ENV")
	if env == "" {
		panic("ENV is required.")
	}

	// DATABASE_URL
	dbUrl := os.Getenv("DATABASE_URL")
	if env == "" {
		panic("DATABASE_URL is required.")
	}
	return Config{
		Port:         port,
		Env:          env,
		DATABASE_URL: dbUrl,
	}

}
