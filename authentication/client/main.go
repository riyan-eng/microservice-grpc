package main

import (
	"context"
	"fmt"
	"log"
	"server/pb"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func exampleService() pb.TaskServiceClient {
	port := ":3000"
	conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return pb.NewTaskServiceClient(conn)
}

func main() {

	app := fiber.New()

	example := exampleService()

	app.Get("/example", func(c *fiber.Ctx) error {
		mama, err := example.List(context.Background(), &pb.TaskListRequest{
			Search: c.Query("search"),
			Limit:  10,
			Offset: 0,
		})
		fmt.Println(err)
		fmt.Println(mama)

		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"data":    mama.Data,
			"message": "Berhasil menampilkan data.",
		})
		// return nil
	})

	app.Post("/example", func(c *fiber.Ctx) error {

		return nil
	})

	// mama, err := example.Create(context.Background(), &rpc.TaskCreateRequest{
	// 	Id:     "id",
	// 	Name:   "riyan",
	// 	Detail: "ganteng",
	// })

	// mama, err := example.List(context.Background(), &rpc.TaskListRequest{
	// 	Search: "",
	// 	Limit:  10,
	// 	Offset: 0,
	// })

	// fmt.Println(mama)

	// lala, _ := json.Marshal(mama.Data)
	// fmt.Println(string(lala))
	// fmt.Println(err)

	app.Listen(":3001")
}
