package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host       string
	Port       string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	Domain     string
}

var Cfg *Config

func load() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Error loading .env file")
	}

	var config Config
	config.Host = os.Getenv("HOST")
	if config.Host == "" {
		config.Host = "localhost"
	}
	config.Port = os.Getenv("PORT")
	if config.Port == "" {
		config.Port = "8080"
	}
	config.DBHost = os.Getenv("DB_HOST")
	if config.DBHost == "" {
		config.DBHost = "localhost"
	}
	config.DBPort = os.Getenv("DB_PORT")
	if config.DBPort == "" {
		config.DBPort = "5432"
	}
	config.DBUser = os.Getenv("DB_USER")
	if config.DBUser == "" {
		config.DBUser = "postgres"
	}
	config.DBPassword = os.Getenv("DB_PASSWORD")
	if config.DBPassword == "" {
		config.DBPassword = "postgres"
	}
	config.DBName = os.Getenv("DB_NAME")
	if config.DBName == "" {
		config.DBName = "postgres"
	}
	config.Domain = os.Getenv("DOMAIN")
	if config.Domain == "" {
		config.Domain = "localhost"
	}

	return &config, nil
}

func init() {
	var err error
	Cfg, err = load()
	if err != nil {
		log.Fatal(err)
	}
}
