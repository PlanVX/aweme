package query

import (
	"github.com/PlanVX/aweme/pkg/config"
	"github.com/redis/go-redis/v9"
)

// NewRedisUniversalClient returns a redis client
func NewRedisUniversalClient(config *config.Config) redis.UniversalClient {
	return redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: config.Redis.Addr,
		DB:    config.Redis.DB,
	})
}
