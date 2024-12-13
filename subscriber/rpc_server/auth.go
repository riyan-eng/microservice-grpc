package rpcserver

import (
	"log"
	"server/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func AuthService() pb.AuthServiceClient {
	port := "127.0.0.1:3002"
	conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return pb.NewAuthServiceClient(conn)
}
