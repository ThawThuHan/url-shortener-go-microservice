package grpc

import (
	"context"
	"redirection-service/internal/model"
	"redirection-service/internal/service"
	pb "redirection-service/proto"
)

type RedirectionServer struct {
	pb.UnimplementedRedirectionServiceServer
	urlService       *service.URLService
	accessLogService *service.AccessLogService
}

func NewRedirectionServer(urlService *service.URLService, accessLog *service.AccessLogService) *RedirectionServer {
	return &RedirectionServer{urlService: urlService, accessLogService: accessLog}
}

func (s *RedirectionServer) GetOriginURL(ctx context.Context, req *pb.GetOriginURLRequest) (*pb.GetOriginURLResponse, error) {
	url, err := s.urlService.GetOriginURL(ctx, req.ShortCode)
	if err != nil || url == nil {
		return nil, err
	}
	accessLog := model.AccessLog{ShortURLID: url.ID, Location: req.Location, IPAddress: req.IpAddress, City: req.City}
	if err := s.accessLogService.CreateAccessLog(&accessLog); err != nil {
		return nil, err
	}
	return &pb.GetOriginURLResponse{OriginUrl: url.OriginURL}, nil
}

func (s *RedirectionServer) GetAccessLog(ctx context.Context, req *pb.GetAccessLogRequest) (*pb.GetAccessLogResponse, error) {
	url, err := s.urlService.GetOriginURL(ctx, req.ShortCode)
	if err != nil {
		return nil, err
	}

	accessLogs, err := s.accessLogService.GetAccessLogsByURLID(url.ID)
	if err != nil {
		return nil, err
	}

	var pbLogs []*pb.AccessLog
	for _, accessLog := range accessLogs {
		pbLogs = append(pbLogs, &pb.AccessLog{
			Id:         int64(accessLog.ID),
			ShortUrlId: int64(accessLog.ShortURLID),
			AccessTime: accessLog.CreatedAt.Format("2006-01-02 15:04:05"),
			IpAddress:  accessLog.IPAddress,
			Location:   accessLog.Location,
			City:       accessLog.City,
		})
	}

	return &pb.GetAccessLogResponse{Logs: pbLogs}, nil
}
