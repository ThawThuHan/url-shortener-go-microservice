package main

import (
	"context"
	"log"
	pb "redirection-service/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:5882", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error to connect %v", err)
	}

	defer conn.Close()

	client := pb.NewRedirectionServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.GetOriginURL(ctx, &pb.GetOriginURLRequest{
		ShortCode: "LjBtHO",
		IpAddress: "192.168.10.1",
		City:      "BangKok",
		Location:  "13.7540,100.5014",
	})
	if err != nil {
		log.Println(err.Error())
	}
	log.Printf("response: %v", resp.OriginUrl)

	accessLogResp, err := client.GetAccessLog(ctx, &pb.GetAccessLogRequest{
		ShortCode: "LjBtHO",
	})
	if err != nil {
		log.Println(err.Error())
	}
	log.Printf("access log resp: %v", accessLogResp.Logs)
}
