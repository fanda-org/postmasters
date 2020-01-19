package accounting

import (
	"github.com/fanda-org/postmasters/database/models/system"
)

// Bank model
type Bank struct {
	Ledger        Ledger
	LedgerID      string `gorm:"type:char(36);not null"`
	AccountNumber string `gorm:"size:25;not null;unique_index:uix_bank_accountnumber"`
	// ShortName     *string         `gorm:"size:10"`
	// BankName      string          `gorm:"size:50"`
	AccountType string          `gorm:"size:15;not null"` // Savings, Current, FixedDeposits, etc.,
	IFSCCode    *string         `gorm:"size:15"`
	MICRCode    *string         `gorm:"size:15"`
	BranchCode  *string         `gorm:"size:15"`
	BranchName  *string         `gorm:"size:50"`
	Address     *system.Address `gorm:"foreignkey:AddrID;association_foreignkey:ID"`
	AddrID      *string         `gorm:"type:char(36);"`
	Contact     *system.Contact //`gorm:"foreignkey:ContactID;association_foreignkey:ID"`
	ContactID   *string         `gorm:"type:char(36);"`
	IsDefault   bool
}
