package cmd

import (
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"wikinow/component"
	"wikinow/internal/filetree"
	"wikinow/internal/parser"
	"wikinow/internal/utils"

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
		e := echo.New()
		// mux := http.NewServeMux()

		e.Static("/static", "static")

		// fileServer := http.FileServer(http.Dir("./static/"))
		// mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

		e.GET("/*", handler)
		e.GET("favicon.ico", func(c echo.Context) error {
			return nil
		})

		// mux.HandleFunc("/*", handler)
		// mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		// 	w.WriteHeader(http.StatusNoContent)
		// })

		e.Logger.Fatal(e.Start(":4000"))
	},
}

func handler(c echo.Context) error {
	path := utils.HandlePath(c.Request())
	lines, err := utils.ReadMarkdown(path)
	if err != nil {
		return utils.Render(c, http.StatusInternalServerError, component.Error(err))
	}

	astParser := sitter.NewParser()
	astParser.SetLanguage(markdown.GetLanguage())
	source := []byte(strings.Join(lines, "\n"))

	ctx := parser.CreateCtx()
	err = parser.LoadCtx(ctx, &lines)
	if err != nil {
		return utils.Render(c, http.StatusInternalServerError, component.Error(err))
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error while retrieving the current directory.")
	}

	rootUrl := filepath.Join(wd, "wiki")
	if err := os.MkdirAll(rootUrl, fs.ModePerm); err != nil {
		log.Fatalf("Error creating directory: %s", rootUrl)
	}

	treeRoot := filetree.GetFileTree(rootUrl, c.Request().URL.Path)

	tree, err := astParser.ParseCtx(c.Request().Context(), nil, source)
	if err != nil {
		log.Fatal("Failed to parse source code", err)
	}

	root := tree.RootNode()

	return utils.Render(c, http.StatusOK, component.Layout(root, &lines, treeRoot, ctx, c.Request().URL.Path))
}

func init() {
	rootCmd.AddCommand(startCmd)
}
