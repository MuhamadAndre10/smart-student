package route

import (
	"github.com/MuhamadAndre10/smartStudentApp/web"
	"net/http"
)

func (cr *ConfigRoute) notFound() {
	// 404
	cr.Mux.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "text/html")

		web.NewRender(cr.Log, cr.Cfg).RenderTemplate(w, r, "404", nil)
	})
}
