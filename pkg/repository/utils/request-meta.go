package utils

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

type RequestWithPagination struct {
	Db         *sqlx.DB
	LogicPart  string
	Selected   string
	Pagination QueryPagination
}

func (r *RequestWithPagination) SelectIn(responseVar interface{}) (Pagination, error) {
	var count int
	queryBuilder := NewQueryBuilder(r.LogicPart, r.Pagination.DB())
	countRequest := queryBuilder.WithSelected(SELECTED_COUNT, NO_PAGINATION)
	request := queryBuilder.WithSelected(r.Selected, WITH_PAGINATION)
	if err := r.Db.Select(responseVar, request); err != nil {
		fmt.Fprintf(os.Stderr, "Bad request: %s", err.Error())
		return Pagination{}, err
	}
	r.Db.Get(&count, countRequest)
	pagination := NewPagination(r.Pagination, count)
	return *pagination, nil
}
