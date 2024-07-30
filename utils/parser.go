package utils

import (
	"html/template"
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

func GetText(lines []string, node *sitter.Node) string {
	start := node.StartPoint()
	end := node.EndPoint()

	startLine := start.Row
	startColumn := start.Column
	endLine := end.Row
	endColumn := end.Column
	if startLine == endLine {
		return lines[startLine][startColumn:endColumn]
	}
	text := lines[startLine][startColumn:]
	for i := startLine + 1; i < endLine; i++ {
		text += lines[i]
	}
	text += lines[endLine][:endColumn]
	return text
}

func ParseInline(str string) template.HTML {
	// Parse bold markers (**)
	for {
		fromStart := strings.Index(str, "**")
		if fromStart == -1 {
			break
		}

		fromEnd := strings.Index(str[fromStart+2:], "**")
		if fromEnd == -1 {
			break
		}
		fromEnd += fromStart + 2

		str = str[:fromStart] + "<strong>" + str[fromStart+2:fromEnd] + "</strong>" + str[fromEnd+2:]
	}

	// Parse italic markers (*)
	for {
		fromStart := strings.Index(str, "*")
		if fromStart == -1 {
			break
		}

		fromEnd := strings.Index(str[fromStart+1:], "*")
		if fromEnd == -1 {
			break
		}
		fromEnd += fromStart + 1

		str = str[:fromStart] + "<i>" + str[fromStart+1:fromEnd] + "</i>" + str[fromEnd+1:]
	}

	return template.HTML(str)
}
