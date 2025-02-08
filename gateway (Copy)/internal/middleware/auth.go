package middleware

import (
	"server/util"
	"strings"

	"github.com/gofiber/fiber/v3"
	"google.golang.org/grpc/metadata"
)

func AuthBearer() fiber.Handler {
	return func(c fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		bearerString := string(authHeader)
		if bearerString == "" {
			return util.NewResponse(c).Error("authorization header is required", "", 401)
		}
		token, found := strings.CutPrefix(bearerString, "Bearer ")
		if !found {
			return util.NewResponse(c).Error("undefined token", "", 401)
		}
		md := metadata.Pairs("authorization", token)
		ctx := metadata.NewOutgoingContext(c.Context(), md)

		c.SetUserContext(ctx)
		return c.Next()
	}
}
