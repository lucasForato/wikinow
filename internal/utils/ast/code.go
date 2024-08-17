package ast

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

func GetCode(lines []string, node *sitter.Node) string {
	start := node.StartPoint()
	end := node.EndPoint()

	startRow := start.Row
	startColumn := start.Column
	endRow := end.Row
	endColumn := end.Column
	if startRow == endRow {
		return lines[startRow][startColumn:endColumn]
	}
	allLines := lines[startRow : endRow+1]
	allLines[0] = allLines[0][startColumn:]
	allLines[len(allLines)-1] = allLines[len(allLines)-1][:endColumn]

	text := strings.Join(allLines, "\n")
	return text
}

func GetIndentedCode(lines []string, node *sitter.Node) string {
	lineCount := int(node.ChildCount())
	content := GetText(lines, node)
	indentText := GetText(lines, node.NamedChild(0))
	content = strings.Replace(content, indentText, "\n", lineCount)
	content = strings.Replace(content, "\n", "", 1)

	return content
}
