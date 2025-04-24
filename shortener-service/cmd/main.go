package main

import (
	"log"
	"net"
	"shortener-service/configs"
	"shortener-service/database"
	svr "shortener-service/grpc"
	"shortener-service/internal/repository"
	"shortener-service/internal/service"
	pb "shortener-service/proto"

	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
)

func main() {
	cfg := configs.Cfg

	lis, err := net.Listen("tcp", cfg.Host+":"+cfg.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	db := database.NewDatabase(cfg)
	if err := db.Migrate(); err != nil {
		log.Fatalf("failed to database migrate: %v", err)
	}
	validator := validator.New()
	grpcServer := grpc.NewServer()

	urlRepo := repository.NewShortenerRepo(db.DB)
	urlService := service.NewShortenerService(urlRepo)
	shortenerServer := svr.NewShortenerServer(urlService, validator)
	pb.RegisterShortenerServiceServer(grpcServer, shortenerServer)

	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
