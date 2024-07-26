package parser

import (
	"regexp"
)

type Italic struct {
	Container
}

func NewItalic(raw string, content string, start int, end int) Node {
	italic := new(Italic)
	italic.Type = "Italic"
	italic.Raw = raw
	italic.Start = start
	italic.End = end
	children := Parse(content)
	if children != nil {
		italic.SetChildren(*children)
	}
	return italic
}

func ParseItalic(in string) *[]Node {
	regex := regexp.MustCompile(`(\*(.*?)\*)`)
	segments := regex.FindAllStringSubmatchIndex(in, -1)
	if len(segments) == 0 {
		return nil
	}

	result := []Node{}
	for _, match := range segments {
		italic := NewItalic(
			in[match[0]:match[1]],
			in[match[4]:match[5]],
			match[0],
			match[1],
		)
		result = append(result, italic)
	}
	return &result
}
