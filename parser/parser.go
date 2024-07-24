package parser

import (
	"strings"
	"wikinow/ast"
)

type (
	Node      = ast.Node
	Container = ast.Container
	Document  = ast.Document
	Leaf      = ast.Leaf
)

func NewAstTree(lines []string) Node {
	doc := ast.NewDocument()
	children := []Node{}
	for _, line := range lines {
		children = append(children, ParseLine(line))
	}
	doc.SetChildren(children)
	doc.SetRaw(strings.Join(lines, "\n"))
	return doc
}

func ParseLine(string) Node {
   
}
