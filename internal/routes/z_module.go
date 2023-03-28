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

// Package routes : manage the echo route
package routes

import "go.uber.org/fx"

// Module routes module,
// it provides a new echo instance
// and adds all the routes to it
// and starts the server in fx.Lifecycle
var Module = fx.Module("routes",
	fx.Provide(
		NewEcho,
	),
	fx.Invoke(decorators...),
	fx.Invoke(fx.Annotate(
		NewHTTPServer,
		fx.OnStart(startHook),
		fx.OnStop(stopHook),
	)),
)
