package system

import "github.com/fanda-org/postmasters/database/models"

// Resource model
type Resource struct {
	models.Base
	ResourceName       string           `gorm:"size:50;not null;unique_index:uix_resource_name"`
	Description        *string          `gorm:"size:255"`
	ResourceCategory   ResourceCategory //`gorm:"foreignkey:ResCatID;association_foreignkey:ID"`
	ResourceCategoryID string           `gorm:"type:char(36);not null"`
	DisplayName        *string          `gorm:"size:25"`
}
