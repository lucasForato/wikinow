package ast

type Node interface {
	AsContainer() *Container
	AsLeaf() *Leaf
	AsDocument() *Document
	GetParent() Node
	SetParent(newParent Node)
	GetChildren() []Node
	SetChildren(newChildren []Node)
	Print(spaces int)
}
