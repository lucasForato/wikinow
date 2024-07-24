package parser

type Paragraph struct {
	Container
}

func NewParagraph(content string) *Paragraph {
	paragraph := new(Paragraph)
	paragraph.Type = "Paragraph"
	paragraph.SetRaw(content)
	return paragraph
}
