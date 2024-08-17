package ast

import sitter "github.com/smacker/go-tree-sitter"

func IsOrderedList(node *sitter.Node) bool {
	if node.Type() != "list" || node.ChildCount() == 0 {
		return false
	}

	if node.Child(0).Type() == "list_item" {
		node = node.Child(0)
	}

	if node.ChildCount() == 0 {
		return false
	}

	if node.Child(0).Type() == "list_marker_parenthesis" || node.Child(0).Type() == "list_marker_dot" {
		return true
	}

	return false
}

func HasNestedList(node *sitter.Node) bool {
	for i := 0; i < int(node.ChildCount()); i++ {
		if node.Child(i).Type() == "list" {
			return true
		}
	}
	return false
}

func GetNestedListNode(node *sitter.Node) *sitter.Node {
	for i := 0; i < int(node.ChildCount()); i++ {
		if node.Child(i).Type() == "list" {
			return node.Child(i)
		}
	}
	return nil
}
