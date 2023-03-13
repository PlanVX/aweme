package query

import (
	"github.com/PlanVX/aweme/internal/config"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

// NewGormDB returns a new gorm db instance
func NewGormDB(config *config.Config, logger *zap.Logger) (*gorm.DB, error) {
	l := zapgorm2.New(logger)
	l.SetAsDefault()
	db, err := gorm.Open(mysql.Open(config.MySQL.DSN), &gorm.Config{SkipDefaultTransaction: true, Logger: l, QueryFields: true})
	if err != nil {
		return nil, err
	}
	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		return nil, err
	}
	return db, err
}

// close the gorm db instance
func gormClose(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	return sqlDB.Close()
}
