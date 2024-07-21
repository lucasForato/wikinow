package parser

import (
	"fmt"
	"strings"
	"wikinow/ast"
)

type Italic struct {
	ast.Leaf
}

func (i *Italic) Print(spaces int) {
	tab := strings.Repeat(" ", spaces)
	fmt.Println(tab, "Type: Italic")
	fmt.Println(tab, "Content:", i.Content)
}
