package web

import (
	"embed"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"html/template"
	"net/http"
)

type TemplateData struct {
	StringMap       map[string]string
	IntMap          map[string]int
	FloatMap        map[string]float32
	Data            map[string]interface{}
	Active          string
	CSRFToken       string
	Flash           string
	Warning         string
	Error           string
	IsAuthenticated int
	API             string
	CSSVersion      string
}

type ConfigRender struct {
	Log           *logrus.Logger
	Config        *viper.Viper
	TemplateCache map[string]*template.Template
}

//go:embed templates/*
var templateFS embed.FS

//go:embed static/*
var StaticFS embed.FS

var functionMap = template.FuncMap{}

func (cr *ConfigRender) DefaultData(td *TemplateData, r *http.Request) *TemplateData {
	return td
}

func NewRender(log *logrus.Logger, config *viper.Viper) *ConfigRender {
	tc := make(map[string]*template.Template)
	return &ConfigRender{
		Log:           log,
		Config:        config,
		TemplateCache: tc,
	}
}

func (cr *ConfigRender) RenderTemplate(w http.ResponseWriter, r *http.Request, page string, td *TemplateData, partials ...string) {
	var f *template.Template

	templateToRender := fmt.Sprintf("templates/%s.gohtml", page)

	_, ok := cr.TemplateCache[templateToRender]

	if cr.Config.GetString("app.env") == "production" && ok {
		f = cr.TemplateCache[templateToRender]
	} else {
		// parse template
		partials = append(partials, "sidebar", "navbar")
		f = cr.ParseTemplate(partials, page, templateToRender)
	}

	// if template data nil, create new template data
	if td == nil {
		td = &TemplateData{}
	}

	// set default data
	td = cr.DefaultData(td, r)

	err := f.Execute(w, td)
	if err != nil {
		cr.Log.Errorf("Error executing template: %v", err)
		return
	}

}

// ParseTemplate parses the template and returns the template
func (cr *ConfigRender) ParseTemplate(partials []string, page, templateToRender string) *template.Template {
	var f *template.Template
	var err error

	if len(partials) > 0 {
		for i, partial := range partials {
			partials[i] = fmt.Sprintf("templates/partials/%s.gohtml", partial)
		}
	}

	if len(partials) > 0 {
		files := append([]string{"templates/layouts/base.gohtml", templateToRender}, partials...)

		f, err = template.New(fmt.Sprintf("%s.gohtml", page)).Funcs(functionMap).ParseFS(templateFS, files...)

	} else {
		// Parse the template without partials
		f, err = template.New(fmt.Sprintf("%s.gohtml", page)).Funcs(functionMap).ParseFS(templateFS, "templates/layouts/base.gohtml", templateToRender)
	}

	if err != nil {
		cr.Log.Errorf("Error parsing template: %v", err)
		return nil
	}

	cr.TemplateCache[templateToRender] = f
	return f

}
