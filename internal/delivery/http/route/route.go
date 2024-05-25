package route

import (
	deliveryhttp "github.com/MuhamadAndre10/smartStudentApp/internal/delivery/http"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ConfigRoute struct {
	Mux *mux.Router
	Log *logrus.Logger
	Cfg *viper.Viper

	DashboardController *deliveryhttp.DashboardController
	AuthController      *deliveryhttp.AuthController
	StudentController   *deliveryhttp.StudentController
}

func (cr *ConfigRoute) Setup() {
	cr.fileServer()
	cr.notFound()

	// Route
	cr.setupAdminPanelRoute()
	cr.setupGuestRoute()
}

func (cr *ConfigRoute) setupGuestRoute() {
	cr.Mux.HandleFunc("/login", cr.AuthController.ShowLogin).Methods("GET")
	cr.Mux.HandleFunc("/register", cr.AuthController.ShowRegister).Methods("GET")

}

func (cr *ConfigRoute) setupAdminPanelRoute() {
	cr.Mux.HandleFunc("/dashboard", cr.DashboardController.Dashboard).Methods("GET")
	cr.Mux.HandleFunc("/students", cr.StudentController.ShowData).Methods("GET")
	cr.Mux.HandleFunc("/form-add", cr.StudentController.ShowFormAddData).Methods("GET")

	cr.Mux.HandleFunc("/students", cr.StudentController.InsertData).Methods("POST")
}
