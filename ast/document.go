package ast

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

type Document struct {
	Type     string
	Children []Node
	Raw      string
	Start    int
	End      int
}

func NewDocument() *Document {
	doc := new(Document)
	doc.Type = "Document"
  return doc
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

func (d *Document) GetRaw() string {
	return d.Raw
}

func (d *Document) SetRaw(in string) {
	d.Raw = in
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

func (d *Document) AsJSON() string {
	children := make([]string, len(d.Children))
	for _, child := range d.Children {
		json := child.AsJSON()
		children = append(children, json)
	}

	b, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
