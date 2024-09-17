package api

import (
	"server/infrastructure"
	"server/internal/dto"
	"server/pb"
	"server/util"

	"github.com/gin-gonic/gin"
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
func (m *ServiceServer) ExampleList(c *gin.Context) {
	ctx := c.Request.Context()

	queryParam := new(dto.PaginationReq).Init()
	if err := c.BindQuery(&queryParam); err != nil {
		util.NewResponse(c).Error(err, "", 400)
		return
	}
	pageMeta := util.NewPagination().GetPageMeta(&queryParam.Page, &queryParam.Limit)

	res, err := m.exampleRpcServer.List(ctx, &pb.TaskListRequest{
		Search: queryParam.Search,
		Limit:  int32(*pageMeta.Limit),
		Offset: int32(*pageMeta.Offset),
	})

	if err != nil {
		util.NewResponse(c).GrpcError(err, "")
		return
	}

	meta := util.PaginationMeta{
		Page:       pageMeta.Page,
		Limit:      pageMeta.Limit,
		CountRows:  &res.TotalRows,
		CountPages: util.NewPagination().GetCountPages(&res.TotalRows, pageMeta.Limit),
	}
	util.NewResponse(c).Success(res.Data, meta, infrastructure.Localize("OK_READ"))
}
