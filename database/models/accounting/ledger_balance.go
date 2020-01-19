package accounting

// LedgerBalance model
type LedgerBalance struct {
	Ledger         Ledger      //`gorm:"foreignkey:LedgerID;association_foreignkey:ID"`
	LedgerID       string      `gorm:"type:char(36);primary_key"`
	AccountYear    AccountYear `gorm:"foreignkey:YearID;association_foreignkey:ID"`
	YearID         string      `gorm:"type:char(36);primary_key"`
	OpeningBalance float32
	BalanceSign    string `gorm:"type:char"`
}
