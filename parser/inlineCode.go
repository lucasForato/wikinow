package parser

import (
	"fmt"
	"strings"
)

type InlineCode struct {
	Leaf
}

func (ic *InlineCode) Print(spaces int) {
	tab := strings.Repeat(" ", spaces)
	fmt.Println(tab, "Type: InlineCode")
	fmt.Println(tab, "Content:", ic.Content)
}
