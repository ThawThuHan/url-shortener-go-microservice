package grpc

import (
	"api-gateway/config"
	pb "api-gateway/proto/redirection-service"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewRedirectionClient(cfg *config.Config) pb.RedirectionServiceClient {
	conn, err := grpc.NewClient(cfg.RedirectionService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error to connect %v", err)
	}
	return pb.NewRedirectionServiceClient(conn)
}
