package api

import (
	"server/infrastructure"
	"server/internal/dto"
	"server/pb"
	"server/util"

	"github.com/gin-gonic/gin"
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
func (m *ServiceServer) ExampleCreate(c *gin.Context) {
	ctx := c.Request.Context()
	payload := new(dto.ExampleCreate)

	if err := c.Bind(payload); err != nil {
		util.NewResponse(c).Error(err.Error(), "", 400)
		return
	}

	errors, errT := util.NewValidation().ValidateStruct(*payload)
	if errT != nil {
		util.NewResponse(c).Error(errors, infrastructure.Localize("FAILED_VALIDATION"), 400)
		return
	}

	id := uuid.NewString()

	_, err := m.exampleRpcServer.Create(ctx, &pb.TaskCreateRequest{
		Id:     id,
		Name:   payload.Name,
		Detail: payload.Detail,
	})

	if err != nil {
		// util.NewResponse(c).Error(err.Error(), "")
		util.NewResponse(c).GrpcError(err, "")
		return
	}

	data := map[string]any{
		"id": id,
	}
	util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_CREATE"), 201)
}
