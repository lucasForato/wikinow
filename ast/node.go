package ast

type Node interface {
	AsContainer() *Container
	AsLeaf() *Leaf
	AsDocument() *Document
  GetRaw() string
  SetRaw(in string)
  GetStart() int
  GetEnd() int
	GetChildren() *[]Node
	SetChildren(newChildren []Node)
  AppendChild(newChild Node)
  GetType() string
  AsJSON() string
  Range() int
}
