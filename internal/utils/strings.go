package utils

import (
	"strings"
)

// Normalize removes leading and trailing whitespace and converts the string to lowercase.
func Normalize(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}
