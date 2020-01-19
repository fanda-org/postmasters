package system

// OrgUser model
type OrgUser struct {
	Organization Organization `gorm:"foreignkey:OrgID"` //;association_foreignkey:ID"`
	OrgID        string       `gorm:"type:char(36);primary_key"`
	User         User         // `gorm:"foreignkey:UserID;association_foreignkey:ID"`
	UserID       string       `gorm:"type:char(36);primary_key"`
	Roles        []Role       `gorm:"many2many:user_roles;foreignkey:OrgID,UserID;association_foreignkey:ID;jointable_foreignkey:org_id,user_id;association_jointable_foreignkey:role_id;"`
}
