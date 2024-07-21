package parser

import (
	"fmt"
	"strings"
)

type Bold struct {
	Leaf
}

func (b *Bold) Print(spaces int) {
	tab := strings.Repeat(" ", spaces)
	fmt.Println(tab, "Type: Bold")
	fmt.Println(tab, "Content:", b.Content)
}
