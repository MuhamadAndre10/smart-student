package http

import (
	"github.com/MuhamadAndre10/smartStudentApp/web"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
)

type ManagerStudentController struct {
	Log    *logrus.Logger
	Config *viper.Viper
}

func NewManagerStudentController(log *logrus.Logger, Cfg *viper.Viper) *ManagerStudentController {
	return &ManagerStudentController{
		log,
		Cfg,
	}
}

func (s *ManagerStudentController) ListDataStudent(w http.ResponseWriter, r *http.Request) {
	web.NewRender(s.Log, s.Config).RenderTemplate(w, r, "list_student", &web.TemplateData{
		Active: r.URL.Path,
	})
}
