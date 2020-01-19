package services

import (
	"errors"

	"github.com/fanda-org/postmasters/config"
	"github.com/fanda-org/postmasters/database/models"
	"github.com/fanda-org/postmasters/database/models/system"
	"github.com/jinzhu/gorm"
)

// UsersService struct
type UsersService struct {
	dbConfig *config.DBConfig
	db       *gorm.DB
}

// GetAllUsers from database
func (u *UsersService) GetAllUsers() *[]system.User {
	//db := database.Open(u.dbConfig)
	//defer db.Close()

	users := []system.User{}
	u.db.Find(&users)
	return &users
}

// CreateUser creates user
func (u *UsersService) CreateUser(user *system.User) error {
	//db := database.Open(u.dbConfig)
	//defer db.Close()

	if err := u.db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

// GetUser gets a user from database
func (u *UsersService) GetUser(id string) (*system.User, error) {
	user, err := u.getUserOr404(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser updates user details to the database
func (u *UsersService) UpdateUser(id string, user *system.User) error {
	//db := database.Open(u.dbConfig)
	//defer db.Close()

	dbuser, err := u.getUserOr404(id)
	if err != nil {
		return err
	}
	if dbuser == nil {
		return errors.New("User not found")
	}
	if err := u.db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

// DeleteUser by id
func (u *UsersService) DeleteUser(id string) error {
	//db := database.Open(u.dbConfig)
	//defer db.Close()

	user, err := u.getUserOr404(id)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("User not found")
	}
	if err := u.db.Delete(&user).Error; err != nil {
		return err
	}
	return nil
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
func (u *UsersService) getUserOr404(id string) (*system.User, error) {
	//db := database.Open(u.dbConfig)
	//defer db.Close()

	user := system.User{}
	// system.User{UserName: name}
	filter := system.User{Base: models.Base{ID: id}}
	if err := u.db.First(&user, filter).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
