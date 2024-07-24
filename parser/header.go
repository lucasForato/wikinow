package parser

type Header struct {
	Container

	Level int8
}

func ParseHeader(content string, parent Node) []Node {
	var nodes []Node

	ParseBold(content)

	return nodes
}
