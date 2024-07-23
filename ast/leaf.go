package ast

import (
	"fmt"
	"strings"
)

type Leaf struct {
	Parent   Node
	Content  string
	Original string
}

func (l *Leaf) AsDocument() *Document {
	return nil
}

func (l *Leaf) AsContainer() *Container {
	return nil
}

func (l *Leaf) AsLeaf() *Leaf {
	return l
}

func (l *Leaf) GetParent() Node {
	return l.Parent
}

func (l *Leaf) SetParent(newParent Node) {
	l.Parent = newParent
}

func (l *Leaf) GetChildren() []Node {
	return nil
}

func (l *Leaf) SetChildren(newChildren []Node) {
	if len(newChildren) != 0 {
		panic("leaf node cannot have children")
	}
}

func (l *Leaf) Print(spaces int) {
	tab := strings.Repeat(" ", spaces)

	fmt.Println(tab, "Content:", l.Content)
}
