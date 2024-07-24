package ast

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

type Leaf struct {
	// Parent   Node
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
//
// func (l *Leaf) GetParent() Node {
// 	return l.Parent
// }
//
// func (l *Leaf) SetParent(newParent Node) {
// 	l.Parent = newParent
// }

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
