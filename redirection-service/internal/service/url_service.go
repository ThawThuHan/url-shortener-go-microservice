package service

import (
	"context"
	"redirection-service/internal/model"
	"redirection-service/internal/repository"
)

type URLService struct {
	urlRepo *repository.URLRepository
}

func NewURLService(repo *repository.URLRepository) *URLService {
	return &URLService{urlRepo: repo}
}

func (s *URLService) GetOriginURL(ctx context.Context, shortCode string) (*model.URL, error) {
	return s.urlRepo.GetOriginURL(ctx, shortCode)
}
