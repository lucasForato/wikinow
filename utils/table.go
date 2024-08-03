package utils

import (
	"strings"
	"wikinow/internal/parser"

	sitter "github.com/smacker/go-tree-sitter"
)

func GetTableHeader(node *sitter.Node, lines []string) []string {
  child := node.NamedChild(0)
  content := []string{}

  for i := 0; i < int(child.ChildCount()); i++ {
    if child.Child(i).Type() == "pipe_table_cell" {
      text := parser.GetText(lines, child.Child(i))
      content = append(content, strings.TrimSpace(text))
    }
  }
  return content
}
