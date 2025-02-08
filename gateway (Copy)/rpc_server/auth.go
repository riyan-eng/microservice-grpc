package rpcserver

import (
	"fmt"
	"log"
	"server/env"
	"server/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func AuthService() pb.AuthServiceClient {
	port := env.NewEnv().SERVICE_AUTH_PORT
	fmt.Println("grpc auth target:", port)
	conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return pb.NewAuthServiceClient(conn)
}
