package utils

import "strings"

func AddLimit1(builder *strings.Builder) {
	builder.WriteString(" LIMIT 1")
}
