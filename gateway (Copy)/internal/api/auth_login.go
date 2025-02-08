package api

import (
	"server/infrastructure"
	"server/internal/dto"
	"server/pb"
	"server/util"

	"github.com/gofiber/fiber/v3"
)

// @Summary     Login
// @Tags       	Authentication
// @Accept		json
// @Produce		json
// @Param       body	body  dto.AuthLogin	true  "body"
// @Success      200  {object}  util.SuccessResponse
// @Failure      400  {object}  util.ErrorResponse
// @Router		/auth/login [post]
func (m *ServiceServer) AuthLogin(c fiber.Ctx) error {
	ctx := c.Context()
	payload := new(dto.AuthLogin)

	if err := c.Bind().JSON(payload); err != nil {
		return util.NewResponse(c).Error(err.Error(), "", 400)
	}

	errors, err := util.NewValidation().ValidateStruct(*payload)
	if err != nil {
		return util.NewResponse(c).Error(errors, infrastructure.Localize("FAILED_VALIDATION"), 400)
	}

	data, err := m.authRpcServer.Login(ctx, &pb.AuthLoginRequest{
		Username: payload.Username,
		Password: payload.Password,
	})

	if err != nil {
		return util.NewResponse(c).GrpcError(err, "")
	}

	return util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_READ"))
}
