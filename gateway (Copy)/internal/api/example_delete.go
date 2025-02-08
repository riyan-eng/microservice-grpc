package api

import (
	"server/infrastructure"
	"server/pb"
	"server/util"

	"github.com/gofiber/fiber/v3"
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
func (m *ServiceServer) ExampleDelete(c fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	_, err := m.exampleRpcServer.Delete(ctx, &pb.TaskDeleteRequest{
		Id: id,
	})
	if err != nil {
		return util.NewResponse(c).GrpcError(err, "")
	}

	data := map[string]any{
		"id": id,
	}
	return util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_DELETE"))
}
