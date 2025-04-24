package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host               string
	Port               string
	ShortenerService   string
	RedirectionService string
	FrontendURL        string
}

var Env *Config

func loadEnv() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	var cfg Config

	cfg.Host = os.Getenv("HOST")
	if cfg.Host == "" {
		cfg.Host = "localhost"
	}

	cfg.Port = os.Getenv("PORT")
	if cfg.Port == "" {
		cfg.Port = "8080"
	}

	cfg.ShortenerService = os.Getenv("SHORTENER_SERVICE")
	if cfg.ShortenerService == "" {
		cfg.ShortenerService = "localhost:5881"
	}

	cfg.RedirectionService = os.Getenv("REDIRECTION_SERVICE")
	if cfg.RedirectionService == "" {
		cfg.RedirectionService = "localhost:5882"
	}

	cfg.FrontendURL = os.Getenv("FRONTEND_URL")
	if cfg.FrontendURL == "" {
		cfg.FrontendURL = "http://localhost:5173"
	}

	return &cfg
}

func init() {
	Env = loadEnv()
}
