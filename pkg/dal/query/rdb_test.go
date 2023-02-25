package query

import (
	"github.com/PlanVX/aweme/pkg/config"
	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRedisUniversalClient(t *testing.T) {
	s := miniredis.RunT(t)
	c := config.Config{
		Redis: struct {
			Addr     []string `yaml:"address"`
			Password string   `yaml:"password"`
			DB       int      `yaml:"db"`
		}{
			Addr: []string{s.Addr()},
		},
	}
	client := NewRedisUniversalClient(&c)
	assert.NotNil(t, client)
}
