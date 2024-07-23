package parser

import (
	"fmt"
	"strings"
)

type Bold struct {
	Container

	Content string
}

func NewBold(content string, parent Node) *Bold {
	bold := new(Bold)
	bold.Parent = parent
	bold.Content = strings.Trim(content, "**")
	return bold
}

func (b *Bold) Print(spaces int) {
	tab := strings.Repeat(" ", spaces)
	fmt.Println(tab, "Type: Bold")
	fmt.Println(tab, "Content:", b.Content)
}
