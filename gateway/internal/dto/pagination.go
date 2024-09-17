package dto

type PaginationReq struct {
	Page   int32    `form:"page"`
	Limit  int32    `form:"per_page"`
	Search string `form:"search"`
	Order  string `form:"order"`
}

func (p PaginationReq) Init() PaginationReq {
	p.Page = 1
	p.Limit = 10
	p.Order = "desc"
	return p
}
