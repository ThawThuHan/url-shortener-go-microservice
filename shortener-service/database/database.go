package database

import (
	"log"
	"shortener-service/configs"
	"shortener-service/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase(cfg *configs.Config) *Database {
	dsn := "host=" + cfg.DBHost + " user=" + cfg.DBUser + " password=" + cfg.DBPassword + " dbname=" + cfg.DBName + " port=" + cfg.DBPort + " sslmode=disable timezone=Asia/Bangkok"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return &Database{DB: db}
}

func (d *Database) Migrate() error {
	return d.DB.AutoMigrate(&model.URL{})
}
