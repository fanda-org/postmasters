package database

import (
	"github.com/fanda-org/postmasters/config"
	"github.com/fanda-org/postmasters/database/models/accounting"
	"github.com/fanda-org/postmasters/database/models/system"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"    // empty or blank
	_ "github.com/jinzhu/gorm/dialects/mysql"    // empty or blank
	_ "github.com/jinzhu/gorm/dialects/postgres" // empty or blank
)

//var dialect = config.GetConfig().DB.Dialect
//var connectionString = config.GetConfig().DB.GetConnectionString()

// New creates db object (constructor)
func New(dbConfig *config.DBConfig) (db *gorm.DB, err error) {
	db, err = gorm.Open(dbConfig.Dialect, dbConfig.GetConnectionString())
	if err != nil {
		return
	}
	return
}

// Open database connection
func Open(dbConfig *config.DBConfig) (db *gorm.DB) {
	db, _ = gorm.Open(dbConfig.Dialect, dbConfig.GetConnectionString())
	return
}

// Migrate migrates to database
func Migrate(dbConfig *config.DBConfig) (err error) {
	db, err := New(dbConfig)
	if err != nil {
		return
	}
	defer db.Close()
	// Log
	// db.SetLogger(gorm.Logger{revel.TRACE})
	// db.SetLogger(log.New(os.Stdout, "\r\n", 0))

	// System models - Migration
	db.AutoMigrate(&system.Address{}, &system.Contact{})
	db.AutoMigrate(&system.Organization{}, &system.User{}, &system.OrgUser{}, &system.Role{})
	db.AutoMigrate(&system.ResourceCategory{}, &system.Resource{}, &system.RolePrivilege{})
	// Accounting models - Migration
	db.AutoMigrate(&accounting.LedgerGroup{}, &accounting.Ledger{})
	db.AutoMigrate(&accounting.Bank{}, &accounting.Party{})
	// AccountYear, LedgerBalance models - Migration
	db.AutoMigrate(&accounting.AccountYear{}, &accounting.LedgerBalance{})

	// org_addresses
	db.Table("org_addresses").AddForeignKey("org_id", "organizations(id)", "RESTRICT", "RESTRICT")
	db.Table("org_addresses").AddForeignKey("addr_id", "addresses(id)", "RESTRICT", "RESTRICT")
	// org_contacts
	db.Table("org_contacts").AddForeignKey("org_id", "organizations(id)", "RESTRICT", "RESTRICT")
	db.Table("org_contacts").AddForeignKey("contact_id", "contacts(id)", "RESTRICT", "RESTRICT")
	// org_users
	db.Table("org_users").AddForeignKey("org_id", "organizations(id)", "RESTRICT", "RESTRICT")
	db.Table("org_users").AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	// user_roles
	db.Table("user_roles").AddForeignKey("org_id, user_id", "org_users(org_id, user_id)", "RESTRICT", "RESTRICT")
	db.Table("user_roles").AddForeignKey("role_id", "roles(id)", "RESTRICT", "RESTRICT")

	// org<=>addresses
	//db.Model(&models.Organization{}).AddForeignKey("billing_address_id", "addresses(id)", "RESTRICT", "RESTRICT")
	//db.Model(&models.Organization{}).AddForeignKey("shipping_address_id", "addresses(id)", "RESTRICT", "RESTRICT")
	//db.Model(&models.Organization{}).AddForeignKey("remittance_address_id", "addresses(id)", "RESTRICT", "RESTRICT")
	// org<=>contacts
	//db.Model(&models.Organization{}).AddForeignKey("primary_contact_id", "contacts(id)", "RESTRICT", "RESTRICT")
	//db.Model(&models.Organization{}).AddForeignKey("secondary_contact_id", "contacts(id)", "RESTRICT", "RESTRICT")

	// resource -> resource_category
	db.Model(&system.Resource{}).AddForeignKey("resource_category_id", "resource_categories(id)", "RESTRICT", "RESTRICT")
	// roleprivilege
	db.Model(&system.RolePrivilege{}).AddForeignKey("role_id", "roles(id)", "RESTRICT", "RESTRICT")
	db.Model(&system.RolePrivilege{}).AddForeignKey("resource_id", "resources(id)", "RESTRICT", "RESTRICT")

	// legergroup
	db.Model(&accounting.LedgerGroup{}).AddForeignKey("parent_id", "ledger_groups(id)", "RESTRICT", "RESTRICT")
	db.Model(&accounting.LedgerGroup{}).AddForeignKey("org_id", "organizations(id)", "RESTRICT", "RESTRICT")
	// ledger
	db.Model(&accounting.Ledger{}).AddForeignKey("ledger_group_id", "ledger_groups(id)", "RESTRICT", "RESTRICT")
	db.Model(&accounting.Ledger{}).AddForeignKey("parent_id", "ledgers(id)", "RESTRICT", "RESTRICT")
	db.Model(&accounting.Ledger{}).AddForeignKey("org_id", "organizations(id)", "RESTRICT", "RESTRICT")
	// bank
	db.Model(&accounting.Bank{}).AddForeignKey("ledger_id", "ledgers(id)", "RESTRICT", "RESTRICT")
	db.Model(&accounting.Bank{}).AddForeignKey("addr_id", "addresses(id)", "RESTRICT", "RESTRICT")
	db.Model(&accounting.Bank{}).AddForeignKey("contact_id", "contacts(id)", "RESTRICT", "RESTRICT")
	// party
	db.Model(&accounting.Party{}).AddForeignKey("ledger_id", "ledgers(id)", "RESTRICT", "RESTRICT")
	db.Model(&accounting.Party{}).AddForeignKey("party_org_id", "organizations(id)", "RESTRICT", "RESTRICT")

	// accountyear
	db.Model(&accounting.AccountYear{}).AddForeignKey("org_id", "organizations(id)", "RESTRICT", "RESTRICT")
	// ledgerbalance
	db.Model(&accounting.LedgerBalance{}).AddForeignKey("ledger_id", "ledgers(id)", "RESTRICT", "RESTRICT")
	db.Model(&accounting.LedgerBalance{}).AddForeignKey("year_id", "account_years(id)", "RESTRICT", "RESTRICT")

	return
}
