package utils

import "strings"

func IsHeader(line string) bool {
	return strings.HasPrefix(line, "# ")
}

func GetHeaderContent(line string) (string, int8) {
	if strings.HasPrefix(line, "###### ") {
		return strings.TrimPrefix(line, "###### "), 6
	}
	if strings.HasPrefix(line, "##### ") {
		return strings.TrimPrefix(line, "##### "), 5
	}
	if strings.HasPrefix(line, "#### ") {
		return strings.TrimPrefix(line, "#### "), 4
	}
	if strings.HasPrefix(line, "### ") {
		return strings.TrimPrefix(line, "### "), 3
	}
	if strings.HasPrefix(line, "## ") {
		return strings.TrimPrefix(line, "## "), 2
	}
	if strings.HasPrefix(line, "# ") {
		return strings.TrimPrefix(line, "# "), 1
	}
	panic("Content is not a header")
}
