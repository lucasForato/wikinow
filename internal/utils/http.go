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
