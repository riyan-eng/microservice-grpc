gen:
	@protoc --go_out . --go-grpc_out . proto/*.proto
dev:
	go run main.go
deploy:
	docker-compose down; docker-compose build; docker-compose up -d