package cmd

import (
	"context"
	"net/http"
	"path"
	"strings"
	"wikinow/component"
	"wikinow/internal/parser"
	"wikinow/internal/utils"

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

	astParser := sitter.NewParser()
	astParser.SetLanguage(markdown.GetLanguage())

	sourceCode := []byte(strings.Join(lines, "\n"))

	c := parser.CreateCtx()
	if err := parser.LoadCtx(c, &lines); err != nil {
		return Render(ctx, http.StatusInternalServerError, component.Error(err))
	}

	tree, err := astParser.ParseCtx(context.Background(), nil, sourceCode)
	if err != nil {
		log.Fatal("Failed to parse source code", err)
	}

	root := tree.RootNode()
	str := utils.ConvertTreeToJson(root, lines)
	utils.JsonPrettyPrint(str)

	return Render(ctx, http.StatusOK, component.Layout(root, &lines, c))
}

func handlePath(ctx echo.Context) string {
	url := ctx.Request().URL.Path

	if url == "/" {
		url = "/main"
	}
	file := strings.Join([]string{url, "md"}, ".")
	return path.Join("wiki", file)
}
