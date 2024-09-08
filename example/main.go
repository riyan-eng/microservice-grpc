package main

import (
	"log"
	"net"
	"runtime"
	"server/env"
	"server/infrastructure"
	"server/interceptor"
	"server/internal/handler"
	"server/internal/repository"
	"server/internal/service"
	"server/pb"

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

	infrastructure.ConnectSqlDB()
	infrastructure.ConnectSqlxDB()
}

func main() {
	srv := grpc.NewServer(grpc.ChainUnaryInterceptor(
		// selector.UnaryServerInterceptor(
		// 	auth.UnaryServerInterceptor(interceptor.Authenticator),
		// 	selector.MatchFunc(interceptor.AuthMatcher),
		// ),
		interceptor.Unary(),
	))

	dao := repository.NewDAO(infrastructure.SqlDB, infrastructure.SqlxDB)
	exampleService := service.NewExampleService(&dao)

	service := handler.NewService(dao, exampleService)
	pb.RegisterTaskServiceServer(srv, service)

	log.Println("Starting RPC server:", env.NewEnv().SERVER_HOST+":"+env.NewEnv().SERVER_PORT)
	l, err := net.Listen("tcp", env.NewEnv().SERVER_HOST+":"+env.NewEnv().SERVER_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", env.NewEnv().SERVER_HOST+":"+env.NewEnv().SERVER_PORT, err)
	}
	log.Fatal(srv.Serve(l))
}
