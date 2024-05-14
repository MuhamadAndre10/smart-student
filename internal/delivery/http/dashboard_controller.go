package http

import (
	"github.com/MuhamadAndre10/smartStudentApp/web"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
)

type DashboardController struct {
	Log    *logrus.Logger
	Config *viper.Viper
}

func NewAdminController(log *logrus.Logger, Cfg *viper.Viper) *DashboardController {
	return &DashboardController{
		log,
		Cfg,
	}
}

func (ac *DashboardController) Dashboard(w http.ResponseWriter, r *http.Request) {

	web.NewRender(ac.Log, ac.Config).RenderTemplate(w, r, "dashboard", &web.TemplateData{
		Active: r.URL.Path,
	})
}
