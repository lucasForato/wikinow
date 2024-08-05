package utils

import (
	"strings"
	"wikinow/types"

	sitter "github.com/smacker/go-tree-sitter"
)

func HasListChild(node *sitter.Node) bool {
	for {
		if node.NextSibling() == nil {
			break
		}
		if node.NextSibling().Type() == "list" {
			return true
		}
		node = node.NextSibling()
	}
	return false
}

func GetListChild(node *sitter.Node) *sitter.Node {
	for {
		if node.NextSibling() == nil {
			break
		}
		if node.NextSibling().Type() == "list" {
			return node.NextSibling()
		}
		node = node.NextSibling()
	}
	return nil
}

func IsOrderedList(node *sitter.Node) bool {
	if node.Type() != "list" || node.ChildCount() == 0 {
		return false
	}

	if node.Child(0).Type() == "list_item" {
		node = node.Child(0)
	}

	if node.ChildCount() == 0 {
		return false
	}

	if node.Child(0).Type() == "list_marker_parenthesis" || node.Child(0).Type() == "list_marker_dot" {
		return true
	}

	return false
}

func HasNestedList(node *sitter.Node) bool {
	for i := 0; i < int(node.ChildCount()); i++ {
		if node.Child(i).Type() == "list" {
			return true
		}
	}
	return false
}

func GetNestedListNode(node *sitter.Node) *sitter.Node {
	for i := 0; i < int(node.ChildCount()); i++ {
		if node.Child(i).Type() == "list" {
			return node.Child(i)
		}
	}
	return nil
}

func GetParagraphChild(node *sitter.Node) *sitter.Node {
	for i := 0; i < int(node.ChildCount()); i++ {
		if node.Child(i).Type() == "paragraph" {
			return node.Child(i)
		}
	}
	return nil
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

func NextSiblingIsBlockContinuation(node *sitter.Node) bool {
	if node.NextSibling() == nil {
		return false
	}
	if node.NextSibling().Type() == "block_continuation" {
		return true
	}
	return false
}

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

func GetLanguage(node *sitter.Node, lines []string) types.Language {
	for i := 0; i < int(node.NamedChildCount()); i++ {
		if node.NamedChild(i).Type() == "info_string" {
			return GetText(lines, node.NamedChild(i))
		}
	}
	return ""
}

func IsLanguage(node *sitter.Node, language string, lines []string) bool {
	return GetText(lines, node) == language
}
