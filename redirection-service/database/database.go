package database

import (
	"redirection-service/configs"
	"redirection-service/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg *configs.Config) (*gorm.DB, error) {
	dsn := "host=" + cfg.DB_HOST + " user=" + cfg.DB_USER + " password=" + cfg.DB_PASS + " dbname=" + cfg.DB_NAME + " port=" + cfg.DB_PORT + " sslmode=disable timezone=Asia/Bangkok"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.AccessLog{})
}
