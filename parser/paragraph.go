package parser

import (
	"fmt"
	"strings"
	"wikinow/ast"
)

type Paragraph struct {
	ast.Container

	Content string
}

func (p *Paragraph) Print(spaces int) {
	tab := strings.Repeat(" ", spaces)
	fmt.Println(tab, "Type: Paragraph")
	fmt.Println(tab, "Content:", p.Content)
	p.Container.Print(spaces + 2)
}
