package system

import "github.com/fanda-org/postmasters/database/models"

// Contact model
type Contact struct {
	models.Base
	Salutation  *string `gorm:"size:5"`
	FirstName   *string `gorm:"size:50;index:ix_contact_firstname"`
	LastName    *string `gorm:"size:50"`
	Email       *string `gorm:"size:100;index:ix_contact_email"`
	WorkPhone   *string `gorm:"size:25"`
	Mobile      *string `gorm:"size:25;index:ix_contact_mobile"`
	Designation *string `gorm:"size:25"`
	Department  *string `gorm:"size:25"`
	IsPrimary   bool    `gorm:"not null"`
}
