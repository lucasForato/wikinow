package ast

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

type Leaf struct {
	Type    string
	Content string
	Raw     string
	Start   int
	End     int
	// Parent   Node
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

func (l *Leaf) AppendChild(child Node) {
	panic("Leaf cannot have children")
}

func (l *Leaf) AppendChildren(children []Node) {
	panic("Leaf cannot have children")
}

func (l *Leaf) GetRaw() string {
	return l.Raw
}

func (l *Leaf) SetRaw(in string) {
	l.Raw = in
}

func (l *Leaf) GetChildren() []Node {
	return nil
}

func (l *Leaf) SetChildren(newChildren []Node) {
	if len(newChildren) != 0 {
		panic("leaf node cannot have children")
	}
}

func (l *Leaf) AsJSON() string {
	b, err := json.Marshal(l)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
