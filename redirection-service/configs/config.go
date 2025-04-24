package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HOST           string
	PORT           string
	DB_USER        string
	DB_PASS        string
	DB_NAME        string
	DB_HOST        string
	DB_PORT        string
	Redis_Host     string
	Redis_Port     string
	Redis_Password string
}

var Env *Config

func loadEnv() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	var cfg Config
	cfg.HOST = os.Getenv("HOST")
	if cfg.HOST == "" {
		cfg.HOST = "localhost"
	}

	cfg.PORT = os.Getenv("PORT")
	if cfg.PORT == "" {
		cfg.PORT = "8080"
	}

	cfg.DB_USER = os.Getenv("DB_USER")
	if cfg.DB_USER == "" {
		cfg.DB_USER = "postgres"
	}

	cfg.DB_PASS = os.Getenv("DB_PASS")
	if cfg.DB_PASS == "" {
		cfg.DB_PASS = "postgres"
	}

	cfg.DB_NAME = os.Getenv("DB_NAME")
	if cfg.DB_NAME == "" {
		cfg.DB_NAME = "postgres"
	}

	cfg.DB_HOST = os.Getenv("DB_HOST")
	if cfg.DB_HOST == "" {
		cfg.DB_HOST = "localhost"
	}

	cfg.DB_PORT = os.Getenv("DB_PORT")
	if cfg.DB_PORT == "" {
		cfg.DB_PORT = "5432"
	}

	cfg.Redis_Host = os.Getenv("Redis_HOST")
	if cfg.Redis_Host == "" {
		cfg.Redis_Host = "localhost"
	}

	cfg.Redis_Port = os.Getenv("Redis_PORT")
	if cfg.Redis_Port == "" {
		cfg.Redis_Port = "6379"
	}

	cfg.Redis_Password = os.Getenv("Redis_PASS")

	return &cfg
}

func init() {
	Env = loadEnv()
}
