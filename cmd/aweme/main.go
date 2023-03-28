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

package main

import (
	"github.com/PlanVX/aweme/internal/api"
	"github.com/PlanVX/aweme/internal/config"
	"github.com/PlanVX/aweme/internal/dal/query"
	"github.com/PlanVX/aweme/internal/logic"
	"github.com/PlanVX/aweme/internal/otel"
	"github.com/PlanVX/aweme/internal/routes"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

// @title aweme
// @version 1.0
// @description aweme api
// @contact.name PlanVX
// @contact.url https://github.com/PlanVX
// @license.name Apache 2.0
// @license.url https://github.com/PlanVX/aweme/blob/main/LICENSE
// @host localhost:8080
// @BasePath /v1
func main() {
	app := fx.New(
		fx.Provide(config.NewConfig, newZapLogger),
		fx.WithLogger(fxLogger),
		otel.Module,
		query.Module,
		logic.Module,
		api.Module,
		routes.Module,
	)
	app.Run()
}

// replace the default logger with wrapped zap logger
func fxLogger(logger *zap.Logger) fxevent.Logger { return &fxevent.ZapLogger{Logger: logger} }

// newZapLogger returns a new zap logger
func newZapLogger(c *config.Config) (*zap.Logger, error) {
	if !c.Release {
		return zap.NewDevelopment()
	}
	return zap.NewProduction()
}
