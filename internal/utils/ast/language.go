package ast

import sitter "github.com/smacker/go-tree-sitter"

type Language = string

const (
	Markdown   Language = "markdown"
	JavaScript Language = "javascript"
	TypeScript Language = "typescript"
	Go         Language = "go"
	Rust       Language = "rust"
	Ruby       Language = "ruby"
	Python     Language = "python"
	C          Language = "c"
	Cpp        Language = "cpp"
	CSharp     Language = "csharp"
	PHP        Language = "php"
	HTML       Language = "html"
	CSS        Language = "css"
)

func GetLanguage(node *sitter.Node, lines []string) Language {
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
