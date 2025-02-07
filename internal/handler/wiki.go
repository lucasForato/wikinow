package handler

import (
	"context"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"wikinow/component"
	"wikinow/infra/config"
	"wikinow/infra/logger"
	"wikinow/internal/filetree"
	"wikinow/internal/parser"
	"wikinow/internal/utils"

	"github.com/labstack/echo/v4"
	sitter "github.com/smacker/go-tree-sitter"
	markdown "github.com/smacker/go-tree-sitter/markdown/tree-sitter-markdown"
)

func Wiki(c echo.Context) error {
	path := utils.HandlePath(c.Request())
	lines, err := utils.ReadMarkdown(path)
	if err != nil {
		return utils.Render(c, http.StatusInternalServerError, component.Error(err))
	}

	ctx := parser.CreateCtx()
	err = parser.LoadCtx(ctx, &lines)
	if err != nil {
		return utils.Render(c, http.StatusInternalServerError, component.Error(err))
	}

	astRoot, err := getAstTree(lines, c)
	if err != nil {
		return utils.Render(c, http.StatusInternalServerError, component.Error(err))
	}

	filetreeRoot, err := getFileTree(c)
	if err != nil {
		return utils.Render(c, http.StatusInternalServerError, component.Error(err))
	}

	if c.Request().Header.Get("HX-Request") == "true" {
		return utils.Render(c, http.StatusOK, component.Content(astRoot, &lines, ctx))
	}
	title, err := config.GetTitle()
	if err != nil {
		logger.Error(err)
	}

	return utils.Render(c, http.StatusOK, component.Layout(astRoot, &lines, filetreeRoot, ctx, c.Request().URL.Path, title))
}

func getAstTree(lines []string, c echo.Context) (*sitter.Node, error) {
	astParser := sitter.NewParser()
	astParser.SetLanguage(markdown.GetLanguage())
	source := []byte(strings.Join(lines, "\n"))
	astTree, err := astParser.ParseCtx(c.Request().Context(), nil, source)
	if err != nil {
		return nil, err
	}
	astRoot := astTree.RootNode()
	return astRoot, nil
}

func TestAst(lines []string) (*sitter.Node, error) {
	astParser := sitter.NewParser()
	astParser.SetLanguage(markdown.GetLanguage())
	source := []byte(strings.Join(lines, "\n"))
	astTree, err := astParser.ParseCtx(context.Background(), nil, source)
	if err != nil {
		return nil, err
	}
	astRoot := astTree.RootNode()
	return astRoot, nil
}

func getFileTree(c echo.Context) (*filetree.TreeNode, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	rootPath := filepath.Join(wd, "wiki")
	if err := os.MkdirAll(rootPath, fs.ModePerm); err != nil {
		return nil, err
	}
	filetreeRoot, _ := filetree.GetFileTree(rootPath, c.Request().URL.Path)
	return filetreeRoot, nil
}
