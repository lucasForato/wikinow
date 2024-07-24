package ast

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

type Container struct {
	Type     string
	// Parent   Node
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
//
// func (c *Container) GetParent() Node {
// 	return c.Parent
// }
//
// func (c *Container) SetParent(newParent Node) {
// 	c.Parent = newParent
// }

func (c *Container) GetChildren() []Node {
	return c.Children
}

func (c *Container) SetChildren(newChildren []Node) {
	if len(newChildren) == 0 {
		panic("children received an invalid value")
	}
	c.Children = newChildren
}

func (c *Container) AsJSON() string {
	b, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
