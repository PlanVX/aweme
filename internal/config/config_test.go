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
