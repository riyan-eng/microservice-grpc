package api

import (
	"server/infrastructure"
	"server/pb"
	"server/util"

	"github.com/gin-gonic/gin"
)

// @Summary     Delete
// @Tags       	Example
// @Accept		json
// @Produce		json
// @Param       id	path	string	true	"id"
// @Success      200  {object}  util.SuccessResponse
// @Failure      400  {object}  util.ErrorResponse
// @Router      /example/{id} [delete]
// @Security 	BearerAuth
func (m *ServiceServer) ExampleDelete(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	_, err := m.exampleRpcServer.Delete(ctx, &pb.TaskDeleteRequest{
		Id: id,
	})
	if err != nil {
		util.NewResponse(c).GrpcError(err, "")
		return
	}

	data := map[string]any{
		"id": id,
	}
	util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_DELETE"))
}
