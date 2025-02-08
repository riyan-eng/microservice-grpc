package api

import (
	"server/infrastructure"
	"server/internal/dto"
	"server/pb"
	"server/util"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

// @Summary     Create
// @Tags       	Example
// @Accept		json
// @Produce		json
// @Param       body	body  dto.ExampleCreate	true  "body"
// @Success      200  {object}  util.SuccessResponse
// @Failure      400  {object}  util.ErrorResponse
// @Router		/example [post]
// @Security 	BearerAuth
func (m *ServiceServer) ExampleCreate(c fiber.Ctx) error {
	ctx := c.Context()
	payload := new(dto.ExampleCreate)

	if err := c.Bind().JSON(payload); err != nil {
		return util.NewResponse(c).Error(err.Error(), "", 400)
	}

	errors, err := util.NewValidation().ValidateStruct(*payload)
	if err != nil {
		return util.NewResponse(c).Error(errors, infrastructure.Localize("FAILED_VALIDATION"), 400)
	}

	id := uuid.NewString()

	_, err = m.exampleRpcServer.Create(ctx, &pb.TaskCreateRequest{
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
	return util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_CREATE"), 201)
}
