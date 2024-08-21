package filetree

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
)

// PrintTreeAsJSON prints the TreeNode as JSON
func (n *TreeNode) PrintTreeAsJSON() {
	jsonData, err := json.MarshalIndent(n, "", "  ")
	if err != nil {
		log.Fatal("Error while marshalling tree to JSON:", err)
	}
	log.Println(string(jsonData))
}

func ParseAnchor(node TreeNode) template.HTML {
  return template.HTML(fmt.Sprintf(`<a href="%s" class="text-orange-400">%s</a>`, node.RelativePath, node.Title))
}
