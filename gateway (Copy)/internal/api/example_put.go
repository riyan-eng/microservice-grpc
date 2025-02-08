package api

import (
	"server/infrastructure"
	"server/internal/dto"
	"server/pb"
	"server/util"

	"github.com/gofiber/fiber/v3"
)

// @Summary     Put
// @Tags       	Example
// @Accept		json
// @Produce		json
// @Param       id	path	string	true	"id"
// @Param       body	body  dto.ExamplePut	true  "body"
// @Router      /example/{id} [put]
// @Security 	BearerAuth
// @Success      200  {object}  util.SuccessResponse
// @Failure      400  {object}  util.ErrorResponse
func (m *ServiceServer) ExamplePut(c fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	payload := new(dto.ExamplePut)

	if err := c.Bind().JSON(payload); err != nil {
		return util.NewResponse(c).Error(err.Error(), "", 400)
	}

	errors, err := util.NewValidation().ValidateStruct(*payload)
	if err != nil {
		return util.NewResponse(c).Error(errors, infrastructure.Localize("FAILED_VALIDATION"), 400)
	}

	_, err = m.exampleRpcServer.Put(ctx, &pb.TaskPutRequest{
		Id:     id,
		Name:   payload.Name,
		Detail: payload.Detail,
	})
	if err != nil {
		return util.NewResponse(c).GrpcError(err, "")
	}

	data := map[string]any{
		"id": id,
	}
	return util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_READ"))
}
