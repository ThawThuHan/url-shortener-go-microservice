package grpc

import (
	"context"
	"errors"
	"shortener-service/configs"
	"shortener-service/internal/service"
	"shortener-service/internal/utility"
	pb "shortener-service/proto"

	"time"

	"github.com/go-playground/validator/v10"
)

type ShortenerServer struct {
	pb.UnimplementedShortenerServiceServer
	service   *service.ShortenerService
	validator *validator.Validate
}

func NewShortenerServer(service *service.ShortenerService, validator *validator.Validate) *ShortenerServer {
	return &ShortenerServer{service: service, validator: validator}
}

func (s *ShortenerServer) CreateShortURL(ctx context.Context, req *pb.CreateShortURLRequest) (*pb.CreateShortURLResponse, error) {
	err := s.validator.Var(req.OriginUrl, "required,url")
	if err != nil {
		return nil, err
	}
	if utility.IsOwnDomain(req.OriginUrl, configs.Cfg.Domain) {
		return nil, errors.New("cannot short own url")
	}
	url, err := s.service.CreateShortURL(req.OriginUrl, req.SessionId)
	if err != nil {
		return nil, err
	}

	return &pb.CreateShortURLResponse{ShortCode: url.ShortCode, OriginUrl: url.OriginURL, CreatedAt: url.CreatedAt.Format(time.RFC3339)}, nil
}

func (s *ShortenerServer) GetOriginURLs(ctx context.Context, req *pb.GetOriginURLsRequest) (*pb.GetOriginURLsResponse, error) {
	urls, err := s.service.GetOriginURLs(req.SessionId)
	if err != nil {
		return nil, err
	}

	var response []*pb.Urls
	for _, url := range *urls {
		response = append(response, &pb.Urls{
			Id:        int64(url.ID),
			SessionId: url.SessionID,
			OriginUrl: url.OriginURL,
			ShortCode: url.ShortCode,
			CreatedAt: url.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &pb.GetOriginURLsResponse{Urls: response}, nil
}
