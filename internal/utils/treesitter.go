package utils

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func HasListChild(node *sitter.Node) bool {
  for {
    if node.NextSibling() == nil {
      break
    }

    if node.NextSibling().Type() == "list" {
      return true
    }

    node = node.NextSibling()
  }

  return false
}

func GetListChild(node *sitter.Node) *sitter.Node {
  for {
    if node.NextSibling() == nil {
      break
    }
    if node.NextSibling().Type() == "list" {
      return node.NextSibling()
    }
    node = node.NextSibling()
  }
  return nil
}
