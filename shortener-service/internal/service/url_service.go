package service

import (
	"shortener-service/internal/model"
	"shortener-service/internal/repository"
	"shortener-service/internal/utility"
)

type ShortenerService struct {
	repo *repository.ShortenerRepo
}

func NewShortenerService(repo *repository.ShortenerRepo) *ShortenerService {
	return &ShortenerService{repo: repo}
}

func (s *ShortenerService) CreateShortURL(originUrl string, sessionId string) (*model.URL, error) {
	shortCode := utility.GenerateShortCode()

	url := &model.URL{OriginURL: originUrl, ShortCode: shortCode, SessionID: sessionId}
	if err := s.repo.CreateShortURL(url); err != nil {
		return nil, err
	}

	return url, nil
}

func (s *ShortenerService) GetOriginURLs(sessionId string) (*[]model.URL, error) {
	return s.repo.GetOriginURLs(sessionId)
}
