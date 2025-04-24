package grpc

import (
	"api-gateway/config"
	pb "api-gateway/proto/shortener-service"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewShortenerClient(cfg *config.Config) pb.ShortenerServiceClient {
	conn, err := grpc.NewClient(cfg.ShortenerService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error to connect %v", err)
	}
	return pb.NewShortenerServiceClient(conn)
}
