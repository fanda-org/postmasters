package accounting

import (
	"github.com/fanda-org/postmasters/database/models/system"
)

// Party model
type Party struct {
	Ledger            Ledger
	LedgerID          string              `gorm:"type:char(36);not null"`
	PartyOrganization system.Organization `gorm:"foreignkey:PartyOrgID"`
	PartyOrgID        string              `gorm:"type:char(36);not null"`
	PartyType         string              `gorm:"size:5;not null"`  // CUST-Customer, SUPP-Suplier, VEND-Vendor, BUYR-Buyer, EMPL-Employee
	PaymentTerm       string              `gorm:"size:10;not null"` // IMM-Immediate, NET7, NET15, NET30, NET45, NET60, ONDATE
	CreditLimit       float32
}
