package api

import (
	"encoding/json"
	"net/http"

	"github.com/fanda-org/postmasters/api/handler"
	"github.com/fanda-org/postmasters/database/models/system"
	"github.com/fanda-org/postmasters/services"
	"github.com/gorilla/mux"
)

// UsersAPI struct
type UsersAPI struct {
	service *services.UsersService
}

// GetAllUsers from database
func (u *UsersAPI) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	handler.RespondJSON(w, http.StatusOK, u.service.GetAllUsers())
}

// CreateUser creates user
func (u *UsersAPI) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := system.User{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		handler.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := u.service.CreateUser(&user); err != nil {
		handler.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	handler.RespondJSON(w, http.StatusCreated, user)
}

// GetUser gets a user from database
func (u *UsersAPI) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	user, err := u.service.GetUser(id)
	if err != nil {
		handler.RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	handler.RespondJSON(w, http.StatusOK, user)
}

// UpdateUser updates user details to the database
func (u *UsersAPI) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	user, err := u.service.GetUser(id)
	if err != nil {
		handler.RespondError(w, http.StatusNotFound, err.Error())
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		handler.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := u.service.UpdateUser(id, user); err != nil {
		handler.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	handler.RespondJSON(w, http.StatusOK, user)
}

// DeleteUser by id
func (u *UsersAPI) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	err := u.service.DeleteUser(id)
	if err != nil {
		handler.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	handler.RespondJSON(w, http.StatusNoContent, nil)
}

// func DisableEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	name := vars["name"]
// 	employee := getEmployeeOr404(db, name, w, r)
// 	if employee == nil {
// 		return
// 	}
// 	employee.Disable()
// 	if err := db.Save(&employee).Error; err != nil {
// 		respondError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondJSON(w, http.StatusOK, employee)
// }

// func EnableEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	name := vars["name"]
// 	employee := getEmployeeOr404(db, name, w, r)
// 	if employee == nil {
// 		return
// 	}
// 	employee.Enable()
// 	if err := db.Save(&employee).Error; err != nil {
// 		respondError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondJSON(w, http.StatusOK, employee)
// }

// getUserOr404 gets a employee instance if exists, or respond the 404 error otherwise
// func (u *UsersAPI) getUserOr404(id string, w http.ResponseWriter, r *http.Request) *system.User {
// 	user := &system.User{}

// 	// system.User{UserName: name}
// 	filter := &system.User{Base: models.Base{ID: id}}
// 	if err := DB.First(user, filter).Error; err != nil {
// 		handler.RespondError(w, http.StatusNotFound, err.Error())
// 		return nil
// 	}
// 	return user
// }
