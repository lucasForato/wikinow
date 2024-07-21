package parser

import (
	"fmt"
	"strings"
)

type Link struct {
	Leaf
	Url string
}

func (l *Link) Print(spaces int) {
	tab := strings.Repeat(" ", spaces)
	fmt.Println(tab, "Type: Link")
	fmt.Println(tab, "Url:", l.Url)
	fmt.Println(tab, "Content:", l.Content)
}
