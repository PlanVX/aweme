package query

import (
	"context"
	"github.com/PlanVX/aweme/internal/config"
	driver "github.com/go-sql-driver/mysql"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

// NewGormDB returns a new gorm db instance
func NewGormDB(config *config.Config, logger *zap.Logger, lf fx.Lifecycle) (*gorm.DB, error) {
	l := zapgorm2.New(logger)
	l.SetAsDefault()
	db, err := gorm.Open(mysql.Open(genDsn(config)),
		&gorm.Config{SkipDefaultTransaction: true, Logger: l, QueryFields: true})
	lf.Append(fx.Hook{OnStop: func(ctx context.Context) error {
		sqlDB, e := db.DB()
		if e != nil {
			return e
		}
		return sqlDB.Close()
	}})
	if config.Release == false { // debug mode when not release
		db = db.Debug()
	}
	return db, err
}

func genDsn(config *config.Config) string {
	m := driver.Config{
		User:                 config.MySQL.Username,
		Passwd:               config.MySQL.Password,
		Net:                  "tcp",
		Addr:                 config.MySQL.Address,
		DBName:               config.MySQL.Database,
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	return m.FormatDSN()
}
