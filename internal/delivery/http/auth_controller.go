package http

import (
	"github.com/MuhamadAndre10/smartStudentApp/web"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
)

type AuthController struct {
	Log    *logrus.Logger
	Config *viper.Viper
}

func NewAuthController(log *logrus.Logger, Cfg *viper.Viper) *AuthController {
	return &AuthController{
		log,
		Cfg,
	}
}

func (ac *AuthController) ShowLogin(w http.ResponseWriter, r *http.Request) {
	web.NewRender(ac.Log, ac.Config).RenderTemplate(w, r, "login", nil)
}

func (ac *AuthController) ShowRegister(w http.ResponseWriter, r *http.Request) {
	web.NewRender(ac.Log, ac.Config).RenderTemplate(w, r, "register", nil)
}
