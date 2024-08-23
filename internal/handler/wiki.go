package handler

import (
	"net/http"
	"strings"
	"wikinow/component"
	"wikinow/internal/parser"
	"wikinow/internal/service"
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

	ctx := parser.CreateCtx()
	err = parser.LoadCtx(ctx, &lines)
	if err != nil {
		return utils.Render(c, http.StatusInternalServerError, component.Error(err))
	}

	astParser := sitter.NewParser()
	astParser.SetLanguage(markdown.GetLanguage())
	source := []byte(strings.Join(lines, "\n"))
	astTree, err := astParser.ParseCtx(c.Request().Context(), nil, source)
	if err != nil {
		log.Fatal("Failed to parse source code", err)
	}
	root := astTree.RootNode()

	fileTree, err := service.GetFiletree(c)
	if err != nil {
		return utils.Render(c, http.StatusInternalServerError, component.Error(err))
	}

	return utils.Render(c, http.StatusOK, component.Layout(root, &lines, fileTree, ctx, c.Request().URL.Path))
}
