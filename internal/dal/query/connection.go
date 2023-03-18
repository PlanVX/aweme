package query

import (
	"github.com/PlanVX/aweme/internal/config"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

// NewGormDB returns a new gorm db instance
func NewGormDB(config *config.Config, logger *zap.Logger) (*gorm.DB, error) {
	l := zapgorm2.New(logger)
	l.SetAsDefault()
	return gorm.Open(mysql.Open(config.MySQL.DSN), &gorm.Config{SkipDefaultTransaction: true, Logger: l, QueryFields: true})
}

// close the gorm db instance
func gormClose(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	return sqlDB.Close()
}
