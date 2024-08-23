package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

func GetRequestBody[T any](req *http.Request, out *T) error {
	if req.Body == nil {
		return errors.New("request body is empty")
	}

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(req.Body); err != nil {
		return fmt.Errorf("error reading request body: %w", err)
	}

	if buf.Len() == 0 {
		return errors.New("request body is empty")
	}

	req.Body = io.NopCloser(buf)

	if err := json.NewDecoder(req.Body).Decode(out); err != nil {
		return fmt.Errorf("error parsing request body: %w", err)
	}

	return nil
}

func GetResponseBody[T any](res *http.Response, out *T) error {
	if res.Body == nil {
		return errors.New("request body is empty")
	}

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(res.Body); err != nil {
		return fmt.Errorf("error reading request body: %w", err)
	}

	if buf.Len() == 0 {
		return errors.New("request body is empty")
	}

	res.Body = io.NopCloser(buf)

	if err := json.NewDecoder(res.Body).Decode(out); err != nil {
		return fmt.Errorf("error parsing request body: %w", err)
	}

	return nil
}
