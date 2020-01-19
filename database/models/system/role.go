package system

import "github.com/fanda-org/postmasters/database/models"

// Role model
type Role struct {
	models.Base
	RoleName     string       `gorm:"size:50;not null;unique_index:uix_role_name"`
	Description  *string      `gorm:"size:255"`
	Organization Organization `gorm:"foreignkey:OrgID"` //;association_foreignkey:ID"`
	OrgID        *string      `gorm:"type:char(36)"`    // NULL = System Role
	// Users        []OrgUser `gorm:"many2many:user_roles"`
}
