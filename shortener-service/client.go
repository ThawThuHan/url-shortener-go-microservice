package main

import (
	"context"
	"log"
	"time"

	pb "shortener-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:5881",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewShortenerServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create short URL
	resp, err := client.CreateShortURL(ctx, &pb.CreateShortURLRequest{
		OriginUrl: "http://example.com/test1",
		SessionId: "12345678",
	})
	if err != nil {
		log.Fatalf("CreateShortUrl failed: %v", err)
	}
	log.Printf("Short URL created: %s", resp.ShortCode)

	// Get original URL
	originalResp, err := client.GetOriginURLs(ctx, &pb.GetOriginURLsRequest{
		SessionId: "12345678",
	})
	if err != nil {
		log.Fatalf("GetOriginalUrl failed: %v", err)
	}
	log.Printf("Original URL: %s", originalResp.Urls)
}
