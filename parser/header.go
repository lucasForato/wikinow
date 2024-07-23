package parser

import (
	"slices"
	"wikinow/utils"
)

type Header struct {
	Container

	Level int8
}

func ParseHeader(content string, parent Node) []Node {
	var nodes []Node

	substrings, matchIndices := utils.FindMatches(utils.BoldPattern, content)
	for i, sub := range substrings {
		if slices.Contains(matchIndices, i) {
			bold := NewBold(sub, parent)
			nodes = append(nodes, bold)
			continue
		}
		paragraph := NewParagraph(sub, parent)
		nodes = append(nodes, paragraph)
	}

	return nodes
}
