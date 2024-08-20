package utils

import (
	"net/http"
	"path"
	"strings"

	"github.com/a-h/templ"
)

func Render(w http.ResponseWriter, r *http.Request, statusCode int, t templ.Component) {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(r.Context(), buf); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// simpleTemplate := `<html><head><link href="./static/css/style.css" rel="stylesheet"></head><body><h1 class="bg-red-500">Test</h1></body></html>`
	// w.Header().Set("Content-Type", "text/html")
	// w.Write([]byte(simpleTemplate))

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(statusCode)
	w.Write([]byte(buf.String()))
}

func HandlePath(r *http.Request) string {
	url := r.URL.Path

	if url[len(url)-1] == '/' {
		url += "main"
	}

	file := strings.Join([]string{url, "md"}, ".")
	return path.Join("wiki", file)
}
