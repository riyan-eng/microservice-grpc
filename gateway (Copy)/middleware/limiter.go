package middleware

import (
	"context"
	"log"
	"server/config"

	"github.com/gofiber/fiber/v3"
)

func Limiter() fiber.Handler {
	return func(c fiber.Ctx) error {
		ctx := context.Background()
		_, _, _, ok, err := config.LimiterStore.Take(ctx, "key")
		if err != nil {
			log.Fatal(err)
		}
		if !ok {
			return c.Status(429).SendString("too many request")
		}

		return c.Next()
	}
}
