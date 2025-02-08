package middleware

import (
	"github.com/google/uuid"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/requestid"
)

func RequestId() fiber.Handler {
	return requestid.New(requestid.Config{
		Header: "X-Custom-Header",
		Generator: func() string {
			newRequestID := uuid.NewString()
			return newRequestID
		},
	})
}
