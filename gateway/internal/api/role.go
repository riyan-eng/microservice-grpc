package api

import (
	"server/infrastructure"
	"server/pb"
	"server/util"

	"github.com/gin-gonic/gin"
)

// @Summary      List
// @Tags       	 Role
// @Produce      json
// @Param        search		query   string	false  "search"
// @Param        page		query   int		false  "page"
// @Param        per_page	query   int		false  "per_page"
// @Router       /role [get]
// @Success      200  {object}  util.SuccessResponse
// @Failure      400  {object}  util.ErrorResponse
func (m *ServiceServer) RoleAccessList(c *gin.Context) {
	ctx := c.Request.Context()

	// queryParam := new(dto.PaginationReq).Init()
	// if err := c.BindQuery(&queryParam); err != nil {
	// 	util.NewResponse(c).Error(err, "", 400)
	// 	return
	// }
	// pageMeta := util.NewPagination().GetPageMeta(&queryParam.Page, &queryParam.Limit)

	res, err := m.permissionRpcServer.RoleAccessList(ctx, &pb.PermissionRoleAccessRequest{})
	if err != nil {
		util.NewResponse(c).GrpcError(err, "")
		return
	}

	util.NewResponse(c).Success(res.Data, nil, infrastructure.Localize("OK_READ"))
}
