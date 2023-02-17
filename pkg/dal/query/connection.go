package query

import (
	"github.com/PlanVX/aweme/pkg/config"
	driver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewGormDB returns a new gorm db instance
func NewGormDB(config *config.Config) (*gorm.DB, error) {
	m := driver.Config{
		User:                 config.MySQL.Username,
		Passwd:               config.MySQL.Password,
		Net:                  "tcp",
		Addr:                 config.MySQL.Address,
		DBName:               config.MySQL.Database,
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	return gorm.Open(mysql.Open(m.FormatDSN()), &gorm.Config{})
}
