package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(c echo.Context, status int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(c.Request().Context(), buf); err != nil {
		return c.String(status, err.Error())
	}

	return c.HTMLBlob(http.StatusOK, buf.Bytes())
}

func HandlePath(r *http.Request) string {
	url := r.URL.Path

	if url[len(url)-1] == '/' {
		url += "main"
	}

	return strings.Join([]string{url[1:], "md"}, ".")
}

func GetJSONBody(c echo.Context) (map[string]interface{}, error) {
	jsonBody := make(map[string]interface{})

	if err := json.NewDecoder(c.Request().Body).Decode(&jsonBody); err != nil {
		return jsonBody, errors.New("Error parsing request body.")
	}

	return jsonBody, nil
}
