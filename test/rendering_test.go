package test

import (
	"github.com/MuhamadAndre10/smartStudentApp/internal/config"
	"github.com/MuhamadAndre10/smartStudentApp/web"
	"github.com/stretchr/testify/assert"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRenderingTemplate(t *testing.T) {

	log := config.NewLog()
	cfg := config.NewViper()
	cfg.Set("app.env", "development")

	render := web.NewRender(log, cfg)
	render.TemplateCache = make(map[string]*template.Template)

	request := httptest.NewRequest("GET", "/login", nil)

	w := httptest.NewRecorder()

	render.RenderTemplate(w, request, "login", nil)

	assert.Equal(t, http.StatusOK, w.Code)

}
