package parser

import (
	"fmt"
	"regexp"
)

type Bold struct {
	Container
}

func NewBold(raw string, content string, start int, end int, ctx *Ctx) Node {
  fmt.Println("NewBold")
	bold := new(Bold)
	bold.Type = "Bold"
	bold.Raw = raw
	bold.Start = start
	bold.End = end
	children, _ := Parse(content, ctx)
	if children != nil {
		bold.SetChildren(*children)
	}
	return bold
}

func ParseBold(in string, ctx *Ctx) *[]Node {
	regex := regexp.MustCompile(`\*\*(.+?)\*\*|__(.+?)__`)
	segments := regex.FindAllStringSubmatchIndex(in, -1)
	if len(segments) == 0 {
		return nil
	}

	result := []Node{}
	for _, match := range segments {
		bold := NewBold(
			in[match[0]:match[1]],
			in[match[2]:match[3]],
			match[0],
			match[1],
      ctx,
		)
		result = append(result, bold)
	}
	return &result
}
