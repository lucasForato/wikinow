package handler

import (
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"wikinow/component"
	"wikinow/internal/filetree"
	"wikinow/internal/types"
	"wikinow/internal/utils"

	"github.com/labstack/echo/v4"
)

func GETSearch(c echo.Context) error {
	return utils.Render(c, http.StatusOK, component.SearchModal())
}

func POSTSearch(c echo.Context) error {
	q := utils.Normalize(c.FormValue("q"))

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	rootPath := filepath.Join(wd, "wiki")
	if err := os.MkdirAll(rootPath, fs.ModePerm); err != nil {
		return err
	}

	output := []types.SearchResult{}

	err = filepath.WalkDir(rootPath, func(path string, entry fs.DirEntry, err error) error {
    if entry.IsDir() {
      return nil
    }

		exists, err := utils.FindInFile(path, q)

		if exists {
			title, err := utils.GetFileTitle(path)
      if err != nil {
        return err
      }


			result := &types.SearchResult{
				Title: title,
				Path:  filetree.NormalizePath(path, rootPath),
			}
			output = append(output, *result)
		}

		return nil
	})
	if err != nil {
		return err
	}

	return utils.Render(c, http.StatusOK, component.SearchResults(output))
}
