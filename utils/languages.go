package utils

import (
	"wikinow/types"

	sitter "github.com/smacker/go-tree-sitter"
)

func GetLanguage(node *sitter.Node, lines []string) types.Language {
	for i := 0; i < int(node.NamedChildCount()); i++ {
		if node.NamedChild(i).Type() == "info_language" {
			return GetText(lines, node)
		}
	}
  return ""
}

func IsLanguage(node *sitter.Node, language string, lines []string) bool {
	return GetText(lines, node) == language
}
