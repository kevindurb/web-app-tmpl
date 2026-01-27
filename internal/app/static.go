package app

import (
	"io/fs"
	"net/http"

	"github.com/kevindurb/web-app-tmpl/web"
)

func (a *App) Static() http.Handler {
	staticFs, _ := fs.Sub(web.Files, "static")
	return http.FileServer(http.FS(staticFs))
}
