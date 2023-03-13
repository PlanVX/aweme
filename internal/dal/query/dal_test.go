package query

import (
	"context"
	"github.com/PlanVX/aweme/internal/config"
	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestNewRedisUniversalClient(t *testing.T) {
	s := miniredis.RunT(t)
	c := config.Config{}
	c.Redis.Addr = []string{s.Addr()}
	client := NewRedisUniversalClient(&c, zap.L())
	assert.NotNil(t, client)
	ctx := context.Background()
	t.Run("HIncr on non number", func(t *testing.T) {
		err := client.HSet(ctx, "test", "test", "test").Err()
		assert.NoError(t, err)
		client.HIncrBy(ctx, "test", "test", 1)
	})

	t.Run("HIncr Keys on non number", func(t *testing.T) {
		err := client.HSet(ctx, "test", "test", "test").Err()
		assert.NoError(t, err)
		fields := []HashField{
			{
				Key:   "test",
				Field: "test",
			},
			{
				Key:   "1",
				Field: "test",
			},
		}
		client.HKeyFieldsIncrBy(ctx, fields, 1)
	})
	err := closeRedis(client)
	assert.NoError(t, err)
	otel, err := RedisOtel(client)
	assert.NoError(t, err)
	assert.NotNil(t, otel)
	s.Close()
}
