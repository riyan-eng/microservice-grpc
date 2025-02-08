package api

import (
	"server/infrastructure"
	"server/pb"
	"server/util"

	"github.com/gofiber/fiber/v3"
)

// @Summary     Me
// @Tags       	Authentication
// @Accept		json
// @Produce		json
// @Success      200  {object}  util.SuccessResponse
// @Failure      400  {object}  util.ErrorResponse
// @Router		/auth/me [get]
// @Security 	BearerAuth
func (m *ServiceServer) AuthMe(c fiber.Ctx) error {
	ctx := c.Context()
	data, err := m.authRpcServer.Me(ctx, &pb.AuthMeRequest{})
	if err != nil {
		return util.NewResponse(c).GrpcError(err, "")
	}

	return util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_READ"))
}
