package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"wikinow/internal/filetree"
	"wikinow/internal/utils"

	"github.com/labstack/echo/v4"
)

func GetFiletree(c echo.Context) (*filetree.TreeNode, error) {
	body, err := json.Marshal(map[string]string{
		"path": c.Request().URL.Path,
	})

	payload := bytes.NewBuffer(body)

	if err != nil {
		return nil, errors.New("Error marshaling request body")
	}

	scheme := "http"
	if c.Request().TLS != nil {
		scheme = "https"
	}

	url := fmt.Sprintf("%s://%s/api/filetree", scheme, c.Request().Host)
	req, err := http.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		return nil, errors.New("Error creating request")
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	node := &filetree.TreeNode{}
	if err := utils.GetResponseBody(res, node); err != nil {
		return nil, err
	}

	return node, nil
}
