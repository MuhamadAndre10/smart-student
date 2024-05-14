package config

import (
	"github.com/MuhamadAndre10/smartStudentApp/internal/delivery/http"
	"github.com/MuhamadAndre10/smartStudentApp/internal/delivery/http/route"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB        *gorm.DB
	Mux       *mux.Router
	Log       *logrus.Logger
	Config    *viper.Viper
	Validator *validator.Validate
}

func Bootstrap(config *BootstrapConfig) {

	// setup controller
	adminController := http.NewAdminController(config.Log, config.Config)
	authController := http.NewAuthController(config.Log, config.Config)
	studentController := http.NewManagerStudentController(config.Log, config.Config)

	routeConfig := route.ConfigRoute{
		Mux: config.Mux,
		Log: config.Log,
		Cfg: config.Config,

		DashboardController: adminController,
		AuthController:      authController,
		StudentController:   studentController,
	}

	routeConfig.Setup()
}
