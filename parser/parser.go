package parser

import (
	"strings"
	"wikinow/ast"
	"wikinow/utils"
)

type (
	Node      = ast.Node
	Container = ast.Container
	Document  = ast.Document
	Leaf      = ast.Leaf
)

func NewAstTree(lines []string) Node {
	doc := ast.NewDocument()
	for _, line := range lines {
		children := Parse(line)
		if children == nil {
			continue
		}
		doc.AppendChildren(*children)
	}
	doc.SetRaw(strings.Join(lines, "\n"))
	utils.JsonPrettyPrint(doc.AsJSON())
	return doc
}

func Parse(in string) *[]Node {
	result := []Node{}

	if bold := ParseBold(in); bold != nil {
		result = append(result, *bold...)
	}
	if header := ParseHeader(in); header != nil {
		result = append(result, *header...)
	}
	if len(result) == 0 {
		return nil
	}

	return &result
}
