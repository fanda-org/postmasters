package system

import "github.com/fanda-org/postmasters/database/models"

// Address model
type Address struct {
	models.Base
	Attention  *string `gorm:"size:50"`
	AddrLine1  *string `gorm:"size:100"`
	AddrLine2  *string `gorm:"size:100"`
	City       *string `gorm:"size:25;index:ix_address_city"`
	State      *string `gorm:"size:25"` // Alternative for State => sql keyword
	PostalCode *string `gorm:"size:15"`
	Country    *string `gorm:"size:25"`
	Phone      *string `gorm:"size:25;index:ix_address_phone"`
	Fax        *string `gorm:"size:25"`
	AddrType   string  `gorm:"size:5;not null"` //BILL-Billing, SHIP-Shipping, REMT-Remittance
}
