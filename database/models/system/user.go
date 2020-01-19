package system

import (
	"errors"
	"time"

	"github.com/fanda-org/postmasters/database/models"
	"golang.org/x/crypto/bcrypt"
	// "github.com/jinzhu/gorm"
	// "github.com/satori/go.uuid"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	// "database/sql"
)

// User model
type User struct {
	models.Base
	UserName      string         `gorm:"size:50;not null;unique_index:uix_user_name" json:"userName"`
	Email         string         `gorm:"size:100;not null;unique_index:uix_user_email" json:"email"`
	PasswordHash  string         `gorm:"size:255;not null" json:"-"`
	Salutation    *string        `gorm:"size:5" json:"salutation,omitempty"`
	FirstName     *string        `gorm:"size:50;index:ix_user_firstname" json:"firstName,omitempty"`
	LastName      *string        `gorm:"size:50" json:"lastName,omitempty"`
	WorkPhone     *string        `gorm:"size:25" json:"workPhone,omitempty"`
	Mobile        *string        `gorm:"size:25;index:ix_user_mobile" json:"mobile,omitempty"`
	LoginAt       *time.Time     `json:"loginAt,omitempty"`
	Organizations []Organization `gorm:"many2many:org_users;foreignkey:ID;association_foreignkey:ID;jointable_foreignkey:user_id;association_jointable_foreignkey:org_id;" json:"-"`
}

// Roles []Role `gorm:"many2many:user_roles"`

// SetPassword - What's bcrypt? https://en.wikipedia.org/wiki/Bcrypt
// Golang bcrypt doc: https://godoc.org/golang.org/x/crypto/bcrypt
// You can change the value in bcrypt.DefaultCost to adjust the security index.
// 	err := User.setPassword("password0")
func (u *User) SetPassword(password string) error {
	if len(password) == 0 {
		return errors.New("Password should not be empty")
	}
	bytePassword := []byte(password)
	// Make sure the second param `bcrypt generator cost` between [4, 32)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.PasswordHash = string(passwordHash)
	return nil
}

// VerifyPassword - Database will only save the hashed string, you should check it by util function.
// if err := serModel.checkPassword("password0"); err != nil { password error }
func (u *User) VerifyPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.PasswordHash)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}
