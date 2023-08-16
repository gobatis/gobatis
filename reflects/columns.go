package reflects

import "strings"

func SplitColumns(s string) []string {
	columns := strings.Split(s, ",")
	for i := range columns {
		columns[i] = strings.TrimSpace(columns[i])
	}
	return columns
}

func ExtractTag(s string) string {
	idx := strings.Index(s, ";")
	if idx == -1 {
		return s
	}
	return s[:idx]
}
