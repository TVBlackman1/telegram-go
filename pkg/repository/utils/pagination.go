package utils

import "math"

type QueryPagination struct {
	Page int
	Size int
}

func (p *QueryPagination) DB() DBPagination {
	offset := (p.Page - 1) * p.Size
	return DBPagination{
		Offset: offset,
		Limit:  p.Size,
	}
}

func (p *QueryPagination) CheckValues(maxSize int) {
	p.checkMaxPageSize(maxSize)
	p.checkPageNumber()
}

func (p *QueryPagination) checkPageNumber() {
	if p.Page == 0 {
		p.Page = 1
	}
}

func (p *QueryPagination) checkMaxPageSize(maxSize int) {
	if p.Size == 0 || p.Size > maxSize {
		p.Size = maxSize
	}
}

type Pagination struct {
	Page  int
	Pages int
	Limit int
}

type DBPagination struct {
	Offset int
	Limit  int
}

func NewPagination(queryPagination QueryPagination, count int) *Pagination {
	pagination := new(Pagination)
	pagination.Pages = getPages(queryPagination.Size, count)
	pagination.Page = queryPagination.Page
	pagination.Limit = queryPagination.Size
	return pagination
}

func getPages(pageLimit int, count int) int {
	tmp := float64(count) / float64(pageLimit)
	return int(math.Ceil(tmp))
}
