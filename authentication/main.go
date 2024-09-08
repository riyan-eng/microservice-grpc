package main

import (
	"log"
	"net"
	"runtime"
	"server/config"
	"server/env"
	"server/infrastructure"
	"server/internal/handler"
	"server/internal/repository"
	"server/internal/service"
	"server/pb"

	"google.golang.org/grpc"
)

func init() {
	numCPU := runtime.NumCPU()
	if numCPU <= 1 {
		runtime.GOMAXPROCS(1)
	} else {
		runtime.GOMAXPROCS(numCPU - 1)
	}
	env.LoadEnvironmentFile()
	env.NewEnv()

	infrastructure.ConnectSqlDB()
	infrastructure.ConnectSqlxDB()
	infrastructure.ConnRedis()
}

func main() {
	srv := grpc.NewServer()

	dao := repository.NewDAO(infrastructure.SqlDB, infrastructure.SqlxDB, infrastructure.Redis, config.NewEnforcer())
	exampleService := service.NewExampleService(&dao)
	authService := service.NewAuthService(&dao)

	service := handler.NewService(dao,
		exampleService,
		authService,
	)

	pb.RegisterTaskServiceServer(srv, service)
	pb.RegisterAuthServiceServer(srv, service)

	log.Println("Starting RPC server:", env.NewEnv().SERVER_HOST+":"+env.NewEnv().SERVER_PORT)
	l, err := net.Listen("tcp", env.NewEnv().SERVER_HOST+":"+env.NewEnv().SERVER_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", env.NewEnv().SERVER_HOST+":"+env.NewEnv().SERVER_PORT, err)
	}
	log.Fatal(srv.Serve(l))
}
