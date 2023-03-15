// Package routes : manage the echo route
package routes

import "go.uber.org/fx"

// Module routes module,
// it provides a new echo instance
// and adds all the routes to it
// and starts the server in fx.Lifecycle
var Module = fx.Module("routes",
	fx.Provide(NewEcho),
	fx.Invoke(decorators...),
	fx.Invoke(StartServer),
)
