package parser

import (
	"regexp"
	"strings"
)

type Bold struct {
	Container

	Content string
}

func NewBold(content string) *Bold {
	bold := new(Bold)
  bold.Type = "Bold"
	bold.Content = strings.Trim(content, "**")
	return bold
}

func ParseBold(content string) []string {
	regex := regexp.MustCompile(`(\*\*(.*?)\*\*)`)
	segments := regex.FindAllStringSubmatch(content, -1)

	result := []string{}
	for _, match := range segments {
		text := match[len(match)-1]
		result = append(result, strings.Trim(text, "*"))
	}
	return result
}
