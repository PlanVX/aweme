#!/usr/bin/env bash
#
# Copyright (c) 2023 The PlanVX Authors.
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#     http://www.apache.org/licenses/LICENSE-2.
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

# generate swagger documentation
go install github.com/swaggo/swag/cmd/swag@latest
swag init -g cmd/aweme/main.go -p snakecase

# optional: generate mocks for testing
# mockery code generation
# for mocking interfaces to test logic
if [ -n "$MOCKERY" ]; then
  go install github.com/vektra/mockery/v2@v2.20.0
  go generate ./...
fi
