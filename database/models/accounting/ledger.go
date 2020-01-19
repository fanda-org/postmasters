package accounting

import (
	"github.com/fanda-org/postmasters/database/models"
	"github.com/fanda-org/postmasters/database/models/system"
)

// Ledger model
type Ledger struct {
	models.Base
	LedgerCode  string  `gorm:"size:15;not null;unique_index:uix_ledger_code"`
	LedgerName  string  `gorm:"size:50;not null;unique_index:uix_ledger_name"`
	Description *string `gorm:"size:255"`
	//LedgerOrg   system.Organization //`gorm:"foreignkey:LedgerOrgID;association_foreignkey:ID"`
	//LedgerOrgID   *string             `gorm:"type:char(36)"`
	LedgerGroup   LedgerGroup
	LedgerGroupID string `gorm:"type:char(36);not null"`
	//LedgerType    string              `gorm:"size:5"` // CUST-Customer, SUPP-Supplier, RETN-Return, BANK-Bank
	Parent       *Ledger             // `gorm:"foreignkey:ParentID;association_foreignkey:ID"`
	ParentID     string              `gorm:"type:char(36);not null"`
	Organization system.Organization `gorm:"foreignkey:OrgID;association_foreignkey:ID"`
	OrgID        *string             `gorm:"type:char(36);unique_index:uix_ledger_code,uix_ledger_name"` // NULL = System Ledger
	IsSystem     bool
}
