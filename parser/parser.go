package parser

import (
	"wikinow/ast"
	"wikinow/utils"
)

type Node = ast.Node
type Container = ast.Container
type Document = ast.Document
type Leaf = ast.Leaf

func NewAstTree(lines []string) Node {
	doc := new(ast.Document)

	for _, line := range lines {
		if utils.IsHeader(line) {
			content, level := utils.GetHeaderContent(line)
			header := new(Header)
			header.Parent = doc
			header.Children = ParseHeader(content, header)
			header.Level = level
			doc.Children = append(doc.Children, header)
		}
	}
	return doc
}
