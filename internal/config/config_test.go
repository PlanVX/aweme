package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
	_, err := NewConfig()
	assert.Error(t, err)
}

func TestParseConfig(t *testing.T) {
	data, prefix := []byte(`
api:
  prefix: $API_PREFIX`), "/test"
	assert.NoError(t, os.Setenv("API_PREFIX", prefix))
	c, err := parseConfig(data)
	assert.NoError(t, err)
	assert.NoError(t, os.Setenv("API_PREFIX", prefix))
	assert.Equal(t, prefix, c.API.Prefix)
}
