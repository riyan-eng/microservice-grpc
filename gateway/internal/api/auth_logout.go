package api

import (
	"server/infrastructure"
	"server/pb"
	"server/util"

	"github.com/gin-gonic/gin"
)

// @Summary     Logout
// @Tags       	Authentication
// @Accept		json
// @Produce		json
// @Success      200  {object}  util.SuccessResponse
// @Failure      400  {object}  util.ErrorResponse
// @Router		/auth/logout/ [delete]
// @Security 	BearerAuth
func (m *ServiceServer) AuthLogout(c *gin.Context) {
	ctx := c.Request.Context()

	_, err := m.authRpcServer.Logout(ctx, &pb.AuthLogoutRequest{})
	if err != nil {
		util.NewResponse(c).GrpcError(err, "")
		return
	}

	util.NewResponse(c).Success(nil, nil, infrastructure.Localize("OK_LOGOUT"))
}
