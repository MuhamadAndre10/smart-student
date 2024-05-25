package http

import (
	"github.com/MuhamadAndre10/smartStudentApp/web"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
)

type StudentController struct {
	Log    *logrus.Logger
	Config *viper.Viper
}

func NewManagerStudentController(log *logrus.Logger, Cfg *viper.Viper) *StudentController {
	return &StudentController{
		log,
		Cfg,
	}
}

func (s *StudentController) ShowData(w http.ResponseWriter, r *http.Request) {
	web.NewRender(s.Log, s.Config).RenderTemplate(w, r, "list_student", &web.TemplateData{
		Active: r.URL.Path,
	})
}

func (s *StudentController) ShowFormAddData(w http.ResponseWriter, r *http.Request) {
	web.NewRender(s.Log, s.Config).RenderTemplate(w, r, "form_add_student", &web.TemplateData{
		Active: r.URL.Path,
	})
}

func (s *StudentController) InsertData(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}
