package utils

import (
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
