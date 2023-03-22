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
