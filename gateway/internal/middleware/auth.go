package middleware

import (
	"server/util"
	"strings"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
)

func AuthBearer() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		bearerString := string(authHeader)
		if bearerString == "" {
			util.NewResponse(c).Error("authorization header is required", "", 400)
			return
		}
		token, found := strings.CutPrefix(bearerString, "Bearer ")
		if !found {
			util.NewResponse(c).Error("undefined token", "", 400)
			return

		}
		md := metadata.Pairs("authorization", token)
		ctx := metadata.NewOutgoingContext(c, md)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
