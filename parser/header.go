package parser

import (
	"regexp"
)

type Header struct {
	Container

	Level int8
}

func ParseHeader(content string, parent Node) []Node {
	var nodes []Node

	boldPattern := regexp.MustCompile(`(\*\*(.*?)\*\*)`)
	segments := boldPattern.FindAllStringSubmatchIndex(content, -1)

	lastIndex := 0
	for _, match := range segments {
		if match[0] > lastIndex {
			paragraph := new(Paragraph)
			paragraph.Parent = parent
			paragraph.Content = content[lastIndex:match[0]]
			nodes = append(nodes, paragraph)
		}

		bold := new(Bold)
		bold.Parent = parent
		bold.Content = content[match[2]:match[3]]
		nodes = append(nodes, bold)

		lastIndex = match[1]
	}

	if lastIndex < len(content) {
		paragraph := new(Paragraph)
		paragraph.Parent = parent
		paragraph.Content = content[lastIndex:]
		nodes = append(nodes, paragraph)
	}

	return nodes
}
