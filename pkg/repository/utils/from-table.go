package utils

import (
	"fmt"
	"strings"
)

func AddPrimaryTableToBuilder(builder *strings.Builder, tableName string) {
	fmt.Fprintf(builder, "FROM %s", tableName)
}
