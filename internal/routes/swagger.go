//go:build swagger
// +build swagger

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

package routes

import (
	"github.com/PlanVX/aweme/docs"
	"github.com/PlanVX/aweme/internal/config"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// trick to make sure decorators is initialized before it is used
var _ = func() any {
	decorators = append(decorators, registerSwagger)
	return nil
}()

// registerSwagger registers swagger docs route
func registerSwagger(config *config.Config, e *echo.Echo) *echo.Echo {
	docs.SwaggerInfo.BasePath = config.API.Prefix // set swagger base path same as echo group prefix
	e.GET("/swagger/*", echoSwagger.WrapHandler)  // add swagger docs route
	return e
}
