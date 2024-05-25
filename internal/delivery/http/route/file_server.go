package route

import (
	"github.com/MuhamadAndre10/smartStudentApp/web"
	"io/fs"
	"net/http"
)

func (cr *ConfigRoute) fileServer() {
	// File server
	dir, _ := fs.Sub(web.StaticFS, "static")
	fileServer := http.FileServer(http.FS(dir))
	cr.Mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))
}
