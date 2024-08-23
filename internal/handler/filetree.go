package handler

import (
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"wikinow/internal/filetree"
	"wikinow/internal/utils"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func Filetree(c echo.Context) error {
	body := make(map[string]interface{})
	err := utils.GetRequestBody(c.Request(), &body)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error while retrieving the current directory.")
	}

	rootUrl := filepath.Join(wd, "wiki")
	if err := os.MkdirAll(rootUrl, fs.ModePerm); err != nil {
		log.Fatalf("Error creating directory: %s", rootUrl)
	}

	treeRoot, _ := filetree.GetFileTree(rootUrl, body["path"].(string))
	return c.JSON(http.StatusOK, treeRoot)
}
