package api

import (
	"server/infrastructure"
	"server/internal/dto"
	"server/pb"
	"server/util"

	"github.com/gofiber/fiber/v3"
)

// @Summary      List
// @Tags       	 Example
// @Produce      json
// @Param        search		query   string	false  "search"
// @Param        page		query   int		false  "page"
// @Param        per_page	query   int		false  "per_page"
// @Router       /example [get]
// @Success      200  {object}  util.SuccessResponse
// @Failure      400  {object}  util.ErrorResponse
func (m *ServiceServer) ExampleList(c fiber.Ctx) error {
	ctx := c.Context()

	queryParam := new(dto.PaginationReq).Init()
	if err := c.Bind().Query(&queryParam); err != nil {
		return util.NewResponse(c).Error(err, "", 400)
	}
	pageMeta := util.NewPagination().GetPageMeta(&queryParam.Page, &queryParam.Limit)

	res, err := m.exampleRpcServer.List(ctx, &pb.TaskListRequest{
		Search: queryParam.Search,
		Limit:  int32(*pageMeta.Limit),
		Offset: int32(*pageMeta.Offset),
	})

	if err != nil {
		return util.NewResponse(c).GrpcError(err, "")
	}

	meta := util.PaginationMeta{
		Page:       pageMeta.Page,
		Limit:      pageMeta.Limit,
		CountRows:  &res.TotalRows,
		CountPages: util.NewPagination().GetCountPages(&res.TotalRows, pageMeta.Limit),
	}
	return util.NewResponse(c).Success(res.Data, meta, infrastructure.Localize("OK_READ"))
}
