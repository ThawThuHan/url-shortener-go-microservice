package service

import (
	"redirection-service/internal/model"
	"redirection-service/internal/repository"
)

type AccessLogService struct {
	repo *repository.AccessLogRepo
}

func NewAccessLogService(repo *repository.AccessLogRepo) *AccessLogService {
	return &AccessLogService{repo: repo}
}

func (s *AccessLogService) CreateAccessLog(accessLog *model.AccessLog) error {
	return s.repo.CreateAccessLog(accessLog)
}

func (s *AccessLogService) GetAccessLogsByURLID(shortURLID uint) ([]model.AccessLog, error) {
	return s.repo.GetAccessLogByURLID(shortURLID)
}
