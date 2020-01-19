package system

// RolePrivilege model
type RolePrivilege struct {
	Role        Role
	RoleID      string   `gorm:"type:char(36);primary_key"`
	Resource    Resource //`gorm:"foreignkey:ResourceID;association_foreignkey:ID"`
	ResourceID  string   `gorm:"type:char(36);primary_key"`
	CanCreate   bool     `gorm:"not null"`
	CanRead     bool     `gorm:"not null"`
	CanUpdate   bool     `gorm:"not null"`
	CanDelete   bool     `gorm:"not null"`
	CanActivate bool     `gorm:"not null"`
}
