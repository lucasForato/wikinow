package parser

import (
	"fmt"
	"strings"
)

type Paragraph struct {
	Container

	Content string
}

func NewParagraph(content string, parent Node) *Paragraph {
	paragraph := new(Paragraph)
	paragraph.Parent = parent
	paragraph.Content = content
	return paragraph
}

func (p *Paragraph) Print(spaces int) {
	tab := strings.Repeat(" ", spaces)
	fmt.Println(tab, "Type: Paragraph")
	fmt.Println(tab, "Content:", p.Content)
	p.Container.Print(spaces + 2)
}
