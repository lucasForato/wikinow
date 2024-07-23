package parser

import "fmt"

type Header struct {
	Container

	Level int8
}

func ParseHeader(content string, parent Node) []Node {
	var nodes []Node
  fmt.Println(content)

  FindBold(content)	

	return nodes
}
