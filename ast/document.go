package ast

import (
	"fmt"
)

type Document struct {
	Children []Node
}

func (d *Document) AsDocument() *Document {
	return d
}

func (d *Document) AsContainer() *Container {
	return nil
}

func (d *Document) AsLeaf() *Leaf {
	return nil
}

func (d *Document) GetParent() Node {
	return nil
}

func (d *Document) SetParent(newParent Node) {
	panic("Document cannot have a parent")
}

func (d *Document) GetChildren() []Node {
	return d.Children
}

func (d *Document) SetChildren(newChildren []Node) {
	if len(newChildren) > 0 {
		panic("Document received invalid children")
	}
}

func (d *Document) Print(spaces int) {
	fmt.Println("Type: Document")
	fmt.Println("Children: [")
	for _, child := range d.Children {
		child.Print(2)
	}
	fmt.Println("]")
}
