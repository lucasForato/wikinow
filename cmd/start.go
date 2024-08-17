package cmd

import (
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"wikinow/component"
	"wikinow/internal/parser"
	"wikinow/internal/utils"

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
		mux := http.NewServeMux()
		mux.HandleFunc("/*", handler)
		mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
		})

		log.Info("Starting server at port 4000")
		log.Fatal(http.ListenAndServe(":4000", mux))
	},
}

func handler(w http.ResponseWriter, r *http.Request) {
	path := utils.HandlePath(r)
	lines, err := utils.ReadMarkdown(path)
	if err != nil {
		utils.Render(w, r, http.StatusInternalServerError, component.Error(err))
		return
	}

	astParser := sitter.NewParser()
	astParser.SetLanguage(markdown.GetLanguage())
	source := []byte(strings.Join(lines, "\n"))

	ctx := parser.CreateCtx()
	err = parser.LoadCtx(ctx, &lines)
	if err != nil {
		utils.Render(w, r, http.StatusInternalServerError, component.Error(err))
		return
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error while retrieving the current directory.")
	}

	rootUrl := filepath.Join(wd, "wiki")
	if err := os.MkdirAll(rootUrl, fs.ModePerm); err != nil {
		log.Fatalf("Error creating directory: %s", rootUrl)
	}

	treeRoot := utils.GetFileTree(rootUrl, r.URL.Path)

	tree, err := astParser.ParseCtx(r.Context(), nil, source)
	if err != nil {
		log.Fatal("Failed to parse source code", err)
	}

	root := tree.RootNode()

	utils.Render(w, r, http.StatusOK, component.Layout(root, &lines, treeRoot, ctx))
}

func init() {
	rootCmd.AddCommand(startCmd)
}
