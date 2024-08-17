package ast

import (
	"encoding/json"
	"log"

	sitter "github.com/smacker/go-tree-sitter"
)

func ConvertTreeToJson(node *sitter.Node, lines []string) string {
	if node.IsNull() {
		return "[]"
	}

	maps := map[string]interface{}{}
	children := []interface{}{}
	count := int(node.NamedChildCount())

	for i := 0; i < count; i++ {
		child := ConvertTreeToJson(node.NamedChild(i), lines)
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
