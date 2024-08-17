package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"strings"
)

// This function removes the class attribute from an HTML tag.
//
// Example: `<strong class="bold">bold</strong>` -> `<strong>bold</strong>`
func RemoveClass(s template.HTML) string {
	str := string(s)
	var result strings.Builder
	for {
		start := strings.Index(str, ` class="`)
		if start == -1 {
			result.WriteString(str)
			break
		}
		result.WriteString(str[:start])
		str = str[start:]
		end := strings.Index(str, `"`)
		str = str[end+1:]
		end = strings.Index(str, `"`)
		if end == -1 {
			result.WriteString(str)
			break
		}
		str = str[end+1:]
	}
	return result.String()
}

func JsonPrettyPrint(in string) *error {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "  ")
	if err != nil {
		return &err
	}
	fmt.Print(out.String())
	return nil
}

// PrintTreeAsJSON prints the TreeNode as JSON
func (n *TreeNode) PrintTreeAsJSON() {
	jsonData, err := json.MarshalIndent(n, "", "  ")
	if err != nil {
		log.Fatal("Error while marshalling tree to JSON:", err)
	}
	log.Println(string(jsonData))
}

func ParseAnchor(node TreeNode) template.HTML {
	return template.HTML(fmt.Sprintf(`<a href="%s" class="text-amber-600">%s</a>`, node.RelativePath, node.Title))
}

func GetIconClasses(level int) string {
	str := fmt.Sprintf("%dpx", (level * 8))
	return fmt.Sprintf("ml-[%s]", str)
}
