package web

import (
	"html/template"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/fanda-org/postmasters/config"
	"github.com/fanda-org/postmasters/router"
	"github.com/fanda-org/postmasters/services"
	"github.com/gorilla/mux"
)

var tmpls = make(map[string]*template.Template)

// NewUsersController method
func newUsersController(env *config.Env) *UsersController {
	service := services.NewUsersService(env.Config.DB)
	controller := UsersController{service: service, env: env}
	return &controller
}
func newHomeController(env *config.Env) *HomeController {
	return &HomeController{env: env}
}

// Web has router and db instances
type Web struct {
	Router          *mux.Router
	Env             *config.Env
	home            *HomeController
	UsersController *UsersController
}

// Initialize with predefined configuration (constructor)
func (a *Web) Initialize(env *config.Env, router *mux.Router) {
	a.Router = router
	a.Env = env
	//a.Templates = make(map[string]*template.Template)
	a.home = newHomeController(env)
	a.UsersController = newUsersController(env)
	a.setRoutes()
}

// setRoutes sets all required routers
func (a *Web) setRoutes() {
	a.setUsersRoutes()
}

// SetUsersRoutes method
func (a *Web) setUsersRoutes() {
	router.Get(a.Router, "/", a.home.Index)
	router.Get(a.Router, "/dashboard", a.home.Dashboard)
	// Routing for handling the projects
	router.Get(a.Router, "/users", a.UsersController.GetAllUsers)
	//web.Post(a.Router, "/users", a.UsersController.CreateUser)
	router.Get(a.Router, "/users/{id}", a.UsersController.GetUser)
	//web.Put(a.Router, "/users/{id}", a.UsersController.UpdateUser)
	//web.Delete(a.Router, "/users/{id}", a.UsersController.DeleteUser)

	//a.Put("/employees/{title}/disable", a.DisableEmployee)
	//a.Put("/employees/{title}/enable", a.EnableEmployee)
}

// getTemplates gets templates with shared one
func getTemplates(name string, viewName string) (tmpl *template.Template, err error) {
	if tmpl, ok := tmpls[name]; ok {
		return tmpl, nil
	}
	var sharedFiles []string
	//for _, dir := range []string{"./views/_shared","./views/"+path} {
	dir := "./views/_shared"
	files2, _ := ioutil.ReadDir(dir)
	for _, file := range files2 {
		filename := file.Name()
		if strings.HasSuffix(filename, ".gohtml") {
			filePath := filepath.Join(dir, filename)
			sharedFiles = append(sharedFiles, filePath)
		}
	}
	allFiles := append(sharedFiles, "views\\"+viewName+".gohtml")

	tmpl, err = template.New(name).ParseFiles(allFiles...)
	//tmpls[name] = tmpl
	return
}
