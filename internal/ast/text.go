package ast

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

func GetText(lines []string, node *sitter.Node) string {
	start := node.StartPoint()
	end := node.EndPoint()

	startRow := start.Row
	endRow := end.Row
	startCol := start.Column
	endCol := end.Column
	if startRow == endRow {
		return lines[startRow][startCol:endCol]
	}
	text := lines[startRow][startCol:]
	for i := startRow + 1; i < endRow; i++ {
		text += lines[i]
	}
	text += lines[endRow][:endCol]
	return text
}

func SplitQuote(node *sitter.Node, lines []string) []string {
	inline := node.Child(0)
	text := GetText(lines, inline)

	splits := []string{}
	for {
		index := strings.IndexRune(text, '\u003e')
		if index == -1 {
			splits = append(splits, text)
			break
		}
		splits = append(splits, text[:index])
		text = text[index+1:]
	}
	return splits
}
