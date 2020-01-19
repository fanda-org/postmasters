package config

import (
	"html/template"
	"log"

	"github.com/gorilla/mux"
)

// Env struct
type Env struct {
	Config *Config
	//DB          *gorm.DB
	Logger     *log.Logger
	APIRouter  *mux.Router
	WebRouter  *mux.Router
	Views      *template.Template
	UsersViews *template.Template
}
