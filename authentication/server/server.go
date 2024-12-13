package server

import (
	"log"
	"net"
	"runtime"
	"server/config"
	"server/env"
	"server/infrastructure"
	"server/interceptor"
	"server/internal/controller"
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

func Start() { // New Start function to encapsulate the main logic
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

	// initiate cache
	cache, err := infrastructure.ConnRedis()
	if err != nil {
		log.Fatalf("Cache connection error: %v", err)
	}

	// initiate permission
	enforcer, err := config.NewEnforcer(sqlDB)
	if err != nil {
		log.Fatalf("Enforce error: %v", err)
	}

	myConfig := config.Config{
		SqlDB:      sqlDB,
		SqlXDB:     sqlxDB,
		Cache:      cache,
		Permission: enforcer,
	}

	token := util.NewToken(cache)
	security := util.NewSecurity()
	errorUtil := util.NewError()

	myUtil := util.Util{
		Token:    token,
		Security: security,
		Error:    errorUtil,
	}

	dao := repository.NewDAO(&myConfig)
	authService := service.NewAuthService(
		&myUtil,
		&dao,
		enforcer,
	)

	service := controller.NewService(
		&myUtil,
		dao,
		authService,
	)

	newInterceptor := interceptor.NewUnaryInterceptor(service)

	srv := grpc.NewServer(grpc.ChainUnaryInterceptor(
		newInterceptor.Unary(),
	))

	pb.RegisterAuthServiceServer(srv, service)
	pb.RegisterPermissionServiceServer(srv, service)

	log.Println("Starting RPC server:", env.NewEnv().SERVER_HOST+":"+env.NewEnv().SERVER_PORT)
	l, err := net.Listen("tcp", env.NewEnv().SERVER_HOST+":"+env.NewEnv().SERVER_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", env.NewEnv().SERVER_HOST+":"+env.NewEnv().SERVER_PORT, err)
	}
	log.Fatal(srv.Serve(l))
}
