package parser

import (
	"regexp"
)

type Italic struct {
	Container
}

func NewItalic(raw string, content string, start int, end int, ctx *Ctx) Node {
	italic := new(Italic)
	italic.Type = "Italic"
	italic.Raw = raw
	italic.Start = start
	italic.End = end
	children, _ := Parse(content, ctx)
	if children != nil {
		italic.SetChildren(*children)
	}
	return italic
}

func ParseItalic(in string, ctx *Ctx) *[]Node {
	regex := regexp.MustCompile(`\*(.+?)\*|_(.+?)_`)
	segments := regex.FindAllStringSubmatchIndex(in, -1)
	if len(segments) == 0 {
		return nil
	}

	result := []Node{}
	for _, match := range segments {
		raw := in[match[0]:match[1]]
		content := in[match[2]:match[3]]
		if content[0] == '*' || content[len(content)-1] == '*' {
			continue
		}

		italic := NewItalic(
			raw,
			content,
			match[0],
			match[1],
			ctx,
		)
		result = append(result, italic)
	}
	return &result
}
