package parser

import (
	"regexp"
)

type Header struct {
	Container
	Level int
}

func (h *Header) AsHeader() *Header {
	return h
}

func NewHeader(raw string, content string, start int, end int, level int, ctx *Ctx) Node {
	header := new(Header)
	header.Type = "Header"
	header.Raw = raw
	header.Start = start
	header.End = end
	children, _ := Parse(content, ctx)
	if children != nil {
		header.SetChildren(*children)
	}
	header.Level = level
	return header
}

func ParseHeader(in string, ctx *Ctx) *[]Node {
	if h6 := ParseH6(in, ctx); h6 != nil {
		return h6
	}
	if h5 := ParseH5(in, ctx); h5 != nil {
		return h5
	}
	if h4 := ParseH4(in, ctx); h4 != nil {
		return h4
	}
	if h3 := ParseH3(in, ctx); h3 != nil {
		return h3
	}
	if h2 := ParseH2(in, ctx); h2 != nil {
		return h2
	}
	if h1 := ParseH1(in, ctx); h1 != nil {
		return h1
	}
	return nil
}

func ParseH1(in string, ctx *Ctx) *[]Node {
	regex := regexp.MustCompile(`(#{1}\s)(.*)`)
	segments := regex.FindAllStringSubmatchIndex(in, -1)
	if len(segments) == 0 {
		return nil
	}

	result := []Node{}
	for _, match := range segments {
		bold := NewHeader(
			in[match[0]:match[1]],
			in[match[4]:match[5]],
			match[0],
			match[1],
			1,
      ctx,
		)
		result = append(result, bold)
	}
	return &result
}

func ParseH2(in string, ctx *Ctx) *[]Node {
	regex := regexp.MustCompile(`(#{2}\s)(.*)`)
	segments := regex.FindAllStringSubmatchIndex(in, -1)
	if len(segments) == 0 {
		return nil
	}

	result := []Node{}
	for _, match := range segments {
		bold := NewHeader(
			in[match[0]:match[1]],
			in[match[4]:match[5]],
			match[0],
			match[1],
			2,
      ctx,
		)
		result = append(result, bold)
	}
	return &result
}

func ParseH3(in string, ctx *Ctx) *[]Node {
	regex := regexp.MustCompile(`(#{3}\s)(.*)`)
	segments := regex.FindAllStringSubmatchIndex(in, -1)
	if len(segments) == 0 {
		return nil
	}

	result := []Node{}
	for _, match := range segments {
		bold := NewHeader(
			in[match[0]:match[1]],
			in[match[4]:match[5]],
			match[0],
			match[1],
			3,
      ctx,
		)
		result = append(result, bold)
	}
	return &result
}

func ParseH4(in string, ctx *Ctx) *[]Node {
	regex := regexp.MustCompile(`(#{4}\s)(.*)`)
	segments := regex.FindAllStringSubmatchIndex(in, -1)
	if len(segments) == 0 {
		return nil
	}

	result := []Node{}
	for _, match := range segments {
		bold := NewHeader(
			in[match[0]:match[1]],
			in[match[4]:match[5]],
			match[0],
			match[1],
			4,
      ctx,
		)
		result = append(result, bold)
	}
	return &result
}

func ParseH5(in string, ctx *Ctx) *[]Node {
	regex := regexp.MustCompile(`(#{5}\s)(.*)`)
	segments := regex.FindAllStringSubmatchIndex(in, -1)
	if len(segments) == 0 {
		return nil
	}

	result := []Node{}
	for _, match := range segments {
		bold := NewHeader(
			in[match[0]:match[1]],
			in[match[4]:match[5]],
			match[0],
			match[1],
			5,
      ctx,
		)
		result = append(result, bold)
	}
	return &result
}

func ParseH6(in string, ctx *Ctx) *[]Node {
	regex := regexp.MustCompile(`(#{6}\s)(.*)`)
	segments := regex.FindAllStringSubmatchIndex(in, -1)
	if len(segments) == 0 {
		return nil
	}

	result := []Node{}
	for _, match := range segments {
		bold := NewHeader(
			in[match[0]:match[1]],
			in[match[4]:match[5]],
			match[0],
			match[1],
			6,
      ctx,
		)
		result = append(result, bold)
	}
	return &result
}
