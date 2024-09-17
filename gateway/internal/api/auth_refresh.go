package api

import (
	"server/infrastructure"
	"server/internal/dto"
	"server/pb"
	"server/util"

	"github.com/gin-gonic/gin"
)

// @Summary     Refresh
// @Tags       	Authentication
// @Accept		json
// @Produce		json
// @Param       body	body  dto.AuthRefresh	true  "body"
// @Success      200  {object}  util.SuccessResponse
// @Failure      400  {object}  util.ErrorResponse
// @Router		/auth/refresh/ [post]
func (m *ServiceServer) AuthRefresh(c *gin.Context) {
	ctx := c.Request.Context()
	payload := new(dto.AuthRefresh)

	if err := c.Bind(payload); err != nil {
		util.NewResponse(c).Error(err.Error(), "", 400)
		return
	}

	errors, errT := util.NewValidation().ValidateStruct(*payload)
	if errT != nil {
		util.NewResponse(c).Error(errors, infrastructure.Localize("FAILED_VALIDATION"), 400)
		return
	}
	token, err := m.authRpcServer.Refresh(ctx, &pb.AuthRefreshRequest{
		Token: payload.Token,
	})
	if err != nil {
		util.NewResponse(c).GrpcError(err, "")
		return
	}

	data := map[string]any{
		"access_token":    token.AccessToken,
		"access_expired":  token.AccessExpired,
		"refresh_token":   token.RefreshToken,
		"refresh_expired": token.RefreshExpired,
	}
	util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_READ"))
}
