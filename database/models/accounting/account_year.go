package accounting

import (
	"github.com/fanda-org/postmasters/database/models"
	"github.com/fanda-org/postmasters/database/models/system"
	"time"
)

//AccountYear model
type AccountYear struct {
	models.Base
	YearCode     string `gorm:"size:15;not null;unique_index:uix_year_code"`
	YearBegin    time.Time
	YearEnd      time.Time
	Organization system.Organization `gorm:"foreignkey:OrgID;association_foreignkey:ID"`
	OrgID        string              `gorm:"type:char(36);not null;unique_index:uix_year_code"`
}
