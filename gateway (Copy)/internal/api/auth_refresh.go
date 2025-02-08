package api

import (
	"server/infrastructure"
	"server/internal/dto"
	"server/pb"
	"server/util"

	"github.com/gofiber/fiber/v3"
)

// @Summary     Refresh
// @Tags       	Authentication
// @Accept		json
// @Produce		json
// @Param       body	body  dto.AuthRefresh	true  "body"
// @Success      200  {object}  util.SuccessResponse
// @Failure      400  {object}  util.ErrorResponse
// @Router		/auth/refresh [post]
func (m *ServiceServer) AuthRefresh(c fiber.Ctx) error {
	ctx := c.Context()
	payload := new(dto.AuthRefresh)

	if err := c.Bind().JSON(payload); err != nil {
		return util.NewResponse(c).Error(err.Error(), "", 400)
	}

	errors, err := util.NewValidation().ValidateStruct(*payload)
	if err != nil {
		return util.NewResponse(c).Error(errors, infrastructure.Localize("FAILED_VALIDATION"), 400)
	}
	token, err := m.authRpcServer.Refresh(ctx, &pb.AuthRefreshRequest{
		Token: payload.Token,
	})
	if err != nil {
		return util.NewResponse(c).GrpcError(err, "")
	}

	data := map[string]any{
		"access_token":    token.AccessToken,
		"access_expired":  token.AccessExpired,
		"refresh_token":   token.RefreshToken,
		"refresh_expired": token.RefreshExpired,
	}
	return util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_READ"))
}
