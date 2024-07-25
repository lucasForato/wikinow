package ast

type Node interface {
	AsContainer() *Container
	AsLeaf() *Leaf
	AsDocument() *Document
  GetRaw() string
  SetRaw(in string)
	GetChildren() []Node
	SetChildren(newChildren []Node)
  AppendChild(newChild Node)
  AsJSON() string
}
