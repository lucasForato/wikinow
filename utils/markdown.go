package utils

import (
	"encoding/json"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	sitter "github.com/smacker/go-tree-sitter"

)

func ReadMarkdown(path string) []string {
	mainInput, err := os.ReadFile(path)
	if err != nil {
		log.WithFields(log.Fields{
			"source": path,
		}).Fatal("Error reading main documentation file.")
	}

	file := string(mainInput)

	return strings.Split(file, "\n")
}

func IndexOfSubstring(str, subStr string) int {
	for i := 0; i < len(str); i++ {
		if str[i:i+len(subStr)] == subStr {
			return i
		}
	}
	return -1
}

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

func ConvertTreeToJson(node *sitter.Node, lines []string) string {
	if node.IsNull() {
		return "[]"
	}

	maps := map[string]interface{}{}
	children := []interface{}{}
	count := int(node.ChildCount())

	for i := 0; i < count; i++ {
		child := ConvertTreeToJson(node.Child(i), lines)
		children = append(children, json.RawMessage(child))
	}

	maps[node.Type()] = map[string]interface{}{
		"content":  GetText(lines, node),
		"children": children,
	}

	item, err := json.Marshal(maps)
	if err != nil {
		log.Fatal("failed to parse json", err)
	}
	return string(item)
}
