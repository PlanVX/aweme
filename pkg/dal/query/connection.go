package query

import (
	"context"
	"github.com/PlanVX/aweme/pkg/config"
	driver "github.com/go-sql-driver/mysql"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

// NewGormDB returns a new gorm db instance
func NewGormDB(config *config.Config, logger *zap.Logger, lf fx.Lifecycle) (*gorm.DB, error) {
	m := driver.Config{
		User:                 config.MySQL.Username,
		Passwd:               config.MySQL.Password,
		Net:                  "tcp",
		Addr:                 config.MySQL.Address,
		DBName:               config.MySQL.Database,
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	l := zapgorm2.New(logger)
	l.SetAsDefault()
	db, err := gorm.Open(mysql.Open(m.FormatDSN()), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 l,
	})
	lf.Append(fx.Hook{OnStop: func(ctx context.Context) error {
		sqlDB, e := db.DB()
		if e != nil {
			return e
		}
		return sqlDB.Close()
	}})
	return db, err
}
