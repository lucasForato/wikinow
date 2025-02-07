package handler

import (
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"wikinow/infra/logger"
	"wikinow/internal/filetree"
	"wikinow/internal/utils"

	"github.com/labstack/echo/v4"
)

func Filetree(c echo.Context) error {
	body := make(map[string]interface{})
	err := utils.GetRequestBody(c.Request(), &body)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	wd, err := os.Getwd()
	if err != nil {
		logger.Error("Error while retrieving the current directory.")
	}

	rootUrl := filepath.Join(wd, "wiki")
	if err := os.MkdirAll(rootUrl, fs.ModePerm); err != nil {
		logger.Error(fmt.Sprintf("Error creating directory: %s", rootUrl))
	}

	treeRoot, _ := filetree.GetFileTree(rootUrl, body["path"].(string))
	return c.JSON(http.StatusOK, treeRoot)
}
