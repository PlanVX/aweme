package query

import (
	"github.com/PlanVX/aweme/internal/config"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"moul.io/zapgorm2"
)

// NewGormDB returns a new gorm db instance
func NewGormDB(config *config.Config, logger *zap.Logger) (*gorm.DB, error) {
	l := zapgorm2.New(logger)
	l.SetAsDefault()
	return gorm.Open(mysql.Open(config.MySQL.DSN), &gorm.Config{SkipDefaultTransaction: true, Logger: l, QueryFields: true})
}

func withPlugins(config *config.Config, db *gorm.DB) (*gorm.DB, error) {
	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		return nil, err
	}
	if len(config.MySQL.Replicas) > 0 {
		c := replicas(config.MySQL.Replicas)
		if err := db.Use(dbresolver.Register(c)); err != nil {
			return nil, err
		}
	}
	return db, nil
}

func replicas(replicas []string) dbresolver.Config {
	var dbs []gorm.Dialector
	for _, replica := range replicas {
		dbs = append(dbs, mysql.Open(replica))
	}
	return dbresolver.Config{Replicas: dbs}
}

// close the gorm db instance
func gormClose(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	return sqlDB.Close()
}
