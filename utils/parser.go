package utils

import (
	"context"
	"html/template"
	"regexp"
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

func GetText(lines []string, node *sitter.Node) string {
	start := node.StartPoint()
	end := node.EndPoint()

	startRow := start.Row
	endRow := end.Row
	startCol := start.Column
	endCol := end.Column
	if startRow == endRow {
		return lines[startRow][startCol:endCol]
	}
	text := lines[startRow][startCol:]
	for i := startRow + 1; i < endRow; i++ {
		text += lines[i]
	}
	text += lines[endRow][:endCol]
	return text
}

func GetCode(lines []string, node *sitter.Node) string {
  start := node.StartPoint()
  end := node.EndPoint()

  startRow := start.Row
  startColumn := start.Column
  endRow := end.Row
  endColumn := end.Column
  if startRow == endRow {
    return lines[startRow][startColumn:endColumn]
  }
  allLines := lines[startRow:endRow+1] 
  allLines[0] = allLines[0][startColumn:]
  allLines[len(allLines)-1] = allLines[len(allLines)-1][:endColumn]

  text := strings.Join(allLines, "\n")
  return text
}

func ParseInline(str string, ctx *context.Context) template.HTML {
	str = parseBold(str)
	str = parseItalic(str)
  str = parseImage(str)
	str = parseInlineLink(str)
	str = parseVariable(str, ctx)
  str = parseRefLink(str, ctx)
  str = parseInlineCode(str)

	return template.HTML(str)
}

func parseBold(str string) string {
	for {
		fromStart := strings.Index(str, "**")
		if fromStart == -1 {
			fromStart = strings.Index(str, "__")
		}
		if fromStart == -1 {
			break
		}

		fromEnd := strings.Index(str[fromStart+2:], str[fromStart:fromStart+2])
		if fromEnd == -1 {
			break
		}
		fromEnd += fromStart + 2

		str = str[:fromStart] + "<strong class=\"font-bold\">" + str[fromStart+2:fromEnd] + "</strong>" + str[fromEnd+2:]
	}
	return str
}

func parseItalic(str string) string {
	for {
		fromStart := strings.Index(str, "*")
		if fromStart == -1 {
			fromStart = strings.Index(str, "_")
		}
		if fromStart == -1 {
			break
		}

		fromEnd := strings.Index(str[fromStart+1:], str[fromStart:fromStart+1])
		if fromEnd == -1 {
			break
		}
		fromEnd += fromStart + 1

		str = str[:fromStart] + "<i class=\"italic\">" + str[fromStart+1:fromEnd] + "</i>" + str[fromEnd+1:]
	}
	return str
}

func parseInlineLink(str string) string {
	re := regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)
	str = re.ReplaceAllString(str, `<a href="$2" class="text-amber-600" target="_blank">$1</a>`)

	return str
}

func parseRefLink(str string, ctx *context.Context) string {
  re := regexp.MustCompile(`\[([^\]]+)\]\[([^\]]+)\]`)

  match := re.FindStringSubmatch(str)
  if len(match) == 0 {
    return str
  }
  name := match[2]
  value, ok := GetFromContext(ctx, name)
  if !ok {
    return str 
  }
  str = re.ReplaceAllString(str, `<a href="` + value + `" class="text-amber-600" target="_blank">$1</a>`)
  return str
}

func parseVariable(str string, ctx *context.Context) string {
	for {
		fromStart := strings.Index(str, "{{")
		if fromStart == -1 {
			break
		}

		fromEnd := strings.Index(str, "}}")
		if fromEnd == -1 {
			break
		}

		varName := str[fromStart+2 : fromEnd]
		varValue, ok := GetFromContext(ctx, varName)
		if !ok {
			varValue = ""
		}

		str = str[:fromStart] + varValue + str[fromEnd+2:]
	}
	return str
}

func parseInlineCode(str string) string {
  re := regexp.MustCompile("`([^`]+)`")
  str = re.ReplaceAllString(str, `<code class="bg-zinc-700 p-1 rounded text-amber-600">$1</code>`)
  return str
}

func parseImage(str string) string {
  re := regexp.MustCompile(`!\[([^\]]+)\]\(([^)]+)\)`)
  str = re.ReplaceAllString(str, `<img src="$2" alt="$1" class="w-full" />`)
  return str
}
