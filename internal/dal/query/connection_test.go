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
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

const dsn = "aweme:aweme@tcp(localhost:3306)/aweme?charset=utf8mb4&parseTime=True&loc=Local&allowNativePasswords=true"

func TestPlugin(t *testing.T) {
	_, g, _, err := mockDB(t)
	c := config.Config{}
	c.MySQL.Replicas = []string{dsn}
	_, err = withPlugins(&c, g)
	assert.Nil(t, err)
}

func TestNewGormDB(t *testing.T) {
	c := config.Config{}
	c.MySQL.DSN = dsn
	db, err := NewGormDB(&c, zap.L())
	assert.Nil(t, err)
	assert.NotNil(t, db)
	err = gormClose(db)
	assert.Nil(t, err)
}
