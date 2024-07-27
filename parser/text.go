package parser

import (
	"strings"
	"wikinow/utils"
)

type Text struct {
	Leaf
}

func NewText(raw string, content string, start int, end int) Node {
	text := new(Text)
	text.Type = "Text"
	text.Raw = raw
	text.Content = content
	text.Start = start
	text.End = end
	return text
}


func ParseText(content string, nodes []Node) *[]Node {
	i := 0
	result := []Node{}

	for _, node := range nodes {
		raw := node.GetRaw()
		if strings.Contains(content, raw) {
			index := utils.IndexOfSubstring(content[i:], raw)
			if index != -1 {
				index += i  // Adjust index relative to the entire content
				subStr := content[i:index]
				text := NewText(subStr, subStr, i, index)
				result = append(result, text)
				i = index + len(raw)  // Update i to the end of the current node
			}
		}
	}

	// Append the remaining part of the content after the last node
	if i < len(content) {
		subStr := content[i:]
		text := NewText(subStr, subStr, i, len(content))
		result = append(result, text)
	}

	return &result
}

