package util

import "math"

type pageMeta struct {
	Page   *int32
	Limit  *int32
	Offset *int32
}

type paginationStruct struct{}

func NewPagination() *paginationStruct {
	return &paginationStruct{}
}

func (p *paginationStruct) GetCountPages(totalRows, limit *int32) *int32 {
	countPages := int32(math.Ceil(float64(*totalRows) / float64(*limit)))
	return &countPages
}

func (p *paginationStruct) GetPageMeta(page, limit *int32) *pageMeta {
	if *page < 1 {
		*page = 1
	}
	if *limit < 1 {
		*limit = 1
	}

	offset := *limit * (*page - 1)
	return &pageMeta{
		Page:   page,
		Limit:  limit,
		Offset: &offset,
	}
}
