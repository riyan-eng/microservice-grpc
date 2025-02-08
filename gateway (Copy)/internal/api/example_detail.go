package api

import (
	"server/infrastructure"
	"server/pb"
	"server/util"

	"github.com/gofiber/fiber/v3"
)

// @Summary     Detail
// @Tags       	Example
// @Accept		json
// @Produce		json
// @Param       id	path	string	true	"id"
// @Router      /example/{id} [get]
// @Security 	BearerAuth
// @Success      200  {object}  util.SuccessResponse
// @Failure      400  {object}  util.ErrorResponse
func (m *ServiceServer) ExampleDetail(c fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	res, err := m.exampleRpcServer.Detail(ctx, &pb.TaskDetailRequest{
		Id: id,
	})
	if err != nil {
		return util.NewResponse(c).GrpcError(err, "")
	}

	return util.NewResponse(c).Success(res.Data, nil, infrastructure.Localize("OK_READ"))
}
