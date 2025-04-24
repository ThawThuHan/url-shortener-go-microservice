package repository

import (
	"redirection-service/internal/model"

	"gorm.io/gorm"
)

type AccessLogRepo struct {
	db *gorm.DB
}

func NewAccessLogRepo(db *gorm.DB) *AccessLogRepo {
	return &AccessLogRepo{db: db}
}

func (a *AccessLogRepo) CreateAccessLog(accessLog *model.AccessLog) error {
	return a.db.Create(&accessLog).Error
}

func (a *AccessLogRepo) GetAccessLogByURLID(shortURLID uint) ([]model.AccessLog, error) {
	var accessLogs []model.AccessLog
	if err := a.db.Where("short_url_id = ?", shortURLID).Find(&accessLogs).Error; err != nil {
		return nil, err
	}
	return accessLogs, nil
}
