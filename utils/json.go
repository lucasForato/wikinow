package utils

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
	sitter "github.com/smacker/go-tree-sitter"
)

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
