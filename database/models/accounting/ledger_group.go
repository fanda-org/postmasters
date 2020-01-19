package accounting

import (
	"github.com/fanda-org/postmasters/database/models"
	"github.com/fanda-org/postmasters/database/models/system"
)

// LedgerGroup model
type LedgerGroup struct {
	models.Base
	GroupCode    string              `gorm:"size:12;not null;unique_index:uix_ledger_group_code"` //A-00-0-00000 -> (A/L/I/E)-SEQUENCE-LEVEL-INDEX
	GroupName    string              `gorm:"size:50;not null;unique_index:uix_ledger_group_name"`
	Description  *string             `gorm:"size:255"`
	GroupType    string              `gorm:"size:5;not null"` // AST-Asset, LIA-Liability, INC-Income, EXP-Expenses
	Parent       *LedgerGroup        // `gorm:"foreignkey:ParentID;association_foreignkey:ID"`
	ParentID     string              `gorm:"type:char(36);not null"`
	Organization system.Organization `gorm:"foreignkey:OrgID;association_foreignkey:ID"`
	OrgID        *string             `gorm:"type:char(36);unique_index:uix_ledger_group_code,uix_ledger_group_name"` // NULL = System LedgerGroup
	IsSystem     bool
}
