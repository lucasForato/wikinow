package utils

import (
	"html/template"
	"strings"
)

// This function removes the class attribute from an HTML tag.
//
// Example: `<strong class="bold">bold</strong>` -> `<strong>bold</strong>`
func RemoveClass(s template.HTML) string {
	str := string(s)
	var result strings.Builder
	for {
		start := strings.Index(str, ` class="`)
		if start == -1 {
			result.WriteString(str)
			break
		}
		result.WriteString(str[:start])
		str = str[start:]
		end := strings.Index(str, `"`)
		str = str[end+1:]
		end = strings.Index(str, `"`)
		if end == -1 {
			result.WriteString(str)
			break
		}
		str = str[end+1:]
	}
	return result.String()
}
