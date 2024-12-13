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

// type taskServer struct {
// 	rpc.UnimplementedTaskServiceServer
// }

// func (taskServer) Create(ctx context.Context, req *rpc.TaskCreateRequest) (*rpc.TaskCreateResponse, error) {
// 	fmt.Println("before")
// 	fmt.Println(req)
// 	fmt.Println("after")

// 	return &rpc.TaskCreateResponse{}, nil
// }

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

	myConfig := config.Config{
		SqlDB:  sqlDB,
		SqlXDB: sqlxDB,
	}

	errorUtil := util.NewError()

	myUtil := util.Util{
		Error: errorUtil,
	}

	dao := repository.NewDAO(&myConfig)
	exampleService := service.NewExampleService(&dao)

	service := handler.NewService(&myUtil, dao, exampleService)

	srv := grpc.NewServer(grpc.ChainUnaryInterceptor(
		interceptor.Unary(),
	))
	pb.RegisterTaskServiceServer(srv, service)

	log.Println("Starting RPC server:", env.NewEnv().SERVER_HOST+":"+env.NewEnv().SERVER_PORT)
	l, err := net.Listen("tcp", env.NewEnv().SERVER_HOST+":"+env.NewEnv().SERVER_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", env.NewEnv().SERVER_HOST+":"+env.NewEnv().SERVER_PORT, err)
	}
	log.Fatal(srv.Serve(l))
}
