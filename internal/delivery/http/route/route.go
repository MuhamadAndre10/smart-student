package route

import (
	deliveryhttp "github.com/MuhamadAndre10/smartStudentApp/internal/delivery/http"
	"github.com/MuhamadAndre10/smartStudentApp/web"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/fs"
	"net/http"

	"github.com/gorilla/mux"
)

type ConfigRoute struct {
	Mux *mux.Router
	Log *logrus.Logger
	Cfg *viper.Viper

	DashboardController *deliveryhttp.DashboardController
	AuthController      *deliveryhttp.AuthController
	StudentController   *deliveryhttp.ManagerStudentController
}

func (cr *ConfigRoute) Setup() {

	cr.setupFileServerRoute()
	cr.setupNotFoundHandler()

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
	cr.Mux.HandleFunc("/students", cr.StudentController.ListDataStudent).Methods("GET")
}

func (cr *ConfigRoute) setupNotFoundHandler() {
	// 404
	cr.Mux.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "text/html")

		web.NewRender(cr.Log, cr.Cfg).RenderTemplate(w, r, "404", nil)
	})
}

func (cr *ConfigRoute) setupFileServerRoute() {
	// File server
	dir, _ := fs.Sub(web.StaticFS, "static")
	fileServer := http.FileServer(http.FS(dir))
	cr.Mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))
}
