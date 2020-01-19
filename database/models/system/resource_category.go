package system

import "github.com/fanda-org/postmasters/database/models"

// ResourceCategory model
type ResourceCategory struct {
	models.Base
	CategoryName string  `gorm:"size:50;not null;unique_index:uix_resource_category_name"`
	Description  *string `gorm:"size:255"`
}
