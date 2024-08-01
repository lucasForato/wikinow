package utils

import sitter "github.com/smacker/go-tree-sitter"

type Language = string

const (
  Markdown Language = "markdown"
  JavaScript Language = "javascript"
  Go Language = "go"
  Rust Language = "rust"
  Ruby Language = "ruby"
  TypeScript Language = "typescript"
  Python Language = "python"
  C Language = "c"
  Cpp Language = "cpp"
  CSharp Language = "csharp"
  PHP Language = "php"
  HTML Language = "html"
  CSS Language = "css"
)

func GetLanguage(node *sitter.Node, lines []string) Language {
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