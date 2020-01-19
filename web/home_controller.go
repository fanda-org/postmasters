package web

import (
	"net/http"

	"github.com/fanda-org/postmasters/config"
)

// HomeController struct
type HomeController struct {
	env *config.Env
}

// Index from database
func (h *HomeController) Index(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := getTemplates("home_index", "home\\index")
	err := tmpl.ExecuteTemplate(w, "home_index", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Dashboard from database
func (h *HomeController) Dashboard(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := getTemplates("home_dashboard", "home\\dashboard")
	err := tmpl.ExecuteTemplate(w, "home_dashboard", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
