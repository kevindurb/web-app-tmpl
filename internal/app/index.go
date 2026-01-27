package app

import (
	"net/http"

	"github.com/kevindurb/web-app-tmpl/internal/html"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func (a *App) index(w http.ResponseWriter, r *http.Request) (g.Node, error) {
	return html.Layout(
		h.H1(g.Text("Hello World")),
	), nil
}
