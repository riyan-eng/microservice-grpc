package main

import (
	"consumer-trial/env"
	"consumer-trial/infrastructure"
	"consumer-trial/internal/rabbitmq"
	"log"
	"runtime"
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
	sqlxDB, err := infrastructure.ConnectSqlxDB()
	if err != nil {
		log.Fatalf("Failed to initialize SQLX DB: %v", err)
	}

	// Setup RabbitMQ connection
	conn, ch, err := rabbitmq.Connect()
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ: %v", err)
	}
	defer conn.Close()
	defer ch.Close()

	// Declare RabbitMQ topology (exchange, queue, bindings)
	err = rabbitmq.DeclareTopology(ch)
	if err != nil {
		log.Fatalf("Failed to declare RabbitMQ topology: %v", err)
	}

	// Start consuming messages
	log.Println("Waiting for messages...")
	rabbitmq.ConsumeAndProcessMessages(sqlxDB, ch)
}
