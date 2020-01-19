package system

import "github.com/fanda-org/postmasters/database/models"

// Organization model
type Organization struct {
	models.Base
	OrgCode     string     `gorm:"size:15;not null;unique_index:uix_org_code"`
	OrgName     string     `gorm:"size:50;not null;unique_index:uix_org_name"`
	Description *string    `gorm:"size:255"`
	RegdNumber  *string    `gorm:"size:25"`
	PAN         *string    `gorm:"size:25"`
	TAN         *string    `gorm:"size:25"`
	GSTIN       *string    `gorm:"size:25"`
	Addresses   []*Address `gorm:"many2many:org_addresses;foreignkey:ID;association_foreignkey:ID;jointable_foreignkey:org_id;association_jointable_foreignkey:addr_id;"`
	Contacts    []*Contact `gorm:"many2many:org_contacts;foreignkey:ID;association_foreignkey:ID;jointable_foreignkey:org_id;association_jointable_foreignkey:contact_id;"`
	Users       []*User    `gorm:"many2many:org_users;foreignkey:ID;association_foreignkey:ID;jointable_foreignkey:org_id;association_jointable_foreignkey:user_id;"`
}

//OrgAddresses 	[]OrgAddress 	`gorm:"foreignkey:OrgID;association_foreignkey:ID"`
//Addresses   	[]Address 		`gorm:"many2many:org_addresses"`
//Contacts      []Contact 		`gorm:"many2many:org_contacts"`

// BillingAddress      Address `gorm:"foreignkey:BillingAddressID;association_foreignkey:ID"`
// BillingAddressID    *string `gorm:"type:char(36);"`
// ShippingAddress     Address // `gorm:"foreignkey:ShippingAddressID;association_foreignkey:ID"`
// ShippingAddressID   *string `gorm:"type:char(36);"`
// RemittanceAddress   Address
// RemittanceAddressID *string `gorm:"type:char(36);"`

// PrimaryContact     Contact `gorm:"foreignkey:PrimaryContactID;association_foreignkey:ID"`
// PrimaryContactID   *string `gorm:"type:char(36);"`
// SecondaryContact   Contact
// SecondaryContactID *string `gorm:"type:char(36);"`
