package handlers

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
)

type HandlerFunc func(http.ResponseWriter, *http.Request) error

func Make(h HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("HTTP handler error", "err", err, "path", r.URL.Path)
		}
	}
}

func Render(w http.ResponseWriter, r *http.Request, c templ.Component) error {
	return c.Render(r.Context(), w)
}
