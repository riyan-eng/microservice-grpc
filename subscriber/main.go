package main

import (
	"log"
	"net"
	"runtime"
	"server/config"
	"server/env"
	"server/infrastructure"
	"server/interceptor"
	"server/internal/handler"
	"server/internal/repository"
	"server/internal/service"
	"server/pb"
	"server/util"

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
}

func main() {
	// initiate db connection
	sqlDB, err := infrastructure.ConnectSqlDB()
	if err != nil {
		log.Fatalf("Database sql connection error: %v", err)
	}
	defer sqlDB.Close()

	sqlxDB, err := infrastructure.ConnectSqlxDB()
	if err != nil {
		log.Fatalf("Database sqlx connection error: %v", err)
	}
	defer sqlxDB.Close()

	mqConn, channel, err := infrastructure.ConnectRabbitMq()
	if err != nil {
		log.Fatalf("RabbitMq ch connection error: %v", err)
	}
	defer mqConn.Close()
	defer channel.Close()

	myConfig := config.Config{
		SqlDB:   sqlDB,
		SqlXDB:  sqlxDB,
		Channel: channel,
	}

	errorUtil := util.NewError()

	myUtil := util.Util{
		Error: errorUtil,
	}

	dao := repository.NewDAO(&myConfig)
	exampleService := service.NewExampleService(&dao)
	backgroundService := service.NewBackgroundService(&dao)

	service := handler.NewService(
		&myUtil,
		dao,
		exampleService,
		backgroundService,
	)

	srv := grpc.NewServer(grpc.ChainUnaryInterceptor(
		interceptor.Unary(),
	))
	pb.RegisterTaskServiceServer(srv, service)
	pb.RegisterBackgroundServiceServer(srv, service)

	log.Println("Starting RPC server:", env.NewEnv().SERVER_HOST+":"+env.NewEnv().SERVER_PORT)
	l, err := net.Listen("tcp", env.NewEnv().SERVER_HOST+":"+env.NewEnv().SERVER_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", env.NewEnv().SERVER_HOST+":"+env.NewEnv().SERVER_PORT, err)
	}
	log.Fatal(srv.Serve(l))
}
