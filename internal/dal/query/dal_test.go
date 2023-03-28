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
