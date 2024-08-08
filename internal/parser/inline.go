package parser

import (
	"fmt"
	"html/template"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Extra string

const (
	Title Extra = "title"
)

func ParseInline(str string, c *Ctx, extra *[]Extra) template.HTML {
	str = parseBold(str)
	str = parseItalic(str)
	str = parseImage(str)
	str = parseRefImage(str, c)
	str = parseInlineLink(str)
	str = parseRefLink(str, c)
	str = parseVariable(str, c)
	str = parseInlineCode(str)
	str = parseCodeBlock(str, c, new(RealFileReader))
	str = parseLinkToAnotherFile(str)
	str = parseRefFootnote(str)

	if extra != nil {
		if slices.Contains(*extra, "title") {
			str = parseTitle(str)
		}
	}

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

func parseRefLink(str string, c *Ctx) string {
	re := regexp.MustCompile(`\[([^\]]+)\]\[([^\]]+)\]`)

	match := re.FindStringSubmatch(str)
	if len(match) == 0 {
		return str
	}
	name := match[2]
	value, ok := ReadCtx(c, name)
	if !ok {
		return str
	}
	str = re.ReplaceAllString(str, `<a href="`+value+`" class="text-amber-600" target="_blank">$1</a>`)
	return str
}

func parseRefFootnote(str string) string {
  re := regexp.MustCompile(`\[\^([^\]]+)\]`)
  match := re.FindStringSubmatch(str)
	if len(match) == 0 {
		return str
	}
	str = re.ReplaceAllString(str, `<a href="#`+match[1]+`" class="text-amber-600">$1</a>`)
	return str
}

func parseVariable(str string, c *Ctx) string {
	for {
		fromStart := strings.Index(str, "$var(")
		if fromStart == -1 {
			break
		}

		fromEnd := strings.Index(str, ")")
		if fromEnd == -1 {
			break
		}

		varName := str[fromStart+5 : fromEnd]
		varValue, ok := ReadCtx(c, varName)
		if !ok {
			varValue = ""
		}

		str = str[:fromStart] + varValue + str[fromEnd+1:]
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

func parseRefImage(str string, c *Ctx) string {
	re := regexp.MustCompile(`!\[([^\]]+)\]\[([^\]]+)\]`)
	match := re.FindStringSubmatch(str)
	if len(match) == 0 {
		return str
	}
	name := match[2]
	value, ok := ReadCtx(c, name)
	if !ok {
		return str
	}
	str = re.ReplaceAllString(str, `<img src="`+value+`" alt="$1" class="w-full" />`)
	return str
}

type FileReader interface {
	ReadFile(path string) ([]byte, error)
}

type RealFileReader struct{}

func (r RealFileReader) ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func parseCodeBlock(str string, c *Ctx, reader FileReader) string {
	for {
		fromStart := strings.Index(str, "$code(")
		if fromStart == -1 {
			break
		}

		fromEnd := strings.Index(str, ")")
		if fromEnd == -1 {
			break
		}

		within := str[fromStart+6 : fromEnd]
		split := strings.Split(within, ",")
		for i := range split {
			split[i] = strings.TrimSpace(split[i])
		}

		path, ok := ReadCtx(c, split[0])
		if !ok {
			log.Fatal("Variable not found")
		}

		bytes, err := reader.ReadFile(path)
		if err != nil {
			log.WithFields(log.Fields{
				"source": path,
			}).Fatal("Error reading documentation file.")
		}
		file := string(bytes)
		lines := strings.Split(file, "\n")
		start, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal("Error parsing start line")
		}

		end, err := strconv.Atoi(split[2])
		if err != nil {
			log.Fatal("Error parsing end line")
		}

		lines = lines[start:end]

		code := strings.Join(lines, "\n")

		html := fmt.Sprintf(`<pre class="bg-zinc-700 p-1 rounded text-amber-600"><code class="%s">%s</code></pre>`, getFileType(path), code)
		return html
	}
	return str
}

func getFileType(path string) string {
	split := strings.Split(path, ".")
	fileType := split[len(split)-1]
	switch fileType {
	case "md":
		return "language-markdown"
	case "go":
		return "language-go"
	default:
		return "text"
	}
}

func parseLinkToAnotherFile(str string) string {
	for {
		fromStart := strings.Index(str, "$link(")
		if fromStart == -1 {
			break
		}

		fromEnd := strings.Index(str, ")")
		if fromEnd == -1 {
			break
		}

		within := str[fromStart+6 : fromEnd]
		split := strings.Split(within, ",")
		for i := range split {
			split[i] = strings.TrimSpace(split[i])
		}

		path := split[1]

		link := fmt.Sprintf(`<a href="%s" class="text-amber-600">%s</a>`, path, split[0])
		str = str[:fromStart] + link + str[fromEnd+1:]

	}
	return str
}

func parseTitle(str string) string {
	return strings.Replace(str, "#", "", -1)
}
