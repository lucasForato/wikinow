package ast

import (
	"fmt"
	"strings"
)

type Container struct {
	Parent   Node
	Children []Node
}

func (c *Container) AsDocument() *Document {
	return nil
}

func (c *Container) AsContainer() *Container {
	return c
}

func (c *Container) AsLeaf() *Leaf {
	return nil
}

func (c *Container) GetParent() Node {
	return c.Parent
}

func (c *Container) SetParent(newParent Node) {
	c.Parent = newParent
}

func (c *Container) GetChildren() []Node {
	return c.Children
}

func (c *Container) SetChildren(newChildren []Node) {
	if len(newChildren) == 0 {
		panic("children received an invalid value")
	}
	c.Children = newChildren
}

func (c *Container) Print(spaces int) {
	tab := strings.Repeat(" ", spaces)

	fmt.Println(tab, "{")
	fmt.Println(tab, "  Type: Container")
	fmt.Println(tab, "  Children: [")
	for _, child := range c.Children {
		child.Print(spaces + 4)
	}
	fmt.Println(tab, "  ]")
	fmt.Println(tab, "}")
}
