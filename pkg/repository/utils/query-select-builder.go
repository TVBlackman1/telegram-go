package utils

import (
	"fmt"
	"strings"
)

type ruleResponsePagination int

const (
	WITH_PAGINATION ruleResponsePagination = iota
	NO_PAGINATION
)

type QueryBuilder struct {
	logicPart  string
	pagination DBPagination
}

func NewQueryBuilder(logicPart string, pagination DBPagination) *QueryBuilder {
	return &QueryBuilder{logicPart, pagination}
}

func (q *QueryBuilder) WithSelected(selected string, paginationRule ruleResponsePagination) string {
	var request strings.Builder
	request.WriteString("SELECT ")
	request.WriteString(selected)
	request.WriteRune(' ')
	request.WriteString(q.logicPart)
	if paginationRule == WITH_PAGINATION {
		fmt.Fprintf(&request, " OFFSET %d LIMIT %d", q.pagination.Offset, q.pagination.Limit)
	}
	return request.String()
}

func (q *QueryBuilder) UpdateLogicPart(logicPart string) {
	q.logicPart = logicPart
}
