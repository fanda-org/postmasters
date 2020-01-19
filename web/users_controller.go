package web

import (
	"net/http"

	"github.com/fanda-org/postmasters/config"
	"github.com/fanda-org/postmasters/services"
	"github.com/gorilla/mux"
)

// UsersController struct
type UsersController struct {
	service *services.UsersService
	env     *config.Env
}

// GetAllUsers from database
func (u *UsersController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := u.service.GetAllUsers()
	tmpl, _ := getTemplates("users_index", "users\\index")
	err := tmpl.ExecuteTemplate(w, "users_index", users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	//u.env.UsersViews.Lookup("layout.hbs")
	//tmpl.Execute(w, users)
}

// GetUser from database
func (u *UsersController) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	user, _ := u.service.GetUser(id)
	tmpl, _ := getTemplates("users_view", "users\\view")
	err := tmpl.ExecuteTemplate(w, "users_view", user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	//u.env.UsersViews.Lookup("layout.hbs")
	//tmpl.Execute(w, users)
}
