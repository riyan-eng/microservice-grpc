package api

import (
	"github.com/gofiber/fiber/v3"
)

func (m *ServiceServer) Index(c fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "Full ganteng",
	})
}
