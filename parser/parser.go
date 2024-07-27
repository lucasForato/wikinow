package parser

import (
	"fmt"
	"sort"
	"strings"
	"wikinow/ast"
)

type (
	Node      = ast.Node
	Container = ast.Container
	Document  = ast.Document
	Leaf      = ast.Leaf
)

func NewAstTree(lines []string) Node {
	doc := ast.NewDocument()
	for _, line := range lines {
		children := Parse(line)
		if children == nil {
			continue
		}
		doc.AppendChildren(*children)
	}
	doc.SetRaw(strings.Join(lines, "\n"))
	return doc
}

func Parse(in string) *[]Node {
	result := []Node{}

	if header := ParseHeader(in); header != nil {
		result = append(result, *header...)
	}

	if bold := ParseBold(in); bold != nil {
		result = append(result, *bold...)
	}

	if italic := ParseItalic(in); italic != nil {
		result = append(result, *italic...)
	}

	if text := ParseText(in, result); text != nil {
		for _, item := range result {
			fmt.Println(item.GetType(), ":", item.GetRaw())
		}
		result = append(result, *text...)
	}

	if len(result) == 0 {
		return nil
	}

	if len(result) > 1 {
		return GroupNodes(&result)
	}

	return &result
}

func GroupNodes(nodes *[]Node) *[]Node {
	if len(*nodes) == 1 {
		return nodes
	}

	sorted := make([]Node, len(*nodes))
	copy(sorted, *nodes)
	sort.Sort(ByDiff(sorted))

	p1 := 0
	p2 := 1

	for {
		if p2 >= len(sorted) {
			p1++
			p2 = p1 + 1
		}

		if p1 >= len(sorted)-1 {
			break
		}

		a := sorted[p1]
		b := sorted[p2]

		if b.AsLeaf() != nil {
			p2++
			if p2 >= len(sorted) {
				p1++
				p2 = p1 + 1
			}
			continue
		}

		if a.GetStart() >= b.GetStart() && a.GetEnd() <= b.GetEnd() {
			b.AsContainer().AppendChild(a)
			sorted = append(sorted[:p1], sorted[p1+1:]...)
			p2 = p1 + 1
			continue
		}
		p2++
	}

	for _, node := range sorted {
		container := node.AsContainer()
		if container == nil {
			continue
		}

		GroupNodes(container.GetChildren())
	}

	return &sorted
}

type ByDiff []Node

func (a ByDiff) Len() int           { return len(a) }
func (a ByDiff) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDiff) Less(i, j int) bool { return a[i].Range() < a[j].Range() }
