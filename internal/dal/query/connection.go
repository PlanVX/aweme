/*
 * Copyright (c) 2023 The PlanVX Authors.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
