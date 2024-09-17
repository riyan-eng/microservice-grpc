package rpcserver

import (
	"log"
	"server/env"
	"server/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ExampleService() pb.TaskServiceClient {
	port := env.NewEnv().SERVICE_EXAMPLE_PORT
	conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return pb.NewTaskServiceClient(conn)
}
