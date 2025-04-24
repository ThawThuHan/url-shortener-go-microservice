package repository

import (
	"shortener-service/internal/model"

	"gorm.io/gorm"
)

type ShortenerRepo struct {
	db *gorm.DB
}

func NewShortenerRepo(db *gorm.DB) *ShortenerRepo {
	return &ShortenerRepo{db: db}
}

func (s *ShortenerRepo) CreateShortURL(url *model.URL) error {
	return s.db.Create(url).Error
}

func (s *ShortenerRepo) GetOriginURLs(sessionId string) (*[]model.URL, error) {
	var urls []model.URL
	err := s.db.Where("session_id = ?", sessionId).Find(&urls).Error
	return &urls, err
}
