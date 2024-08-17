// package ast provides utilities for working with the AST (Abstract Syntax Tree) of a markdown file.
package ast

import (
	"regexp"

	sitter "github.com/smacker/go-tree-sitter"
)

func NextSiblingIsBlockContinuation(node *sitter.Node) bool {
	if node.NextSibling() == nil {
		return false
	}
	if node.NextSibling().Type() == "block_continuation" {
		return true
	}
	return false
}

func IsFootnoteRef(node *sitter.Node, lines *[]string) bool {
	str := GetText(*lines, node)
	re := regexp.MustCompile(`\[\^([^\]]+)\]:`)
	match := re.FindStringSubmatch(str)
	if len(match) > 0 {
		return true
	}
	return false
}
