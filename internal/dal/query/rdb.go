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
	"context"

	"github.com/PlanVX/aweme/internal/config"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

// RDB is the database of redis
type RDB struct {
	redis.UniversalClient
	logger *zap.Logger
}

// NewRedisUniversalClient returns a redis client extended with zap logger on incr error
func NewRedisUniversalClient(config *config.Config, logger *zap.Logger) *RDB {
	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    config.Redis.Addr,
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	})
	return &RDB{
		client,
		logger,
	}
}

// closeRedis closes the redis client
func closeRedis(client *RDB) error {
	return client.Close()
}

// RedisOtel extends redis client with open telemetry tracing
func RedisOtel(rdb *RDB) (*RDB, error) {
	if err := redisotel.InstrumentTracing(rdb.UniversalClient); err != nil {
		return nil, err
	}
	return rdb, nil
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
