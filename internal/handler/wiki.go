package handler

import (
	"fmt"
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
)

func Wiki(c echo.Context) error {
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
  fmt.Println(rootUrl)
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
