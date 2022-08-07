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
