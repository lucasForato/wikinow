package parser

import (
	"regexp"
)

type Bold struct {
	Container
}

func NewBold(raw string, content string, start int, end int) Node {
	bold := new(Bold)
	bold.Type = "Bold"
	bold.Raw = raw
	bold.Start = start
	bold.End = end
	children := Parse(content)
	if children != nil {
		bold.SetChildren(*children)
	}
	return bold
}

func ParseBold(in string) *[]Node {
	regex := regexp.MustCompile(`(\*\*(.*?)\*\*)`)
	segments := regex.FindAllStringSubmatchIndex(in, -1)
	if len(segments) == 0 {
		return nil
	}

	result := []Node{}
	for _, match := range segments {
		bold := NewBold(
			in[match[0]:match[1]],
			in[match[4]:match[5]],
			match[0],
			match[1],
		)
		result = append(result, bold)
	}
	return &result
}
