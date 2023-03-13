package query

import (
	"context"
	"github.com/PlanVX/aweme/internal/config"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// RDB is the database of redis
type RDB struct {
	redis.UniversalClient
	logger *zap.Logger
}

// NewRedisUniversalClient returns a redis client extended with zap logger on incr error
func NewRedisUniversalClient(config *config.Config, lf fx.Lifecycle, logger *zap.Logger) *RDB {
	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    config.Redis.Addr,
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	})
	lf.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				return client.Ping(ctx).Err()
			},
			OnStop: func(ctx context.Context) error {
				return client.Close()
			},
		})
	return &RDB{
		client,
		logger,
	}
}

// HashField is the struct for specifying field of redis hash structure
type HashField struct {
	Key   string
	Field string
}

// HKeyFieldsIncrBy for given keys and fields, increment the value
func (r *RDB) HKeyFieldsIncrBy(ctx context.Context, keyFields []HashField, value int64) {
	cmder, err := r.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		for _, keyField := range keyFields {
			pipe.HIncrBy(ctx, keyField.Key, keyField.Field, value)
		}
		return nil
	})
	if err != nil {
		r.logger.Error("HIncrByKeys", zap.Error(err))
		for _, cmd := range cmder {
			r.logger.Error("redis cmd", zap.String("cmd", cmd.String()), zap.Error(cmd.Err()))
		}
	}
}

// HIncrBy for given key and field, increment the value
func (r *RDB) HIncrBy(ctx context.Context, key, field string, value int64) {
	cmd := r.UniversalClient.HIncrBy(ctx, key, field, value)
	if err := cmd.Err(); err != nil {
		r.logger.Error("HIncrBy", zap.Error(err), zap.String("cmd", cmd.String()))
	}
}
