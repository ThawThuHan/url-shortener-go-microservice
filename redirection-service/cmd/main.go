package main

import (
	"log"
	"net"
	"redirection-service/configs"
	"redirection-service/database"
	svr "redirection-service/grpc"
	"redirection-service/internal/repository"
	"redirection-service/internal/service"
	pb "redirection-service/proto"

	"google.golang.org/grpc"
)

func main() {
	cfg := configs.Env

	lis, err := net.Listen("tcp", cfg.HOST+":"+cfg.PORT)
	if err != nil {
		log.Fatalf("Failed to listen on %v, %v", cfg.HOST+":"+cfg.PORT, err)
	}

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}
	database.Migrate(db)

	redisClient, err := database.Redis(*cfg)
	if err != nil {
		log.Fatalf("Failed to connect Redis: %v", err)
	}

	grpcSvr := grpc.NewServer()

	urlRepo := repository.NewURLRepository(db, redisClient)
	accessLogRepo := repository.NewAccessLogRepo(db)
	urlService := service.NewURLService(urlRepo)
	accessLogService := service.NewAccessLogService(accessLogRepo)

	redirectionServer := svr.NewRedirectionServer(urlService, accessLogService)

	pb.RegisterRedirectionServiceServer(grpcSvr, redirectionServer)

	log.Printf("server listening at %v", lis.Addr())
	if err := grpcSvr.Serve(lis); err != nil {
		log.Fatalf("Failed to serve grpc-server, %v", err)
	}
}
