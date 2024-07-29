package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strings"
	"wikinow/utils"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	sitter "github.com/smacker/go-tree-sitter"
	markdown "github.com/smacker/go-tree-sitter/markdown/tree-sitter-markdown"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Initializes wikinow by adding a /wiki directory and configuration file",
	Long: `Initialization command:

  This command will create a config file at the current directory and a /wiki
  directory with a main.md file in it.

  You should run this command before using any other command. 
  `,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Starting server...")
		main()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func main() {
	app := echo.New()
	app.GET("/*", handler)
	app.GET("/favicon.ico", func(ctx echo.Context) error {
		return ctx.NoContent(http.StatusNoContent)
	})
	app.Logger.Fatal(app.Start(":4000"))
}

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}

func handler(ctx echo.Context) error {
	path := handlePath(ctx)
	lines := utils.ReadMarkdown(path)

	parser := sitter.NewParser()
	parser.SetLanguage(markdown.GetLanguage())

	sourceCode := []byte(strings.Join(lines, "\n"))
	tree, err := parser.ParseCtx(context.Background(), nil, sourceCode)
	if err != nil {
		log.Fatal("Failed to parse source code", err)
	}

	root := tree.RootNode()
	str := ConvertTreeToJson(root, lines)
	utils.JsonPrettyPrint(str)

	return Render(ctx, http.StatusOK, nil)
}

func handlePath(ctx echo.Context) string {
	url := ctx.Request().URL.Path
	if url == "/" {
		url = "/main.md"
	}
	return path.Join("wiki", url)
}

func getTextFromPoints(lines []string, start sitter.Point, end sitter.Point) string {
	startLine := start.Row
	startColumn := start.Column
	endLine := end.Row
	endColumn := end.Column
	if startLine == endLine {
		return lines[startLine][startColumn:endColumn]
	}
	text := lines[startLine][startColumn:]
	for i := startLine + 1; i < endLine; i++ {
		text += lines[i]
	}
	text += lines[endLine][:endColumn]
	return text
}

func ConvertTreeToJson(node *sitter.Node, lines []string) string {
	if node.IsNull() {
		return "[]"
	}

	maps := map[string]interface{}{}
	children := []interface{}{}
	count := int(node.ChildCount())

	for i := 0; i < count; i++ {
    fmt.Println(node.Child(i))
    fmt.Println(node.String())

		child := ConvertTreeToJson(node.Child(i), lines)
		children = append(children, json.RawMessage(child))
	}

	maps[node.Type()] = map[string]interface{}{
		"content":  getTextFromPoints(lines, node.StartPoint(), node.EndPoint()),
		"children": children,
	}

	item, err := json.Marshal(maps)
	if err != nil {
		log.Fatal("failed to parse json", err)
	}
	return string(item)
}
