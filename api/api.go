package api

import (
	"github.com/fanda-org/postmasters/config"
	"github.com/fanda-org/postmasters/router"
	"github.com/fanda-org/postmasters/services"
	"github.com/gorilla/mux"
)

// NewUsersAPI constructor for UserApi
func NewUsersAPI(dbConfig *config.DBConfig) *UsersAPI {
	service := services.NewUsersService(dbConfig)
	API := UsersAPI{service: service}
	return &API
}

// App has router and db instances
type App struct {
	Router   *mux.Router
	Env      *config.Env
	UsersAPI *UsersAPI
}

// Initialize with predefined configuration (constructor)
func (a *App) Initialize(env *config.Env, router *mux.Router) {
	a.Router = router
	a.Env = env
	a.UsersAPI = NewUsersAPI(env.Config.DB)
	a.setRoutes()
}

// setRoutes sets all required routers
func (a *App) setRoutes() {
	a.setUsersRoutes()
}

// SetUsersRoutes method
func (a *App) setUsersRoutes() {
	// Routing for handling the projects
	router.Get(a.Router, "/users", a.UsersAPI.GetAllUsers)
	router.Post(a.Router, "/users", a.UsersAPI.CreateUser)
	router.Get(a.Router, "/users/{id}", a.UsersAPI.GetUser)
	router.Put(a.Router, "/users/{id}", a.UsersAPI.UpdateUser)
	router.Delete(a.Router, "/users/{id}", a.UsersAPI.DeleteUser)
	//a.Put("/employees/{title}/disable", a.DisableEmployee)
	//a.Put("/employees/{title}/enable", a.EnableEmployee)
}
