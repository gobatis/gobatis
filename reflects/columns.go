package reflects

import "strings"

func TrimColumns(s string) string {
	return strings.TrimSuffix(s, ",")
}

func ExtractTag(s string) string {
	idx := strings.Index(s, ";")
	if idx == -1 {
		return s
	}
	return s[:idx]
}
