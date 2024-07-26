package ast

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

type Container struct {
	Type     string
	Children []Node
	Raw      string
	Start    int
	End      int
	// Parent   Node
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

func (c *Container) GetRaw() string {
	return c.Raw
}

func (c *Container) SetRaw(in string) {
	c.Raw = in
}

func (c *Container) GetStart() int {
  return c.Start
}

func (c *Container) GetEnd() int {
  return c.End
}

func (c *Container) GetChildren() *[]Node {
	return &c.Children
}

func (c *Container) GetType() string {
  return c.Type
}

func (c *Container) AppendChild(child Node) {
  c.Children = append(c.Children, child)
}

func (c *Container) AppendChildren(children []Node) {
	c.Children = append(c.Children, children...)
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

func (c *Container) Range() int {
  return c.End - c.Start
}
