package parser

import (
	"regexp"
)

type Text struct {
	Leaf
}

func NewText(raw string, content string, start int, end int) Node {
	text := new(Text)
	text.Type = "Text"
	text.Raw = raw
	text.Start = start
	text.End = end
	return text
}

func ParseText(in string) *[]Node {
	regex := regexp.MustCompile(`^[^*_#>\[\]~]+$`)
	segments := regex.FindAllStringSubmatchIndex(in, -1)
	if len(segments) == 0 {
		return nil
	}

	result := []Node{}
	for _, match := range segments {
		bold := NewText(
			in[match[0]:match[1]],
			in[match[0]:match[1]],
			match[0],
			match[1],
		)
		result = append(result, bold)
	}
	return &result
}
